package notify

type Interface struct {
	base *Base
	v    interface{}
}

func NewInterface(v interface{}) *Interface {
	return &Interface{
		base: NewBase(),
		v:    v,
	}
}

func (i *Interface) Notify(v interface{}) {
	i.base.Notify(func() {
		i.v = v
	})
}

func (i *Interface) Listen(interface{}) (interface{}, <-chan struct{}) {
	var v interface{}
	return v, i.base.Listen(func() {
		v = i.v
	})
}
