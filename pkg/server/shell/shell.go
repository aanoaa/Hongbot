package shell

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/aanoaa/hongbot/pkg/server"
)

// Shell Terminal shell
type Shell struct {
	connected bool
}

// Connect to shell
func (s *Shell) Connect(params server.ConnectParams, onConnect func()) {
	s.connected = true
	onConnect()
}

// OnMessage receive a message
func (s *Shell) OnMessage(ch chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for s.connected {
		time.AfterFunc(time.Second/100, func() {
			fmt.Printf("you> ")
		})
		input, _ := reader.ReadString('\n')
		ch <- fmt.Sprintf("you>%s", input)
	}
}

// Close shell
func (s *Shell) Close(onClose func()) {
	s.connected = false
	onClose()
}

// Send message
func (s *Shell) Send(nick string, msg string) {
	fmt.Printf("%s> %s\n", nick, msg)
}
