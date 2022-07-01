// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"log"
	"time"
)

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())

	// Start the pool of workers
	Start()
}

// simplePool implements SimplePool
type simplePool struct {
	maxConcurrent int
	workQueue     chan func()
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	if maxConcurrent < 1 {
		return nil
	}
	sp := &simplePool{
		maxConcurrent,
		make(chan func()),
	}
	return sp
}

func (sp *simplePool) Submit(task func()) {
	sp.workQueue <- task
}

func (sp *simplePool) Start() {
	sp.start()
}

func (sp *simplePool) start() {
	for i := 1; i <= sp.maxConcurrent; i++ {
		log.Printf("simple_worker_pool: Worker %d is up", i)
		go func(workerID int) {
			for task := range sp.workQueue {
				log.Printf("simple_worker_pool: Worker %d is busy", workerID)
				task()
				time.Sleep(time.Second * 3)
				log.Printf("simple_worker_pool: Worker %d is done", workerID)
			}
		}(i)
	}
}
