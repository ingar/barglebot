package slack

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ingar/barglebot"
	"github.com/ingar/barglebot/util"
)

type BotConnection struct {
	url       string
	botUserId string
	messages  chan barglebot.Message
	conn      *websocket.Conn
}

func debug(s string) {
	fmt.Println("[SLACK]", s)
}

func (self BotConnection) handleEvents() {
	debug(fmt.Sprintf("Dialing %s\n", self.url))
	conn, _, _ := websocket.DefaultDialer.Dial(self.url, nil)
	self.conn = conn

	for {
		var o interface{}
		conn.ReadJSON(&o)
		self.handleEvent(o.(map[string]interface{}))
	}
}

func (self BotConnection) handleEvent(o map[string]interface{}) {
	debug(fmt.Sprintf("Received event:\n%s", util.FmtJSON(o)))

	if _, ok := o["subtype"]; !ok && o["type"] == "message" && self.botUserId != o["user"] {
		m := Message{self.conn, o}
		self.messages <- m
		debug("Forwarding on to bot.")
	} else {
		debug("Invalid event for bot consumption.")
	}
}

func Connect(apiToken string, incomingEvents chan barglebot.Message) (users []User) {
	apiCaller := RestAPICaller{apiToken}
	url, userId, users := authenticate(apiCaller)
	connection := BotConnection{url, userId, incomingEvents, nil}
	go connection.handleEvents()
	return
}
