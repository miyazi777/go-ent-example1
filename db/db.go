package db

import (
	"fmt"
	"test/ent"
)

func NewDBClient() *ent.Client {
	user := "user"
	password := "pass"
	port := "5432"
	host := "localhost"
	dbName := "testdb"
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	client, err := ent.Open("postgres", url)
	if err != nil {
		fmt.Printf("failed connecting to postgres: %v", err)
	}
	return client
}

func CloseDB(client *ent.Client) {
	err := client.Close()
	if err != nil {
		fmt.Printf("failed close to db: %v", err)
	}
}
