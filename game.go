package main

import (
	"container/list"
	"time"

	"github.com/gorilla/websocket"
)

// Game struct is the main struct in the game
type Game struct {
	clients *list.List
	gameMap [][]int
}

// RegisterClient should be used when we want to add new player.
func (g *Game) RegisterClient(conn *websocket.Conn) {
	client := &Client{conn, 5, 5, game}
	element := g.clients.PushBack(client)

	go client.loop(element)
}

func (g *Game) getMapWithClients() [][]int {
	currentMap := make([][]int, len(g.gameMap))

	for i, row := range g.gameMap {
		currentMap[i] = make([]int, len(row))
		copy(currentMap[i], row)
	}

	for e := g.clients.Front(); e != nil; e = e.Next() {
		client, ok := e.Value.(*Client)

		if ok {
			currentMap[client.y][client.x] = 2
		}
	}

	return currentMap
}

// ServerUpdateLoop function we use to send informations to clients
func (g *Game) ServerUpdateLoop() {
	for {
		for e := g.clients.Front(); e != nil; e = e.Next() {
			client, ok := e.Value.(*Client)

			if ok {
				err := client.conn.WriteJSON(g.getMapWithClients())

				if err != nil {
					panic(err)
				}
			}
		}

		time.Sleep(45 * time.Millisecond)
	}
}
