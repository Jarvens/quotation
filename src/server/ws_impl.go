// auth: kunlun
// date: 2019-01-09
// description:
package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex //sync
	isClosed  bool
}

// init connection
func InitConn(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{wsConnect: wsConn, inChan: make(chan []byte, 1024), outChan: make(chan []byte, 1024), closeChan: make(chan byte, 1)}
	go conn.readLoop()
	go conn.writeLoop()
	return
}

// read message
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
		fmt.Println("inChan message ...")
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// write message
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
		fmt.Println("outChan message > > > time:", time.Now().UnixNano())
	case <-conn.closeChan:
		err = errors.New("connection is closed")

	}
	return
}

// close connection
func (conn *Connection) Close() {
	conn.wsConnect.Close()
	// sync
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// loop read message
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}

		//sync until write message into inChan
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR
		}
	}

ERR:
	conn.Close()
}

// loop write message
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}
