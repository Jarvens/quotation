/**
* @auth    kunlun
* @date    2019-01-04 15:32
* @version v1.0
* @des     描述：编解码
*
**/
package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

/**
*
* |  Header   | HeaderLen | Code | Version | RequestType | ClientType | Content | ContentLen|
* -------------------------------------------------------------------------------------------
* | quotation |     9     |   0  |   0x1   |     0x1     |    0x1     | abcdefg |     8     |
*
*
**/

const (
	Header    = "quotation"
	HeaderLen = 9
	Version   = 0x1
)

//byte to int  bigEndian pattern
func BytesToInt(b []byte) int {
	byteBuffer := bytes.NewBuffer(b)
	var value int32
	binary.Read(byteBuffer, binary.BigEndian, &value)
	return int(value)
}

//int to byte  bigEndian pattern
func IntToBytes(i int) []byte {
	value := int32(i)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, value)
	return byteBuffer.Bytes()
}

// encode message
func Encode(message []byte) []byte {
	buffer := make([]byte, 0)
	buffer = append([]byte(Header))

	headerBytes := append([]byte(Header), IntToBytes(HeaderLen)...)

	return append(append([]byte(Header), IntToBytes(HeaderLen)...))
}

//decode  message
func Decode(buffer []byte, ch chan []byte) {

}

//byte to hex
func ByteToHex(buffer []byte) string {
	return hex.EncodeToString(buffer)
}

// hex string to byte
func HexToByte(str string) []byte {
	value, _ := hex.DecodeString(str)
	return value
}

//func Packet(message []byte) []byte {
//	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
//}
