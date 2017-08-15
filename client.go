package main

import (
	"container/list"
	"fmt"

	"github.com/gorilla/websocket"
)

// Client struct represents a single player.
type Client struct {
	conn *websocket.Conn
	x    int
	y    int
	game *Game
}

func (client *Client) loop(element *list.Element) {
	conn := client.conn
	defer client.close(element)

	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			break
		}

		client.parseClientMessage(string(p))
	}
}

func (client *Client) close(element *list.Element) {
	client.game.clients.Remove(element)
	client.conn.Close()
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

// func (client Client) getMapForClient(gameMap [][]int) [][]int {
// 	x := client.x
// 	y := client.y

// 	if x
// }
