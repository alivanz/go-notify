package notify

type Listener interface {
	Listen() (interface{}, <-chan struct{})
}
