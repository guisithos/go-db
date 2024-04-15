package main

import (
	"github.com/guisithos/go-db/database"
)

func main() {

	_, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
}
