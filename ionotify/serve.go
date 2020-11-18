package ionotify

import (
	"context"
	"encoding/gob"
	"net"

	"github.com/alivanz/go-notify"
)

// ListenAndServe open tcp listen and serve notify
func ListenAndServe(ctx context.Context, listen string, n *notify.Interface) error {
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		return err
	}
	defer listener.Close()
	return Serve(ctx, listener, n)
}

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
