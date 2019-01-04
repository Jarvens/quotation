/**
* @auth    kunlun
* @date    2019-01-04 17:15
* @version v1.0
* @des     描述：行情 socket 服务端
*
**/

package main

import (
	"codec"
	"encoding/json"
	"fmt"
	log "github.com/alecthomas/log4go"
	"net"
	"protocol"
	"strings"
)

func main() {

	var pro = protocol.TcpProtocol{Version: protocol.Version, Header: protocol.Header}
	fmt.Println("打印协议体：", pro)
	fmt.Println("server starting ...")
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Debug("Fatal error: ", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Debug("Accept Fatal error: ", err.Error())
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	//byte Buffer cache
	tmpBuffer := make([]byte, 0)
	readChan := make(chan []byte, 16)
	go readData(readChan)
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Debug("Server Read message error: ", err.Error())
			return
		}
		//tmpBuffer = codec.Decode(append(tmpBuffer, buffer[:n]...), readChan)
		tmpBuffer = codec.Decode(buffer[:n], readChan)
		obj := protocol.TcpProtocol{}

		_ = json.Unmarshal(tmpBuffer, &obj)
		log.Debug("Object", obj)
		log.Debug("Read message: %s", strings.TrimSpace(string(tmpBuffer)))
	}

}

//init
func init() {
	log.LoadConfiguration("log4go.xml")
}

//read  data
func readData(readChan chan []byte) {

	for {
		select {
		case data := <-readChan:
			log.Debug(string(data))
		}
	}
}

//write data to client
func writeData() {

}
