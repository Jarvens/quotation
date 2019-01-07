// auth: kunlun
// date: 2019-01-07
package protocol

//quote protocol
type QuoteProtocol struct {
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

//request
type BaseReq struct {
}

const (
	Header  = "Quotation"
	Version = 0x1
	Magic   = "magicNyxV0.1"
)

func NewProtocol() TcpProtocol {
	return TcpProtocol{Header: Header, Version: Version}
}
