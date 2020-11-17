package ionotify

import (
	"context"
	"encoding/gob"
	"net"
	"reflect"

	"github.com/alivanz/go-notify"
)

// Serve notify to net.Listener
func Serve(ctx context.Context, listener net.Listener, n *notify.Interface) error {
	for {
		select {
		case <-ctx.Done():
		default:
		}
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			defer conn.Close()
			enc := gob.NewEncoder(conn)
			Encode(ctx, enc, n)
		}()
	}
}

// Listen dial and listen to notify
func Listen(ctx context.Context, host string, Type reflect.Type, n *notify.Interface) error {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err
	}
	defer conn.Close()
	dec := gob.NewDecoder(conn)
	return Decode(ctx, Type, dec, n)
}
