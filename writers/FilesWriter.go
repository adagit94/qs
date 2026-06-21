package writers

import (
	"github.com/adagit94/gotils/queues"
)

type saveFile[D buffData, R any] func(path string, data D) R
type writeTask[R any] = queues.Task[R]
type pathQueue[R any] = queues.ChannelQueue[R]
type pathsQueues[R any] map[string]*pathQueue[R]
type buffSize uint8
type buffData interface {
	string | []byte
}

type filesWriter[D buffData, R any] struct {
	saveFile    saveFile[D, R]
	pathsQueues pathsQueues[R]
	buffSize    buffSize
}

func (fw *filesWriter[D, R]) Schedule(path string, data D) *writeTask[R] {
	queuePtr, found := fw.pathsQueues[path]

	if !found {
		queuePtr = &pathQueue[R]{Queue: make(chan *writeTask[R], fw.buffSize)}
		fw.pathsQueues[path] = queuePtr

		go queuePtr.Loop()
	}

	taskPtr := &writeTask[R]{Work: func() R {
		return fw.saveFile(path, data)
	}}

	queuePtr.Queue <- taskPtr

	return taskPtr
}

type IFilesWriter[D buffData, R any] interface {
	Schedule(path string, data D) *writeTask[R]
}
