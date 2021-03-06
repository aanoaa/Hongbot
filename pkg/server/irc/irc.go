package irc

import (
	"fmt"
	"log"

	"github.com/aanoaa/hongbot/pkg/server"
	irc "github.com/thoj/go-ircevent"
)

// Irc Server implementation
type Irc struct {
	Conn *irc.Connection
}

// Connect to IRC
func (i *Irc) Connect(params server.ConnectParams, onConnect func()) {
	conn := irc.IRC(params.Name, params.User)
	conn.UseTLS = false
	// Welcome
	conn.AddCallback("001", func(e *irc.Event) {
		for _, ch := range params.Channels {
			conn.Join(ch)
		}
	})

	i.Conn = conn
	err := i.Conn.Connect(params.Url)
	if err != nil {
		log.Fatal(err)
	}
	onConnect()
}

// OnMessage receive a message
func (i *Irc) OnMessage(ch chan<- string) {
	i.Conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		ch <- fmt.Sprintf("%s@%s>%s", e.Arguments[0], e.Nick, e.Message())
	})
	i.Conn.Loop()
}

// Close disconnect IRC
func (i *Irc) Close(onClose func()) {
	i.Conn.Quit()
	onClose()
}

// Send message
func (i *Irc) Send(channel, nick, msg string) {
	i.Conn.Privmsg(channel, msg)
}
