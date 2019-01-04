/**
* @auth    kunlun
* @date    2019-01-04 17:15
* @version v1.0
* @des     描述：行情 socket 客户端
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
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Connection Fatal error: ", err.Error())
		return
	}

	defer conn.Close()

	for {
		content := protocol.Student{Name: "张三", Age: 18}
		contentStr, _ := json.Marshal(content)
		log.Debug("Json: %s", string(contentStr))

		time.Sleep(2000 * time.Millisecond)
		log.Debug("开始写入消息")
		_, err := conn.Write(codec.Encode([]byte(contentStr), 0x1, 0x1))
		if err != nil {
			return
		}

	}

}
