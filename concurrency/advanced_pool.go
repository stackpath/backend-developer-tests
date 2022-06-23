package concurrency

import (
	"context"
	"errors"
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

type Job struct {
	Exec func(context.Context)
}

type AdvancedPoolImpl struct {
	runningWorkers chan bool
	waitingWorkers chan Job
	runningGroup   *sync.WaitGroup
	poolCtx        context.Context
	poolCloser     func()
}

func (api *AdvancedPoolImpl) Submit(ctx context.Context, exec func(context.Context)) error {
	for {
		select {
		case api.waitingWorkers <- Job{exec}:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		case <-api.poolCtx.Done():
			return errors.New("ErrPoolClosed")
		}
	}
}

func (api *AdvancedPoolImpl) Close(ctx context.Context) error {
	if api.poolCtx.Err() != nil {
		return errors.New("ErrPoolClosed")
	}
	waitFinish := make(chan bool)
	go func() {
		api.runningGroup.Wait()
		waitFinish <- true
	}()

	select {
	case <-waitFinish:
		close(api.waitingWorkers)
		close(api.runningWorkers)
		close(waitFinish)
		api.poolCloser()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (api *AdvancedPoolImpl) start() {
	for {
		select {
		case api.runningWorkers <- true:
			job := <-api.waitingWorkers
			api.runningGroup.Add(1)
			go func() {
				defer api.runningGroup.Done()
				job.Exec(context.WithValue(api.poolCtx, nil, nil))
				<-api.runningWorkers
			}()
		}
	}
}

// NewAdvancedPool creates a new AdvancedPool. maxSlots is the maximum total
// submitted tasks, running or waiting, that can be submitted before Submit
// blocks waiting for more room. maxConcurrent is the maximum tasks that can be
// running at any one time. An error is returned if maxSlots is less than
// maxConcurrent or if either value is not greater than zero.
func NewAdvancedPool(maxSlots, maxConcurrent int) (AdvancedPool, error) {
	if maxSlots < maxConcurrent || maxSlots < 0 || maxConcurrent < 0 {
		return nil, errors.New("Invalid maxSlots and maxConcurrent")
	}
	ctx, cancel := context.WithCancel(context.Background())
	api := &AdvancedPoolImpl{
		make(chan bool, maxConcurrent),
		make(chan Job, maxSlots),
		&sync.WaitGroup{},
		ctx,
		cancel}
	go api.start()
	return api, nil
}
