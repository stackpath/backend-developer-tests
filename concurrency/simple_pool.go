// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"fmt"
	"sync/atomic"
)

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	GetAllDone() chan bool
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
}

type SimplePoolObject struct {
	maxConcurrent int
	runningCount  atomic.Value //int
	jobWaiting    atomic.Value //bool
	finishedWork  chan bool
	allDone       chan bool
}

func (spo SimplePoolObject) GetAllDone() chan bool {
	return spo.allDone
}

func (spo *SimplePoolObject) Submit(func1 func()) {

	spo.runningCount.Store(spo.runningCount.Load().(int) + 1)
	fmt.Println("count value is ", spo.runningCount)

	if spo.runningCount.Load().(int) > spo.maxConcurrent {
		spo.jobWaiting.Store(true)
		fmt.Println("Waiting for finish work signal")
		<-spo.finishedWork
		if spo.runningCount.Load().(int) <= spo.maxConcurrent {
			spo.jobWaiting.Store(false)
		}
		fmt.Println("After the signal")
	}

	go func() {
		func1()
		spo.runningCount.Store(spo.runningCount.Load().(int) - 1)
		fmt.Println("count in go routine value after the job completion is ", spo.runningCount)
		if spo.jobWaiting.Load().(bool) {
			spo.finishedWork <- true
		}

		if spo.runningCount.Load() == 0 {
			spo.allDone <- true
		}
	}()
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	simplePool := &SimplePoolObject{
		maxConcurrent: maxConcurrent,
		finishedWork:  make(chan bool),
		allDone:       make(chan bool),
	}
	simplePool.runningCount.Store(0)
	return simplePool
}
