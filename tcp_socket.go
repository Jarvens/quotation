package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	fmt.Println("server starting ...")
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept Fatal error: ", err.Error())
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error: ", err.Error())
			return
		}

		fmt.Println("Read message: ", strings.TrimSpace(string(buf)))
	}

}
