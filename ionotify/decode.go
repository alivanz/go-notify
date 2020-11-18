package ionotify

import (
	"context"
	"reflect"

	"github.com/alivanz/go-notify"
)

// Decoder generic decoder
type Decoder interface {
	Decode(v interface{}) error
}

// Decode decode data and push it to notify
func Decode(ctx context.Context, Type reflect.Type, dec Decoder, n *notify.Interface) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		rdata := reflect.New(Type)
		if err := dec.Decode(rdata.Interface()); err != nil {
			return err
		}
		n.Notify(rdata.Interface())
	}
}
