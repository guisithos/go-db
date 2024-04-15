package database

import (
	"database/sql"
	"fmt"

	"github.com/guisithos/go-db/env"
	go_ora "github.com/sijms/go-ora/v2"
)

func ConnectToDatabase() (*sql.DB, error) {
	user := env.LoadEnviroment("DB_USER")
	password := env.LoadEnviroment("DB_PASSWORD")
	host := env.LoadEnviroment("DB_HOST")
	port := env.LoadEnviroment("DB_PORT")
	serviceName := env.LoadEnviroment("DB_SERVICE_NAME")

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
