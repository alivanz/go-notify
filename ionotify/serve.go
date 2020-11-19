package ionotify

import (
	"context"
	"net"

	"github.com/alivanz/go-notify"
)

type Server interface {
	Stop()
	Err() <-chan error
}

type server struct {
	listener net.Listener
	encf     NewEncoderFunc
	n        notify.Listener
	err      chan error
}

// ListenAndServe open tcp listen and serve notify
func ListenAndServe(listen string, n notify.Listener) (Server, error) {
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		return nil, err
	}
	return Serve(listener, n)
}

// Serve notify to net.Listener
func Serve(listener net.Listener, n notify.Listener) (Server, error) {
	return ServeWithEncoder(listener, n, NewEncoder)
}

// ServeWithEncoder notify to net.Listener with custom encoder
func ServeWithEncoder(listener net.Listener, n notify.Listener, encf NewEncoderFunc) (Server, error) {
	s := &server{
		listener: listener,
		encf:     encf,
		n:        n,
		err:      make(chan error, 1),
	}
	go s.run()
	return s, nil
}

func (s *server) run() {
	defer close(s.err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.err <- err
			return
		}
		go func() {
			defer conn.Close()
			enc := s.encf(conn)
			Encode(ctx, enc, s.n)
		}()
	}
}

func (s *server) Stop() {
	s.listener.Close()
}
func (s *server) Err() <-chan error {
	return s.err
}
