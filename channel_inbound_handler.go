package main

//  tcp channel interface
type ChannelInboundHandler interface {

	//register
	ChannelRegistered()

	//unRegistered
	ChannelUnRegistered()

	//inActive
	ChannelInActive()

	//Active
	ChannelActive()

	//channel add
	ChannelAdded()

	//read message
	ChannelRead()
}
