// auth: kunlun
// date: 2019-01-08
// description:
package main

import (
	"codec"
	"fmt"
)

func main() {
	var val uint32
	val = 12
	fmt.Println([]byte{byte(val), byte(val >> 8), byte(val >> 16), byte(val >> 24)})

	byteArr := []byte{12, 0, 0, 0}
	fmt.Println(codec.ByteToUint16(byteArr))

}
