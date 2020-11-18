package ionotify

type Subscription interface {
	Unsubscribe()
	Err() <-chan error
}
