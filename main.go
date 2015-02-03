package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

type EventBase struct {
	Type string
}

type EventMessage struct {
	Type    string
	Channel string
	User    string
	Text    string
	Ts      string
}

const BUF_SIZE = 8192

func handleSlackEvents(conn *websocket.Conn, eventsChannel chan string) {
	buf := make([]byte, BUF_SIZE)
	numBytes, err := conn.Read(buf)
	//	websocket.JSON.Receive(conn, &evt)
	//	fmt.Printf("Message received, type: %s\n", evt.Type)
	fmt.Printf("-----\n")
	fmt.Printf("Message received, %s\n", string(buf))
	fmt.Printf("Bytes Read: %d\n", numBytes)
	fmt.Printf("Error: %v\n", err)
}

func handleEvents(conn *websocket.Conn, incomingEvents chan string) {
	var in []byte

	fmt.Printf("handleEvents()\n")

	for {
		err := websocket.Message.Receive(conn, &in)

		if err != nil {
			fmt.Printf("handleEvents() exiting, error: %s\n", err)
			incomingEvents <- "error"
			return
		}

		incomingEvents <- string(in)
	}
}

func main() {
	apiToken := "xoxb-3593752108-geF16kipFlYuB0kbFFoQNjAR"
	resp, err := http.Get(fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", apiToken))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	type RtmStartResponse struct {
		Ok  bool
		Url string
	}
	var r RtmStartResponse

	je := json.Unmarshal(body, &r)

	incomingEvents := make(chan string)
	conn, _ := websocket.Dial(r.Url, "", "http://localhost")

	go handleEvents(conn, incomingEvents)

	for {
		m := <-incomingEvents
		if m == "error" {
			fmt.Printf("reconnecting...\n")
			conn, _ := websocket.Dial(r.Url, "", "http://localhost")
			go handleEvents(conn, incomingEvents)
		} else {
			fmt.Printf("Incoming message: %s\n", m)
		}
	}

	fmt.Printf("Body: %v\n", r)
	fmt.Printf("Error: %v\n", je)
}
