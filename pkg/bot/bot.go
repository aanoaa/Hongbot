package bot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aanoaa/hongbot/pkg/server"
)

// Interface bot
type Interface interface {
	Run()
	Shutdown()

	Send(msg string)
	Reply(nick string, msg string)

	Hear(pattern string, cb func(msg string))
	Respond(pattern string, cb func(nick string, msg string))
}

// Bot chatbot
type Bot struct {
	Name   string
	Server server.Interface

	Reaction map[*regexp.Regexp]func(msg string)
	Response map[*regexp.Regexp]func(nick string, msg string)
}

func NewBot(name string, server server.Interface) *Bot {
	bot := &Bot{
		Name:     name,
		Server:   server,
		Reaction: make(map[*regexp.Regexp]func(string)),
		Response: make(map[*regexp.Regexp]func(string, string)),
	}
	bot.Hear("ping", func(msg string) { bot.Send("pong") })
	bot.Respond("ping", func(nick string, msg string) { bot.Reply(nick, "pong") })
	bot.Respond("shutdown", func(nick string, msg string) {
		bot.Send("Goodbye cruel world")
		bot.Shutdown()
	})
	return bot
}

// Run power on
func (b *Bot) Run() {
	params := server.ConnectParams{Name: b.Name}
	b.Server.Connect(params, b.Server.OnConnect)

	ch := make(chan string)
	go func() {
		defer close(ch)
		b.Server.OnMessage(ch)
	}()

	called, _ := regexp.Compile(fmt.Sprintf(`^%s:`, b.Name))
	for msg := range ch {
		msg = strings.Trim(msg, "\n")
		nick := ""
		for i, c := range msg {
			if c == '>' {
				nick = msg[:i]
				msg = msg[i+1:]
				break
			}
		}

		if called.MatchString(msg) {
			for pattern := range b.Response {
				if pattern.MatchString(msg) {
					b.Response[pattern](nick, msg)
				}
			}
		} else {
			for pattern := range b.Reaction {
				if pattern.MatchString(msg) {
					b.Reaction[pattern](msg)
				}
			}
		}
	}
}

// Shutdown power off
func (b *Bot) Shutdown() {
	b.Server.Close(b.Server.OnClose)
}

// Send message
func (b *Bot) Send(msg string) {
	b.Server.Send(b.Name, msg)
}

// Reply message
func (b *Bot) Reply(nick string, msg string) {
	b.Server.Send(b.Name, fmt.Sprintf("%s: %s", nick, msg))
}

// Hear pattern and register callback
// hearing <pattern> message
// foo> ping
func (b *Bot) Hear(pattern string, cb func(msg string)) {
	compiled, _ := regexp.Compile(pattern)
	b.Reaction[compiled] = cb
}

// Respond pattern and register callback
// someone call bot with listening <pattern>
// foo> bot: ping
func (b *Bot) Respond(pattern string, cb func(nick string, msg string)) {
	compiled, _ := regexp.Compile(pattern)
	b.Response[compiled] = cb
}
