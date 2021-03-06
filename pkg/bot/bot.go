package bot

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/aanoaa/hongbot/pkg/server"
)

// Interface bot
type Interface interface {
	Run()
	Listen()
	Shutdown()

	Send(channel, msg string)
	Reply(channel, nick, msg string)

	Hear(pattern string, cb func(channel, nick, msg string))
	Respond(pattern string, cb func(channel, nick, msg string))

	OnConnect()
	OnClose(ch <-chan byte)
}

// Bot chatbot
type Bot struct {
	Name   string
	Server server.Interface
	server.ConnectParams

	Reaction map[*regexp.Regexp]func(channel, nick, msg string)
	Response map[*regexp.Regexp]func(channel, nick, msg string)
}

// NewBot create bot instance
func NewBot(params server.ConnectParams, server server.Interface) *Bot {
	bot := &Bot{
		Name:          params.Name,
		Server:        server,
		ConnectParams: params,
		Reaction:      make(map[*regexp.Regexp]func(channel, nick, msg string)),
		Response:      make(map[*regexp.Regexp]func(channel, nick, msg string)),
	}
	bot.Hear("ping", func(channel, nick, msg string) { bot.Send(channel, "pong") })
	bot.Respond("ping", func(channel, nick, msg string) { bot.Reply(channel, nick, "pong") })
	bot.Respond("shutdown", func(channel, nick, msg string) {
		bot.Send(channel, "Goodbye cruel world")
		bot.Shutdown()
	})
	return bot
}

// Run connect to server and listen
func (b *Bot) Run() {
	b.Server.Connect(b.ConnectParams, b.OnConnect)
	b.Listen()
}

// Shutdown power off
func (b *Bot) Shutdown() {
	b.Server.Close(b.OnClose)
}

// Send message
func (b *Bot) Send(channel, msg string) {
	b.Server.Send(channel, b.Name, msg)
}

// Reply message
func (b *Bot) Reply(channel, nick, msg string) {
	b.Send(channel, fmt.Sprintf("%s: %s", nick, msg))
}

// Hear pattern and reaction if matched
// foo> ping
func (b *Bot) Hear(pattern string, cb func(channel, nick, msg string)) {
	compiled, _ := regexp.Compile(pattern)
	b.Reaction[compiled] = cb
}

// Respond pattern and response if matched
// foo> bot: ping
func (b *Bot) Respond(pattern string, cb func(channel, nick, msg string)) {
	compiled, _ := regexp.Compile(pattern)
	b.Response[compiled] = cb
}

// OnConnect connection accepted
func (b *Bot) OnConnect() {
	log.Println("Connection accepted")
	log.Println("<Ctrl + c> to quit")
}

// OnClose
func (b *Bot) OnClose() {
	log.Println("Connection closed")
}

// Listen listen server messages
func (b *Bot) Listen() {
	ch := make(chan string)
	go func() {
		defer close(ch)
		b.Server.OnMessage(ch)
	}()

	called, _ := regexp.Compile(fmt.Sprintf(`^%s:`, b.Name))
	for line := range ch {
		// #channel@nick>message
		line = strings.Trim(line, "\n")
		var channel, nick, message string
		nickIdx := 0
		for i, c := range line {
			switch c {
			case '@':
				channel = line[:i]
				nickIdx = i
			case '>':
				nick = line[nickIdx+1 : i]
				message = line[i+1:]
				break
			default:
			}
		}

		if called.MatchString(message) {
			for pattern := range b.Response {
				if pattern.MatchString(message) {
					b.Response[pattern](channel, nick, message)
				}
			}
		} else {
			for pattern := range b.Reaction {
				if pattern.MatchString(message) {
					b.Reaction[pattern](channel, nick, message)
				}
			}
		}
	}
}
