package ent

import (
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDatabase() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbUser, dbPass, dbHost, dbPort, dbName)

	driver, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	client := NewClient(Driver(driver))

	// 自動マイグレーションツールを実行する
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

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
