package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "index", nil)
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
