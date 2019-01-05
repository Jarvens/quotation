// Codec
// copyright @2019
package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

/**
*
* |  Header   | HeaderLen | Version | RequestType | ClientType | Content | ContentLen|  Crc  |
* -------------------------------------------------------------------------------------------
* | quotation |     9     |   0x1   |     0x1     |    0x1     |   abc   |     3     |  34354|
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
func IntToBytes(val int) []byte {
	value := int32(val)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, value)
	return byteBuffer.Bytes()
}

// encode message
func Encode(message []byte, request byte, client byte) []byte {
	return append(append(append(append([]byte(Header),
		Uint16ToByte(HeaderLen)...),
		byte(Version),
		byte(request),
		byte(client)),
		Uint16ToByte(uint16(len(message)))...),
		message...)
}

//decode  message
func Decode(buffer []byte, ch chan []byte) []byte {
	var i int
	len := len(buffer)
	for i = 0; i < len; i++ {
		if len < HeaderLen {
			break
		}
		//check header
		if string(buffer[i:HeaderLen]) == Header {
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

//byte to hex
func ByteToHex(buffer []byte) string {
	return hex.EncodeToString(buffer)
}

//uint16 to byte
func Uint16ToByte(val uint16) []byte {
	return []byte{byte(val), byte(val >> 8)}

}

// hex string to byte
func HexToByte(str string) []byte {
	value, _ := hex.DecodeString(str)
	return value
}
