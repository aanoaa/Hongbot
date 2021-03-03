package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aanoaa/hongbot/pkg/server"
)

// Shell Terminal shell
type Shell struct {
	connected bool
}

// Connect to shell
func (s *Shell) Connect(params server.ConnectParams, onConnect func(ch <-chan byte)) {
	ch := make(chan byte)
	defer close(ch)
	go onConnect(ch)
	ch <- 1
}

// OnConnect connection accepted
func (s *Shell) OnConnect(ch <-chan byte) {
	<-ch
	s.connected = true
	log.Println("Connection accepted")
	log.Println("<Ctrl + c> to quit")
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
func (s *Shell) Close(onClose func(ch <-chan byte)) {
	ch := make(chan byte)
	defer close(ch)
	go onClose(ch)
	ch <- 1
}

func (s *Shell) OnClose(ch <-chan byte) {
	<-ch
	s.connected = false
	log.Println("Connection closed")
}

// Send message
func (s *Shell) Send(nick string, msg string) {
	fmt.Printf("%s> %s\n", nick, msg)
}
