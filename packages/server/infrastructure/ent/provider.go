package ent

import (
	"context"
	"log"
	cfg "server/shared/config"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDatabase() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dataSourceName := cfg.GetDatabaseDSN()
	driver, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	client := NewClient(Driver(driver))

	// 自動マイグレーションツールを実行する
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Database connection and migration completed successfully!")
	return client, nil
}

func ProvideClient() (*Client, func(), error) {
	client, err := InitDatabase()
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err := client.Close(); err != nil {
			log.Printf("failed closing client: %v", err)
		}
	}
	return client, cleanup, nil
}
