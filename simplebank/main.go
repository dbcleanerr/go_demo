package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go_demo/simplebank/api"
	db "go_demo/simplebank/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:welcome@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:3309"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	server := api.NewServer(db.NewStore(conn))
	server.Start(serverAddress)
}
