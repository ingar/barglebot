package slack

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func Connect(apiToken string, incomingEvents chan interface{}) {
	apiCaller := RestAPICaller{apiToken}
	url := authenticate(apiCaller)

	fmt.Printf("Dialing %s\n", url)
	conn, _, _ := websocket.DefaultDialer.Dial(url, nil)

	var f interface{}
	for {
		conn.ReadJSON(&f)
		incomingEvents <- f
	}
}

