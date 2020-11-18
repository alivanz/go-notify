package main

import (
	"log"
	"time"

	"github.com/alivanz/go-notify"
	"github.com/alivanz/go-notify/ionotify"
)

func main() {
	n := notify.NewInterface(nil)
	server, err := ionotify.ListenAndServe("0.0.0.0:6000", n)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Stop()
	// feed data
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case err := <-server.Err():
			log.Fatal(err)
		case <-ticker.C:
		}
		t := time.Now()
		n.Notify(t)
		log.Printf("server notify %v", t.String())
	}
}
