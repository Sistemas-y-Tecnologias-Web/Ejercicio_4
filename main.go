package main

import (
	"videogames-api/config"
	"videogames-api/server"
)

func main() {
	db := config.NewDB()
	defer db.Close()

	server.Start(db)
}
