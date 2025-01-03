package service

func RepeatFunc[T any, K any](done <-chan K, fn func() T) (out <-chan T) {

	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream

}
