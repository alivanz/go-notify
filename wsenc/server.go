package wsenc

import (
	"net/http"

	"github.com/alivanz/go-notify"
	"github.com/alivanz/go-notify/ionotify"
	"github.com/gorilla/websocket"
)

type Server struct {
	Notify   notify.Listener
	Upgrader *websocket.Upgrader
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := server.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	enc := &WsEncoding{
		Conn: conn,
	}
	ionotify.Encode(r.Context(), enc, server.Notify)
}
