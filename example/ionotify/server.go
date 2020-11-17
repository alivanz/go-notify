package main

import (
	"context"
	"log"
	"time"

	"github.com/alivanz/go-notify"
	"github.com/alivanz/go-notify/ionotify"
)

func main() {
	n := notify.NewInterface(nil)
	go ionotify.ListenAndServe(context.Background(), "0.0.0.0:6000", n)
	// feed data
	for range time.NewTicker(1 * time.Second).C {
		t := time.Now()
		n.Notify(t)
		log.Printf("server notify %v", t.String())
	}
}
