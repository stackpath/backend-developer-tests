package concurrency

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

// ErrPoolClosed is returned from AdvancedPool.Submit when the pool is closed
// before submission can be sent.
var ErrPoolClosed = errors.New("pool closed")

// AdvancedPool is a more advanced worker pool that supports cancelling the
// submission and closing the pool. All functions are safe to call from multiple
// goroutines.
type AdvancedPool interface {
	// Submit submits the given task to the pool, blocking until a slot becomes
	// available or the context is closed. The given context and its lifetime only
	// affects this function and is not the context passed to the callback. If the
	// context is closed before a slot becomes available, the context error is
	// returned. If the pool is closed before a slot becomes available,
	// ErrPoolClosed is returned. Otherwise the task is submitted to the pool and
	// no error is returned. The context passed to the callback will be closed
	// when the pool is closed.
	Submit(context.Context, func(context.Context)) error

	// Close closes the pool and waits until all submitted tasks have completed
	// before returning. If the pool is already closed, ErrPoolClosed is returned.
	// If the given context is closed before all tasks have finished, the context
	// error is returned. Otherwise, no error is returned.
	Close(context.Context) error
}

// advancedPool implements AdvancedPool for managing
// concurrent worker pools
type advancedPool struct {
	waitGroup sync.WaitGroup
	jobs      chan func(context.Context)
	jobClose  chan bool
	jobDone   chan bool
	jobCtx    context.Context
	jobCancel context.CancelFunc
	syncOnce  *sync.Once
}

// NewAdvancedPool creates a new AdvancedPool. maxSlots is the maximum total
// submitted tasks, running or waiting, that can be submitted before Submit
// blocks waiting for more room. maxConcurrent is the maximum tasks that can be
// running at any one time. An error is returned if maxSlots is less than
// maxConcurrent or if either value is not greater than zero.
func NewAdvancedPool(maxSlots, maxConcurrent int) (AdvancedPool, error) {
	if maxConcurrent < 1 {
		return nil, fmt.Errorf("advanced_worker_pool: Failed init, maxConcurrent < 1")
	}
	if maxSlots < 1 {
		return nil, fmt.Errorf("advanced_worker_pool: Failed init, maxSlots < 1")
	}
	if maxSlots < maxConcurrent {
		return nil, fmt.Errorf("advanced_worker_pool: Failed init, maxSlots < maxConcurrent")
	}

	jobDone := make(chan bool)
	jobClose := make(chan bool)
	jobCtx, jobCancel := context.WithCancel(context.Background())
	ap := &advancedPool{
		waitGroup: sync.WaitGroup{},
		jobs:      make(chan func(context.Context), maxSlots),
		jobClose:  jobClose,
		jobCtx:    jobCtx,
		jobCancel: jobCancel,
	}

	func() {
		defer close(jobDone)
		ap.waitGroup.Wait()
	}()

	go func() {
		defer jobCtx.Done()
		<-jobClose
	}()

	// start workers
	for i := 1; i <= maxConcurrent; i++ {
		workerID := i
		log.Printf("advanced_worker_pool: Worker %d is starting", workerID)
		go ap.start(workerID)
	}

	return ap, nil
}

// Submit queues jobs for worker queue
func (ap *advancedPool) Submit(ctx context.Context, job func(context.Context)) error {
	select {
	case ap.jobs <- job:
		return nil
	case <-ap.jobClose:
		return ErrPoolClosed
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Close pool
func (ap *advancedPool) Close(ctx context.Context) error {
	select {
	case <-ap.jobDone:
		ap.syncOnce.Do(func() {
			close(ap.jobClose)
		})
		return nil
	case <-ap.jobClose:
		return ErrPoolClosed
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (ap *advancedPool) start(workerID int) {
	for job := range ap.jobs {
		log.Printf("advanced_worker_pool: Worker %d is busy", workerID)
		ap.waitGroup.Add(1)
		job(ap.jobCtx)
		ap.waitGroup.Done()
		log.Printf("advanced_worker_pool: Worker %d is done", workerID)
	}
}
