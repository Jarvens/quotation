/**
* @auth    kunlun
* @date    2019-01-04 22:48
* @version v1.0
* @des     描述：
*
**/
package utils

import (
	"fmt"
	"testing"
)

//test pass
func Test_Md5Byte(t *testing.T) {
	message := "hello"
	hash := Md5Byte([]byte(message))
	md5Str := fmt.Sprintf("%x", hash)
	fmt.Println("MD5值: ", md5Str)

}

//test pass
func TestCrc32(t *testing.T) {
	message := "crc32"
	var val uint32 = Crc32([]byte(message))
	fmt.Println("Crc32: ", val)
}
