package queues

import (
	"github.com/adagit94/qs/tasks"
)

type ChannelQueue[R any] struct {
	Queue  chan *tasks.Task[R]
	Closed chan bool
}

func (q *ChannelQueue[R]) Enqueue(t *tasks.Task[R]) {
	q.Queue <- t
}

func (q *ChannelQueue[R]) Loop() {
	for {
		t, ok := <-q.Queue

		if !ok {
			q.Closed <- true
			return
		}

		select {
		case <-t.Skip:
			continue

		default:
			t.Result <- t.Work()
		}
	}
}

func (q *ChannelQueue[R]) Close() {
	close(q.Queue)
}
