package notify

// Error run function on different goroutine, and get error channel
func Error(f func() error) <-chan error {
	c := make(chan error, 1)
	go func() {
		defer close(c)
		c <- f()
	}()
	return c
}
