package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Game")
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
	router.ServeFiles("/assets/*filepath", http.Dir("public"))

	log.Fatal(http.ListenAndServe(getPort(), router))
}
