package main

import (
	"context"
	"fmt"
	"test/db"
	"test/ent/migrate"

	_ "github.com/lib/pq"
)

func main() {
	client := db.NewDBClient()
	ctx := context.Background()

	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		fmt.Printf("failed creating schema resources: %v", err)
	}

	db.CloseDB(client)
}
