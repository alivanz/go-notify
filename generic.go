package notify

type Listener interface {
	Listen() (interface{}, <-chan struct{})
}

type Notifier interface {
	Notify(v interface{})
}
