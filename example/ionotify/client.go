package main

import (
	"log"
	"reflect"
	"time"

	"github.com/alivanz/go-notify"
	"github.com/alivanz/go-notify/ionotify"
)

func main() {
	n := notify.NewInterface(nil)
	subs, err := ionotify.Subscribe("127.0.0.1:6000", reflect.TypeOf(time.Time{}), n)
	if err != nil {
		log.Fatal(err)
	}
	defer subs.Unsubscribe()
	c := notify.Closed()
	for {
		select {
		case err := <-subs.Err():
			log.Fatal(err)
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
