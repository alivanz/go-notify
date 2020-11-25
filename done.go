package notify

// Done run function on different goroutine, and close iif function finished
func Done(f func()) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		f()
	}()
	return c
}
