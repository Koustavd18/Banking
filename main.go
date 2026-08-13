package main

import (
	"database/sql"
	"log"

	"github.com/Koustavd18/Banking/api"
	db "github.com/Koustavd18/Banking/db/sqlc"
	"github.com/Koustavd18/Banking/utils"
	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.Addr)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
