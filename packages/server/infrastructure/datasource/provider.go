package datasource

import (
	"context"
	"fmt"
	"log"
	"os"
	"server/domain/lib/util"
	"server/infrastructure/ent"
	"server/infrastructure/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDatabase() (*ent.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbUser, dbPass, dbHost, dbPort, dbName)

	driver, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	client := ent.NewClient(ent.Driver(driver))

	// err = CreateAdminUser(client)
	// if err != nil {
	// 	log.Fatalf("failed to create or find user: %v", err)
	// }

	return client, nil
}

func ProvideClient() (*ent.Client, func(), error) {
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

func CreateAdminUser(client *ent.Client) error {
	ctx := context.Background()

	email := "jumpei20010910@gmail.com"
	password := "Ninja1891"

	existingUser, err := client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("failed to query user: %w", err)
	}

	if existingUser != nil {
		log.Printf("User already exists: %s", email)
		return nil
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	_, err = client.User.Create().
		SetName("関口純平").
		SetEmail(email).
		SetPassword(hashedPassword).
		SetRole(1).
		SetStatus(0).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	log.Printf("New user created: %s", email)
	return nil
}
