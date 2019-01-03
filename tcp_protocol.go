package main

type tcpProtocol struct {
	Code        byte
	Version     byte
	Header      string
	RequestType byte
	ClientType  byte
	Content     string
	HeaderLen   int16
	ContentLen  int16
}

const (
	Header  = "Quotation"
	Version = 0x1
)

func Instance() tcpProtocol {
	return tcpProtocol{Header: Header, Version: Version}
}
