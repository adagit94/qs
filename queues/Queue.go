package queues

import (
	"github.com/adagit94/qs/tasks"
)

type IQueue[R any] interface {
	Enqueue(t *tasks.Task[R])
	Loop()
	Drain()
}