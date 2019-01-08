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
	"domain"
	"encoding/json"
	"fmt"
	"net"
	"time"
	log "utils"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Connection Fatal error: ", err.Error())
		return
	}

	defer conn.Close()

	for {

		data := domain.ResponseData{Dir: "bid", Symbol: "USDT_BTC", Ts: time.Now().UnixNano(), Amount: 0.2, Price: 0.1, DayVolume: 10, DayPrice: 0.5, DayHigh: 0.5, DayLow: 0.2}
		dataStr, _ := json.Marshal(data)
		_, err := conn.Write(codec.QuoteEncode(dataStr))
		log.Info("start send message time: %d ", time.Now().UnixNano())
		time.Sleep(5 * time.Millisecond)
		if err != nil {
			return
		}

	}

}
