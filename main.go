package main

import (
	"database/sql"
	"log"

	"github.com/yantong1775/vaultCore/api"
	db "github.com/yantong1775/vaultCore/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/vaultCore?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(testDB)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
