//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"server/infrastructure/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()
	dir, err := atlas.NewLocalDir("/app/infrastructure/ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeInspect),
		schema.WithDialect(dialect.MySQL),
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("mysql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	err = migrate.NamedDiff(ctx, dbURL, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}

	log.Println("Migration file generated successfully")
}
