// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package con

import "fmt"

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
}

type worker struct {
	id int
	jobs chan func()
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	jobs := make(chan func(), maxConcurrent)

	for w := 1; w <= maxConcurrent; w++ {
		go doWork(w, jobs)
	}
	return worker{jobs: jobs}
}


func (w worker) Submit(f func()) {
	w.jobs <- f
}

func doWork(id int, jobs <-chan func()) {
	for f := range jobs {
		fmt.Printf("Worker %d", id)
		f()
	}
}
