package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := ":8085"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//go 并发关键字
		go handle(conn)
	}

}

//错误检查
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//处理器
func handle(conn net.Conn) {

	//conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		//可读数据为0
		if read_len != 0 {
			var str = string(request[:read_len])
			fmt.Println("来自客户端的问候：", str)
			conn.Write([]byte([]byte("Hello goLang\r\n")))
			fmt.Println("开始回写客户端")
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)

			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
	}
	request = make([]byte, 128)
}
