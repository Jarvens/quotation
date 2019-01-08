/**
* @auth    kunlun
* @date    2019-01-04 17:55
* @version v1.0
* @des     描述：codec 测试
*
**/
package codec

import (
	"domain"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_Encode(t *testing.T) {
	buffer := Encode([]byte("test message protocol"), 0x1, 0x1)
	hexString := ByteToHex(buffer)
	fmt.Println("encode protocol Hex：", hexString)

	ch := make(chan []byte, 16)
	bytes := Decode(buffer, ch)
	fmt.Println("decode : ", string(bytes))
}

func TestQuoteEncode(t *testing.T) {
	data := domain.ResponseData{Dir: "bid", Symbol: "USDT_BTC", Ts: time.Now().UnixNano(), Amount: 0.2, Price: 0.1, DayVolume: 10, DayPrice: 0.5, DayHigh: 0.5, DayLow: 0.2}
	result, _ := json.Marshal(data)
	fmt.Printf("encode json: %v\r\n", string(result))
	var resultBytes []byte
	resultBytes = QuoteEncode(result)

	readChan := make(chan []byte, 16)
	byte1 := QuoteDecode(resultBytes, readChan)
	var pro = domain.ResponseData{}
	json.Unmarshal(byte1, &pro)
	fmt.Printf("decode object: %v \r\n", pro)
	fmt.Println("decode result: ", string(byte1))

}

// benchmark test
// 200000 count invoke        5850 ns/op   1.3s   total
// go test -run=codec_test.go -bench=BenchmarkQuoteDecode
func BenchmarkQuoteDecode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := domain.ResponseData{Dir: "bid", Symbol: "USDT_BTC", Ts: time.Now().UnixNano(), Amount: 0.2, Price: 0.1, DayVolume: 10, DayPrice: 0.5, DayHigh: 0.5, DayLow: 0.2}
		result, _ := json.Marshal(data)
		//fmt.Printf("encode json: %v\r\n", string(result))
		var resultBytes []byte
		resultBytes = QuoteEncode(result)
		//fmt.Println(resultBytes)
		readChan := make(chan []byte, 16)
		byte1 := QuoteDecode(resultBytes, readChan)
		var pro = domain.ResponseData{}
		json.Unmarshal(byte1, &pro)
		//fmt.Printf("decode object: %v \r\n", pro)
		//fmt.Println("decode result: ", string(byte1))
	}
}
