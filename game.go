package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Client struct is represents a single player.
type Client struct {
	conn *websocket.Conn
	x    int
	y    int
}

// Game struct is the main struct in the game
type Game struct {
	clients []Client
}

// RegisterClient should be used when we want to add new player.
func (g *Game) RegisterClient(conn *websocket.Conn) {
	client := Client{conn, 0, 0}
	g.clients = append(g.clients, client)
	go client.loop()
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
		client.y++
	case "move:right":
		client.x++
	case "move:down":
		client.y--
	default:
		fmt.Println("Other: ", msg)
	}
	fmt.Println(client.x, client.y)
}
