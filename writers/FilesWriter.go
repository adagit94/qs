package writers

import (
	"fmt"
	"github.com/adagit94/err"
	"github.com/adagit94/qs/queues"
	"github.com/adagit94/qs/tasks"
)

type SaveFile[D BuffData, R any] func(path string, data D) R
type WriteTask[R any] = tasks.Task[R]
type PathQueue[R any] = queues.ChannelQueue[R]
type PathsQueues[R any] map[string]*PathQueue[R]
type BuffSize uint8
type BuffData interface {
	string | []byte
}

type FilesWriter[D BuffData, R any] struct {
	SaveFile    SaveFile[D, R]
	PathsQueues PathsQueues[R]
	BuffSize    BuffSize
	Closed      chan bool
}

func (fw *FilesWriter[D, R]) Schedule(path string, data D) *WriteTask[R] {
	queuePtr, found := fw.PathsQueues[path]

	if !found {
		queuePtr = &PathQueue[R]{Queue: make(chan *WriteTask[R], fw.BuffSize), Closed: make(chan bool, 1)}
		fw.PathsQueues[path] = queuePtr

		go queuePtr.Loop()
	}

	taskPtr := &WriteTask[R]{Result: make(chan R, 1), Skip: make(chan bool, 1), Work: func() R {
		return fw.SaveFile(path, data)
	}}

	queuePtr.Enqueue(taskPtr)

	return taskPtr
}

func (fw *FilesWriter[D, R]) Close() (bool, *err.Err) {
	for _, v := range fw.PathsQueues {
		v.Close()
	}

	for k, v := range fw.PathsQueues {
		if closed := <-v.Closed; !closed {
			err := &err.Err{Code: err.ChannelNotClosed, Message: fmt.Sprint("Channel for a path ", k, " wasn't closed successfully.")}
			fw.Closed <- false
			return false, err
		}
	}

	fw.Closed <- true
	return true, nil
}

type IFilesWriter[D BuffData, R any] interface {
	Schedule(path string, data D) *WriteTask[R]
	Close() (bool, *err.Err)
}
