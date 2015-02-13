package slack

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Message struct {
	conn    *websocket.Conn
	payload map[string]interface{}
}

func (self Message) Text() string {
	return self.payload["text"].(string)
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
