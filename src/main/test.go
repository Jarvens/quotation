package main

import (
	"codec"
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/jinzhu/configor"
)

func main() {

	var Config = struct {
		AppName string `default:"app name"`
		Socket  struct {
			Port int
		}
	}{}

	configor.Load(&Config, "config.yml")
	fmt.Printf("Config properties: %#v", Config)
	log4go.LoadConfiguration("log4go.xml")
	log4go.Debug("This is quotation log4go \r\n")
	defer log4go.Close()

	var message []byte = append(append([]byte("header"), (codec.IntToBytes(len("header")))...), "hahaha"...)
	var len int = codec.BytesToInt(message[6:10])
	fmt.Println("长度：", len)
	fmt.Println(string(message))

}
