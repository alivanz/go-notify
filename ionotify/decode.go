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
	conn io.ReadCloser
	n    *notify.Interface
	typ  reflect.Type
	err  chan error
}

// Subscribe dial and listen to notify
func Subscribe(host string, Type reflect.Type, n *notify.Interface) (Subscription, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	subs := &decSubs{
		conn: conn,
		n:    n,
		typ:  Type,
		err:  make(chan error, 1),
	}
	go subs.run()
	return subs, nil
}

func (subs *decSubs) run() {
	defer close(subs.err)
	dec := gob.NewDecoder(subs.conn)
	for {
		rdata := reflect.New(subs.typ)
		if err := dec.Decode(rdata.Interface()); err != nil {
			subs.err <- err
			return
		}
		subs.n.Notify(rdata.Interface())
	}
}

func (subs *decSubs) Unsubscribe() {
	subs.conn.Close()
}

func (subs *decSubs) Err() <-chan error {
	return subs.err
}
