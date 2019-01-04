/**
* @auth    kunlun
* @date    2019-01-04 22:32
* @version v1.0
* @des     描述：Hash
*
**/
package hash

import (
	"crypto/md5"
)

const (
	Md5Salt = "MD5Salt"
)

//get crc32 hash
func Crc32(bytes []byte) int32 {

	return 0
}

//check crc32 sum
func CheckSum(src, dest int32) bool {
	if src == dest {
		return true
	}
	return false
}

//md5
func Md5Byte(bytes []byte) [16]byte {
	bytes = append([]byte(Md5Salt))
	return md5.Sum(bytes)
}

//md5
func Md5String(message string) [16]byte {
	message = message + Md5Salt
	return md5.Sum([]byte(message))
}
