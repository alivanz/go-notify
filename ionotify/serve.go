package ionotify

import (
	"context"
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
	return ServeWithDecoder(ctx, listener, n, NewEncoder)
}

// ServeWithDecoder notify to net.Listener with custom decoder
func ServeWithDecoder(ctx context.Context, listener net.Listener, n *notify.Interface, encf NewEncoderFunc) error {
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
			enc := encf(conn)
			Encode(ctx, enc, n)
		}()
	}
}
