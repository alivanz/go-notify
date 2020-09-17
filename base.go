package notify

import (
	"sync"
)

type Base struct {
	c    chan struct{}
	lock sync.RWMutex
}

func NewBase() *Base {
	return &Base{
		c: make(chan struct{}),
	}
}

func (base *Base) Notify(f func()) {
	base.lock.Lock()
	if f != nil {
		f()
	}
	close(base.c)
	base.c = make(chan struct{})
	base.lock.Unlock()
}

func (base *Base) Listen(f func()) <-chan struct{} {
	base.lock.RLock()
	if f != nil {
		f()
	}
	c := base.c
	base.lock.RUnlock()
	return c
}
