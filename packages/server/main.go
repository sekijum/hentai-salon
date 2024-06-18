package main

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "entgo.io/ent/dialect/sql"
    "server/infrastructure/ent"
    "server/presentation/router"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
        log.Fatal("Database environment variables not set properly")
    }

    dataSourceName := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=True"
    driver, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("failed to open sql connection: %v", err)
    }

    client := ent.NewClient(ent.Driver(driver))

    // Run the auto migration tool
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    defer client.Close()

    log.Println("Database connection and migration completed successfully!")

    // Create a new Gin router
    r := gin.Default()

    // Set up the routes
    router.SetupAdminRoutes(r, client)
    router.SetupClientRoutes(r, client)

    // Start the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("failed to run server: %v", err)
    }
}
