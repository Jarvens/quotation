// auth: kunlun
// date: 2019-01-09
// description: websocket server
package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrade = websocket.Upgrader{
	// allows cross
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, EnableCompression: true,
}

func Handler(w http.ResponseWriter, r *http.Request) {

	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)

	// handlerShark
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = InitConn(wsConn); err != nil {
		goto ERR
	}

	// always send message
	go func() {
		var err error
		for {
			if err = conn.WriteMessage([]byte("pong address" + conn.wsConnect.RemoteAddr().String())); err != nil {
				return
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}

func init() {

}
