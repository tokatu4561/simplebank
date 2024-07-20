package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tokatu4561/simplebank/api"
	db "github.com/tokatu4561/simplebank/db/sqlc"
	"github.com/tokatu4561/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannnot load config", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannnot connect to db", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewSever(config, store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannnot start server", err)
	}
}