// auth: kunlun
// date: 2019-01-07
package codec

//quote protocol
type BaseProto struct {
	Magic string `json:"magic"`
	Len   int16  `json:"len"`
	Data  string `json:"data"`
	Crc32 int32  `json:"crc32"`
}

//push protocol
type TcpProtocol struct {
	Code        byte   `json:"code"`
	Version     byte   `json:"version"`
	Header      string `json:"header"`
	RequestType byte   `json:"requestType"`
	ClientType  byte   `json:"clientType"`
	Content     string `json:"content"`
	HeaderLen   int16  `json:"headerLen"`
	ContentLen  int16  `json:"contentLen"`
	Crc         int32  `json:"crc"`
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
