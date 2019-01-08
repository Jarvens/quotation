// Codec
// copyright @2019
package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"utils"
)

/*

|    magic   |  len  |  data  |  crc32  |
------------------------------------------
|magicNyxV0.1|   3   |   ABC  | 4354356 |

*/
const (
	OrderHeader    = "quotation"
	HeaderLen      = 9
	Version        = 0x1
	QuoteHeader    = "magicNyxV0.1"
	QuoteHeaderLen = 12
)

// byte to int  bigEndian pattern
func BytesToInt(b []byte) int {
	byteBuffer := bytes.NewBuffer(b)
	var value int32
	binary.Read(byteBuffer, binary.BigEndian, &value)
	return int(value)
}

// byte to uint16
func ByteToUint16(b []byte) uint16 {
	var val uint16 = 0
	for i := 0; i < len(b); i++ {
		val = val + uint16(uint(b[i])<<uint(8*i))
	}
	return val
}

// int to byte  bigEndian pattern
func IntToBytes(val int) []byte {
	value := int32(val)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, value)
	return byteBuffer.Bytes()
}

// quotation protocol encode
func QuoteEncode(message []byte) []byte {
	crcVal := utils.Crc32(message)
	return append(append(append([]byte(QuoteHeader), Uint16ToByte(uint16(len(message)))...), message...), Uint32ToByte(crcVal)...)
}

// encode message
func Encode(message []byte, request byte, client byte) []byte {
	return append(append(append(append([]byte(OrderHeader),
		Uint16ToByte(HeaderLen)...),
		byte(Version),
		byte(request),
		byte(client)),
		Uint16ToByte(uint16(len(message)))...),
		message...)
}

// quotation protocol decode
func QuoteDecode(buffer []byte, ch chan []byte) []byte {
	var i int
	len := len(buffer)
	var messageLen uint16
	//fmt.Println("comming len: ", len)
	for i = 0; i < len; i++ {
		if len < int(QuoteHeaderLen) {
			break
		}
		//check header
		if string(buffer[i:QuoteHeaderLen]) == QuoteHeader {
			messageLen = ByteToUint16(buffer[i+QuoteHeaderLen : i+QuoteHeaderLen+2])

			//fmt.Println("now len: ", messageLen)
			if len <= i+QuoteHeaderLen+2+int(messageLen)+4 {
				break
			}
			data := buffer[i+QuoteHeaderLen : i+QuoteHeaderLen+int(messageLen)]
			ch <- data
			i += QuoteHeaderLen + int(messageLen) - 1
		}
	}
	if i == len {
		return make([]byte, 0)
	}
	return buffer[i+QuoteHeaderLen+2 : i+QuoteHeaderLen+2+int(messageLen)]
}

// decode  message
func Decode(buffer []byte, ch chan []byte) []byte {
	var i int
	len := len(buffer)
	for i = 0; i < len; i++ {
		if len < HeaderLen {
			break
		}
		//check header
		if string(buffer[i:HeaderLen]) == OrderHeader {
			messageLen := BytesToInt(buffer[i+HeaderLen : i+HeaderLen+4])
			if len < i+HeaderLen+4+messageLen {
				break
			}
			data := buffer[i+HeaderLen+4 : i+HeaderLen+4+messageLen]
			ch <- data
			i += HeaderLen + 4 + messageLen - 1
		}
	}
	if i == len {
		return make([]byte, 0)
	}

	return buffer[i:]
}

// byte to hex
func ByteToHex(buffer []byte) string {
	return hex.EncodeToString(buffer)
}

// uint16 to byte
func Uint16ToByte(val uint16) []byte {
	return []byte{byte(val), byte(val >> 8)}

}

// uint32 to byte
func Uint32ToByte(val uint32) []byte {
	return []byte{byte(val), byte(val >> 8), byte(val >> 16), byte(val >> 24)}
}

// hex string to byte
func HexToByte(str string) []byte {
	value, _ := hex.DecodeString(str)
	return value
}

// byte to string
func ByteToString(b *[]byte) *string {
	str := bytes.NewBuffer(*b)
	result := str.String()
	return &result
}
