package server

// Interface Chat Server
type Interface interface {
	Connect(params ConnectParams, onConnect func())
	Close(onClose func())
	Send(channel, nick, msg string)
	OnMessage(ch chan<- string)
}

// ConnectParams connection parameters
type ConnectParams struct {
	// Name bot name
	// ex) hongbot
	Name string

	// User username
	User string

	// Channels to participate in after connecting
	// ex) #example
	Channels []string

	// Url server url
	// ex) irc.freenode.net:6667
	Url string
}
