package main

import (
	"net/http"

	app "github.com/guisithos/go-db/app/web/dadosProtocolo"
	"github.com/guisithos/go-db/database"
)

func main() {

	_, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", app.HandleWebpage)
	http.ListenAndServe(":8080", nil)
}
