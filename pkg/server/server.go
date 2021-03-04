package server

// Interface Chat Server
type Interface interface {
	Connect(params ConnectParams, onConnect func())
	Close(onClose func())
	Send(nick string, msg string)

	OnMessage(ch chan<- string)
}

// ConnectParams connection parameters
type ConnectParams struct {
	Name string
}
