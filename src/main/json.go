// auth: kunlun
// date: 2019-01-06
// json example

package main

import (
	"codec"
	"encoding/hex"
	"fmt"
)

func main() {

	var data uint16 = 140
	fmt.Println(hex.EncodeToString(codec.Uint16ToByte(data)))
}
