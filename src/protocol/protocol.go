package protocol

/**
* @auth    kunlun
* @date    2019-01-04 09:27
* @version v1.0
* @des     描述：协议
*
**/

//base protocol
type TcpProtocol struct {
	Code        byte   `json:"code"`
	Version     byte   `json:"version"`
	Header      string `json:"header"`
	RequestType byte   `json:"requestType"`
	ClientType  byte   `json:"clientType"`
	Content     []byte `json:"content"`
	HeaderLen   int16  `json:"headerLen"`
	ContentLen  int16  `json:"contentLen"`
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
)

func Instance() TcpProtocol {
	return TcpProtocol{Header: Header, Version: Version}
}
