# hongbot

chatbot such as hubot

## Shell as chatserver

``` go
// main.go
package main

import (
    "github.com/aanoaa/hongbot/pkg/bot"
    "github.com/aanoaa/hongbot/pkg/server/shell"
)

func main() {
    shell := &shell.Shell{}
    bot := bot.NewBot(server.ConnectParams{Name: "hongbot"}, shell)
    bot.Run()
}
```

## IRC as chatserver

``` go
// main.go
package main

import (
    "github.com/aanoaa/hongbot/pkg/bot"
    "github.com/aanoaa/hongbot/pkg/server"
    "github.com/aanoaa/hongbot/pkg/server/irc"
)

func main() {
    irc := &irc.Irc{}
    bot := bot.NewBot(server.ConnectParams{
        Name:     "hongbot",
        User:     "hongbot",
        Channels: []string{"#mychannel"},
        Url:      "localhost:6668",
    }, irc)
    bot.Run()
}
```

```
$ go run main.go
```
