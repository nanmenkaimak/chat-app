package main

import (
	_ "github.com/lib/pq"
	"github.com/nanmenkaimak/chat-app/internal/dbs/postgres"
	"github.com/nanmenkaimak/chat-app/internal/handlers"
)

const portNumber = ":8080"

func main() {
	db, err := postgres.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	routes()
}
