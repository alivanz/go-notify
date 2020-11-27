package wsenc

import (
	"github.com/gorilla/websocket"
)

// WsEncoding a wrapper for websocket
// implements ionotify.Encoder and ionotify.Decoder
type WsEncoding struct {
	Conn *websocket.Conn
}

func (ws *WsEncoding) Encode(v interface{}) error {
	return ws.Conn.WriteJSON(v)
}

func (ws *WsEncoding) Decode(v interface{}) error {
	return ws.Conn.ReadJSON(v)
}
