package main

import (
	"fmt"
	"net"
	"strings"
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
		input := "Hello go"
		timeInput := strings.Trim(input, "\r\n")
		if timeInput == "q" {
			return
		}
		time.Sleep(2 * time.Second)
		_, err := conn.Write([]byte(timeInput))
		if err != nil {
			return
		}

	}

}
