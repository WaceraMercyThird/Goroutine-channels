package service

func OriginalTake[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:

			}
		}
	}()
	return taken
}

func Take(done <-chan int, input <-chan int, count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			select {
			case v, ok := <-input:
				if !ok {
					return
				}
				out <- v
			case <-done:
				return
			}
		}
	}()
	return out
}
