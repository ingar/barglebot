package main

import (
	"fmt"
	"superslackr/transport/slack"
	"superslackr/util"
)

func main() {
	incomingCommands :=  make(chan interface{})
	go slack.Connect("xoxb-3593752108-geF16kipFlYuB0kbFFoQNjAR", incomingCommands)
	for {
		c := <-incomingCommands
		fmt.Printf("\nIncoming message:\n")
		util.DumpJSON(&c)
	}
}
