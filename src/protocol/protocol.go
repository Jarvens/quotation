package protocol

type TcpProtocol struct {
	Code        byte
	Version     byte
	Header      string
	RequestType byte
	ClientType  byte
	Content     string
	HeaderLen   int16
	ContentLen  int16
}
