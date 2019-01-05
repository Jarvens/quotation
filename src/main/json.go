// auth: kunlun
// date: 2019-01-06
// json example

package main

import (
	"encoding/json"
	"fmt"
)

type header struct {
	Encryption  string `json:"encryption"`
	Timestamp   int64  `json:"timestamp"`
	Key         string `json:"key"`
	Parthercode int    `json:"parthercode"`
}

func main() {

	headerObj := header{}

	headerString := `{"encryption":"md5","timestamp":1482463793,"key":"325436534","parthercode":1002}`
	_ = json.Unmarshal([]byte(headerString), &headerObj)
	fmt.Println("json ", headerObj)
}
