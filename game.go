package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

type Game struct {
	clients []Client
}

func (g *Game) RegisterClient(conn *websocket.Conn) {
	g.clients = append(g.clients, Client{conn})
}
