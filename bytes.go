package notify

type Bytes struct {
	base *Base
	v    []byte
}

func NewBytes(v []byte) *Bytes {
	return &Bytes{
		base: NewBase(),
		v:    v,
	}
}

func (i *Bytes) Notify(v []byte) {
	i.base.Notify(func() {
		i.v = v
	})
}

func (i *Bytes) Listen() (interface{}, <-chan struct{}) {
	return i.ListenBytes()
}
func (i *Bytes) ListenBytes() ([]byte, <-chan struct{}) {
	var v []byte
	return v, i.base.Listen(func() {
		v = i.v
	})
}
