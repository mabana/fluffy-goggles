package main

import (
	"container/list"

	"github.com/gorilla/websocket"
)

// Game struct is the main struct in the game
type Game struct {
	clients *list.List
	gameMap [][]int // TODO: Game map should be pointer
}

type welcomeMessage struct {
	GameMap   *[][]int `json:"gameMap"`
	X         int      `json:"x"`
	Y         int      `json:"y"`
	MapWidth  int      `json:"mapWidth"`
	MapHeight int      `json:"mapHeight"`
}

// RegisterClient should be used when we want to add new player.
func (g *Game) RegisterClient(conn *websocket.Conn) {
	client := &Client{conn, 5, 5, game}
	element := g.clients.PushBack(client)

	g.sendGreetingData(client)
	go client.loop(element)
}

func (g *Game) sendGreetingData(client *Client) {
	msg := &welcomeMessage{
		GameMap:   &g.gameMap,
		X:         client.x,
		Y:         client.y,
		MapWidth:  MapWidth,
		MapHeight: MapHeight,
	}

	client.conn.WriteJSON(msg)
}
