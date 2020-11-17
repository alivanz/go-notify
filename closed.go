package notify

var (
	closed chan struct{}
)

func init() {
	closed = make(chan struct{})
	close(closed)
}

func Closed() <-chan struct{} {
	return closed
}
