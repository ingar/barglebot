# barglebot

Quickstart (main.go):

```
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
```
	
