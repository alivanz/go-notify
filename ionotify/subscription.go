package ionotify

// Subscription Subscription object
type Subscription interface {
	Unsubscribe()
	Err() <-chan error
}
