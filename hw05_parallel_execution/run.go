package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	availableErrors := int32(m)

	taskCh := make(chan Task, len(tasks))

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for atomic.LoadInt32(&availableErrors) > 0 {
				task, ok := <-taskCh
				if !ok {
					break
				}

				if err := task(); err != nil {
					atomic.AddInt32(&availableErrors, -1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&availableErrors) > 0 {
			taskCh <- task
		}
	}
	close(taskCh)

	wg.Wait()

	if availableErrors <= 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
