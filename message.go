/*
A Bot Framework

Example:

	package main

	import (
		"fmt"
		"github.com/ingar/barglebot"
		"github.com/ingar/barglebot/transport/slack"
		"os"
	)

	func main() {
		incomingMessages := make(chan barglebot.Message)
		slack.Connect(os.Getenv("SLACK_BOT_API_KEY"), incomingMessages)

		for {
			message := <-incomingMessages
			message.Respond(fmt.Sprintf("You said: '%s'", message.Text()))
		}
	}
*/
package barglebot

type Message interface {
	Text() string
	Respond(string)
	DebugDump() string
}
