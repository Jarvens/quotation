/**
* @auth    kunlun
* @date    2019-01-04 17:55
* @version v1.0
* @des     描述：codec 测试
*
**/
package codec

import (
	"fmt"
	"testing"
)

func Test_Encode(t *testing.T) {
	buffer := Encode([]byte("这是测试协议消息体"), 0x1, 0x1)
	hexString := ByteToHex(buffer)
	fmt.Println("协议码：", hexString)
}
