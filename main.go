package main

import (
	"database/sql"
	"learn/back-end/server"
	"learn/back-end/store/store"
	"learn/back-end/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	server, err := server.NewServer(conf, store.NewStore(conn))
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	log.Fatal("start server: ", server.Start(":3000"))
}
