// auth: kunlun
// date: 2019-01-07
// description: 行情服务
package server

import (
	"codec"
	"domain"
	"encoding/json"
	"net"
	"time"
	log "utils"
)

func Start() {

	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Info("Listener error: %v", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Info("Accept error: %v", err.Error())
			continue
		}
		go loopHandler(conn)
	}

}

//initial load config
func init() {
	log.Debug("tcp server initial func: %v", time.Now())
}

//
func loopHandler(conn net.Conn) {

	defer conn.Close()
	tmpBuffer := make([]byte, 0)
	readChan := make(chan []byte, 16)
	go readData(readChan)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {

			//TODO 处理tcp主动断开
			log.Info("read message error: %v", err.Error())
			return
		}

		tmpBuffer = codec.QuoteDecode(buffer[:n], readChan)
		obj := domain.ResponseData{}
		_ = json.Unmarshal(tmpBuffer, &obj)
		log.Info("receive message: %v", string(tmpBuffer))
	}
}

func readData(ch chan []byte) {
	for {
		select {
		case data := <-ch:
			log.Info("写入数据: %v", string(data))
		}
	}
}
