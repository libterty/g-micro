package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/libterty/g-micro/api"
	db "github.com/libterty/g-micro/db/sqlc"
	"github.com/libterty/g-micro/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannnot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Test connection: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.HTTPServerAddress); err != nil {
		log.Fatal("Server Failer: ", err)
	}

}
