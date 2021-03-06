package main

import (
	"flag"

	"github.com/aanoaa/hongbot/pkg/bot"
	"github.com/aanoaa/hongbot/pkg/server"
	"github.com/aanoaa/hongbot/pkg/server/irc"
	"github.com/aanoaa/hongbot/pkg/server/shell"
)

func main() {
	name := flag.String("name", "hongbot", "bot name")
	serv := flag.String("server", "shell", "server url")
	channel := flag.String("channel", "#shell", "Channels to participate in after connecting")
	flag.Parse()

	var chatbot *bot.Bot
	var chatServer server.Interface
	var params server.ConnectParams
	if *serv != "shell" {
		chatServer = &irc.Irc{}
		params = server.ConnectParams{
			Name:     *name,
			User:     *name,
			Channels: []string{*channel},
			Url:      *serv,
		}
	} else {
		chatServer = &shell.Shell{}
		params = server.ConnectParams{
			Name: *name,
		}
	}

	chatbot = bot.NewBot(params, chatServer)
	chatbot.Run()
}
