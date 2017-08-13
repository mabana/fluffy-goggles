package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Client struct represents a single player.
type Client struct {
	conn *websocket.Conn
	x    int
	y    int
}

func (client *Client) loop() {
	conn := client.conn
	defer conn.Close()

	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			break
		}

		client.parseClientMessage(string(p))
	}
}

func (client *Client) parseClientMessage(msg string) {
	switch msg {
	case "move:left":
		client.x--
	case "move:up":
		client.y--
	case "move:right":
		client.x++
	case "move:down":
		client.y++
	default:
		fmt.Println("Other: ", msg)
	}
}
