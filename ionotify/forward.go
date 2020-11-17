package ionotify

import (
	"context"
	"reflect"

	"github.com/alivanz/go-notify"
)

// Encoder generic encoder
type Encoder interface {
	Encode(v interface{}) error
}

// Decoder generic decoder
type Decoder interface {
	Decode(v interface{}) error
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
