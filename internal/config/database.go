package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetDBConnection() *sql.DB {
	err := godotenv.Load(filepath.Join("..", "..", "deployments", ".env"))
	if err != nil {
		log.Println("No .env file found")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Println(err)
	}

	db.Close()

	connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return db
}

func RunMigrations(db *sql.DB) {
	migrationPath := filepath.Join("..", "..", "internal", "migrations", "*.sql")
	files, err := filepath.Glob(migrationPath)
	if err != nil {
		log.Fatalf("Failed to read migrations: %v", err)
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", file, err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatalf("Failed to execute migration %s: %v", file, err)
		}

		log.Printf("Applied migration: %s", filepath.Base(file))
	}
}
