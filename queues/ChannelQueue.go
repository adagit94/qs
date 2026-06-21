package queues

import (
	"fmt"
	"github.com/adagit94/qs/tasks"
)

type ChannelQueue[R any] struct {
	Queue chan *tasks.Task[R]
}

func (q *ChannelQueue[R]) Enqueue(t *tasks.Task[R]) {
	q.Queue <- t
}

func (q *ChannelQueue[R]) Loop() {
	for {
		t, ok := <-q.Queue

		if !ok {
			return
		}

		select {
		case <-t.Skip:
			fmt.Println("Task skipped.")
			continue

		default:
			t.Result <- t.Work()
		}
	}
}

func (q *ChannelQueue[R]) Drain() {
	close(q.Queue)
}
