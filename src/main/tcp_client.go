/**
* @auth    kunlun
* @date    2019-01-04 17:15
* @version v1.0
* @des     描述：行情 socket 客户端
*
**/
package main

import (
	"encoding/json"
	"fmt"
	log "github.com/alecthomas/log4go"
	"net"
	"protocol"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Connection Fatal error: ", err.Error())
		return
	}

	defer conn.Close()

	pro := protocol.TcpProtocol{Code: 0x1, Version: protocol.Version, Header: protocol.Header, RequestType: 0x3, ClientType: 0x4}
	for {
		content := protocol.Student{Name: "张三", Age: 18}
		contentStr, _ := json.Marshal(content)
		pro.Content = contentStr
		log.Debug("Json: %s", string(contentStr))
		//timeInput := strings.Trim(pro, "\r\n")
		//if timeInput == "q" {
		//	return
		//}
		time.Sleep(2 * time.Millisecond)
		log.Debug("开始写入消息")
		_, err := conn.Write([]byte("Hello"))
		if err != nil {
			return
		}

	}

}
