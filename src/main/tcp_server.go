package main

import (
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
	for {
		buf := make([]byte, 1024)
		log.Debug("BUF 地址：%p", &buf)
		_, err := conn.Read(buf)
		if err != nil {
			log.Debug("Server Read message error: ", err.Error())
			return
		}
		log.Debug("Read message: %s", strings.TrimSpace(string(buf[:8])))
	}

}

//init
func init() {
	log.LoadConfiguration("log4go.xml")
}
