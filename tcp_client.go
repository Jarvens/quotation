package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

//tcp 客户端
func main() {

	//if len(os.Args) != 2 {
	//	fmt.Fprint(os.Stderr, "Usage: %s host:port ", os.Args[0])
	//	os.Exit(1)
	//}

	//service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8085")

	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	for i := 0; i < 5; i++ {
		_, err = conn.Write([]byte("Hello goLang\r\n"))
		checkErr(err)
	}
	result, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println("来自服务端的问候：", string(result))
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
