package main

import (
	"container/list"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{}
var game *Game

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "index", nil)
}

func wss(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Upgrade: ", err)
		return
	}

	game.RegisterClient(conn)
}

func getPort() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	return (":" + port)
}

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/wss", wss)
	router.ServeFiles("/assets/*filepath", http.Dir("public"))

	gameMap := GenerateMap(17, 13)

	game = &Game{list.New(), gameMap}
	go game.ServerUpdateLoop()

	log.Fatal(http.ListenAndServe(getPort(), router))
}
