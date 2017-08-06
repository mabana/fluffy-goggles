package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{}
var game = Game{make([]Client, 0, 100)}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "index", nil)
}

func clientLoop(conn *websocket.Conn) {
	defer conn.Close()

	for {
		/*
			messageType, p, err := conn.ReadMessage()

			if err != nil {
				return
			}

			if err := conn.WriteMessage(messageType, p); err != nil {
				return
			}
		*/
		if _, _, err := conn.NextReader(); err != nil {
			break
		}
	}
}

func Wss(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Upgrade: ", err)
		return
	}

	game.RegisterClient(conn)
	go clientLoop(conn)
}

func getPort() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return (":" + port)
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/wss", Wss)
	router.ServeFiles("/assets/*filepath", http.Dir("public"))

	log.Fatal(http.ListenAndServe(getPort(), router))
}
