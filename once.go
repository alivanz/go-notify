package notify

import (
	"sync"
)

type NotifyOnce struct {
	c    chan struct{}
	once sync.Once
}

func NewNotifyOnce() *NotifyOnce {
	return &NotifyOnce{
		c: make(chan struct{}),
	}
}

func (notify *NotifyOnce) Notify() {
	notify.once.Do(func() {
		close(notify.c)
	})
}

func (notify *NotifyOnce) Listen() <-chan struct{} {
	return notify.c
}
