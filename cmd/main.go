package main

import (
	"net/http"

	"github.com/guisithos/go-db/app"
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
