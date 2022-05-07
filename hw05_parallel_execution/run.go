package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded  = errors.New("errors limit exceeded")
	ErrWrongWorkersSettings = errors.New("wrong workers settings")
)

type Task func() error

func Run(tasks []Task, n, m int) error {
	if n < 1 {
		return ErrWrongWorkersSettings
	}

	if m < 1 {
		return ErrErrorsLimitExceeded
	}

	errLimit := int32(m)

	wg := sync.WaitGroup{}
	wg.Add(n)

	messageBus := make(chan Task, len(tasks))
	for _, task := range tasks {
		messageBus <- task
	}
	close(messageBus)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			worker(messageBus, &errLimit)
		}()
	}

	wg.Wait()

	if atomic.LoadInt32(&errLimit) <= 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(messageBus chan Task, errLimit *int32) {
	for job := range messageBus {
		err := job()
		if err != nil {
			atomic.AddInt32(errLimit, -1)
		}

		if atomic.LoadInt32(errLimit) < 1 {
			return
		}
	}
}
