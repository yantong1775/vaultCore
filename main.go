package main

import (
	"database/sql"
	"log"

	"github.com/yantong1775/vaultCore/api"
	db "github.com/yantong1775/vaultCore/db/sqlc"
	"github.com/yantong1775/vaultCore/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(testDB)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
