package database

import (
	"database/sql"
	"fmt"

	go_ora "github.com/sijms/go-ora/v2"
)

func ConnectToDatabase() (*sql.DB, error) {
	user := ""
	password := ""
	host := ""
	port := 1521
	serviceName := "dbtreinamentos"

	connStr := go_ora.BuildUrl(host, port, serviceName, user, password, nil)

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection to Oracle database successful!")

	return db, nil
}
