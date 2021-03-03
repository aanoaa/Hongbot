package main

import (
	"github.com/aanoaa/hongbot/pkg/bot"
	"github.com/aanoaa/hongbot/pkg/server/shell"
)

func main() {
	server := &shell.Shell{}
	bot := bot.NewBot("hongbot", server)
	bot.Run()
}
