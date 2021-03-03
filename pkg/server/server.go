package server

// Interface Chat Server
type Interface interface {
	Connect(params ConnectParams, onConnect func(ch <-chan byte))
	Close(onClose func(ch <-chan byte))
	Send(nick string, msg string)

	OnConnect(ch <-chan byte)
	OnClose(ch <-chan byte)
	OnMessage(ch chan<- string)
}

// ConnectParams connection parameters
type ConnectParams struct {
	Name string
}
