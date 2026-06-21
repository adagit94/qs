package tasks

type Task[R any] struct {
	Work func() R
	Result chan R 
	Skip chan bool
}
