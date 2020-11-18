package ionotify

import (
	"context"

	"github.com/alivanz/go-notify"
)

// Encoder generic encoder
type Encoder interface {
	Encode(v interface{}) error
}

// Encode push any notify data to encoder
func Encode(ctx context.Context, enc Encoder, n *notify.Interface) error {
	c := notify.Closed()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-c:
		}
		var data interface{}
		data, c = n.Listen()
		if err := enc.Encode(data); err != nil {
			return err
		}
	}
}
