package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nanmenkaimak/chat-app/internal/dbs/postgres"
)

func main() {
	_, err := postgres.New()
	fmt.Println(err)
}
