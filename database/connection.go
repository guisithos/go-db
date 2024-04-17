package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	go_ora "github.com/sijms/go-ora/v2"
)

func ConnectToDatabase() (*sql.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	serviceName := os.Getenv("DB_SERVICE_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	connStr := go_ora.BuildUrl(host, port, serviceName, user, password, nil)

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}

	/* query := "select nr_protocolo, nr_seq_protocolo from protocolo_convenio where nr_seq_protocolo = 210"
	var nrProtocolo string
	var nr_seq_protocolo int

	err = db.QueryRow(query).Scan(&nrProtocolo, &nr_seq_protocolo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Protocolo: ", nrProtocolo, "Sequência: ", nr_seq_protocolo) */

	fmt.Println("Connection to Oracle database successful!")

	//fmt.Println("nrProtocolo:", nrProtocolo)

	db.Close()

	return db, nil
}
