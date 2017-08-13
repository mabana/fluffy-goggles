package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// Client struct represents a single player.
type Client struct {
	conn *websocket.Conn
	x    int
	y    int
}

// Game struct is the main struct in the game
type Game struct {
	clients []Client
	gameMap [][]int
}

// RegisterClient should be used when we want to add new player.
func (g *Game) RegisterClient(conn *websocket.Conn) {
	client := Client{conn, 5, 5}
	g.clients = append(g.clients, client)
	go g.clients[len(g.clients)-1].loop()
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

func (g *Game) getMapWithClients() [][]int {
	currentMap := make([][]int, len(g.gameMap))

	for i, row := range g.gameMap {
		currentMap[i] = make([]int, len(row))
		copy(currentMap[i], row)
	}

	for _, client := range g.clients {
		currentMap[client.y][client.x] = 2
	}

	return currentMap
}

// ServerUpdateLoop function we use to send informations to clients
func (g *Game) ServerUpdateLoop() {
	for {
		for _, client := range g.clients {
			err := client.conn.WriteJSON(g.getMapWithClients())

			if err != nil {
				panic(err)
			}
		}

		time.Sleep(45 * time.Millisecond)
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
