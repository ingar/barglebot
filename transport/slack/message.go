package slack

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
	"strings"
)

type Message struct {
	conn    *websocket.Conn
	payload map[string]interface{}
}

func (self Message) Text() string {
	return self.payload["text"].(string)
}

func (self Message) Tokens() []string {
	return strings.Split(self.Text(), " ")
}

func (self Message) Sender() string {
	userId := self.payload["user"].(string)
	user, _ := FindUserById(userId)
	return user.Name
}

func (self Message) Args() []string {
	return self.Tokens()[1:]
}

func (self Message) Respond(text string) {
	self.conn.WriteJSON(map[string]string{
		"id":      fmt.Sprintf("%v", time.Now().UnixNano()),
		"type":    "message",
		"channel": self.payload["channel"].(string),
		"text":    text,
	})
}

func (self Message) DebugDump() string {
	return fmt.Sprintf("Message: %v", self.payload)
}
