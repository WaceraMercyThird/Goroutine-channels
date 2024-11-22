package service

import (
	"sync"
)

func FanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(C <-chan T) {
		defer wg.Done()
		for i := range C {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, C := range channels {
		wg.Add(1)
		go transfer(C)
	}

	// Close the fannedInStream once all goroutines complete
	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}
