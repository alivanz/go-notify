package main

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/alivanz/go-notify"
	"github.com/alivanz/go-notify/ionotify"
)

func main() {
	n := notify.NewInterface(nil)
	go ionotify.Subscribe(context.Background(), "127.0.0.1:6000", reflect.TypeOf(time.Time{}), n)
	c := notify.Closed()
	for {
		select {
		case <-c:
			var i interface{}
			i, c = n.Listen()
			if i == nil {
				continue
			}
			t := i.(*time.Time)
			log.Printf("server notify %v", t.String())
		}
	}
}
