package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/mehmetcc/todo-backend/internal/config"
	"github.com/mehmetcc/todo-backend/internal/controllers"
	"github.com/mehmetcc/todo-backend/internal/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	err := godotenv.Load(filepath.Join("..", "..", "deployments", ".env"))
	if err != nil {
		log.Println("No .env file found")
	}

	// Run migrations
	dbSQL := config.GetDBConnection()
	config.RunMigrations(dbSQL)
	dbSQL.Close()

	// GORM setup
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize controllers with DB
	controllers.SetDB(db)

	// Setup router
	router := routers.SetupRouter()

	// Run application
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	router.Run(":" + appPort)
}
