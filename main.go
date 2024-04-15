package main

import (
	"/home/gui/go-db/database"
)

func main() {

	_, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
}
