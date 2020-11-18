package ionotify

import (
	"encoding/gob"
	"io"
	"net"
	"reflect"

	"github.com/alivanz/go-notify"
)

// Decoder generic decoder
type Decoder interface {
	Decode(v interface{}) error
}

type decSubs struct {
	dec    Decoder
	closer io.Closer
	n      *notify.Interface
	typ    reflect.Type
	err    chan error
}

// Subscribe dial and listen to notify
func Subscribe(host string, Type reflect.Type, n *notify.Interface) (Subscription, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return SubscribeFromConn(conn, Type, n)
}

// SubscribeFromConn listen to notify
func SubscribeFromConn(conn io.ReadCloser, Type reflect.Type, n *notify.Interface) (Subscription, error) {
	return SubscribeFromDecoder(gob.NewDecoder(conn), conn, Type, n)
}

// SubscribeFromDecoder listen to notify
func SubscribeFromDecoder(dec Decoder, closer io.Closer, Type reflect.Type, n *notify.Interface) (Subscription, error) {
	subs := &decSubs{
		dec:    dec,
		closer: closer,
		n:      n,
		typ:    Type,
		err:    make(chan error, 1),
	}
	go subs.run()
	return subs, nil
}

func (subs *decSubs) run() {
	defer close(subs.err)
	for {
		rdata := reflect.New(subs.typ)
		if err := subs.dec.Decode(rdata.Interface()); err != nil {
			subs.err <- err
			return
		}
		subs.n.Notify(rdata.Interface())
	}
}

func (subs *decSubs) Unsubscribe() {
	subs.closer.Close()
}

func (subs *decSubs) Err() <-chan error {
	return subs.err
}
