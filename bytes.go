package notify

import (
	"reflect"
)

type Bytes Interface

func NewBytes(v []byte) *Bytes {
	return (*Bytes)(NewInterface(v))
}
func (i *Bytes) Type() reflect.Type {
	return reflect.TypeOf([]byte{})
}
func (i *Bytes) Interface() *Interface {
	return (*Interface)(i)
}

func (i *Bytes) Notify(v []byte) {
	i.Interface().Notify(v)
}

func (i *Bytes) Listen() ([]byte, <-chan struct{}) {
	v, c := i.Interface().Listen()
	return v.([]byte), c
}
