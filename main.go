package main

import (
	"database/sql"
	"log"

	"github.com/tokatu4561/simple-bank/api"
	db "github.com/tokatu4561/simple-bank/db/sqlc"
	"github.com/tokatu4561/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannnot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannnot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewSever(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannnot start server", err)
	}
}