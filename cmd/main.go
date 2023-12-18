package main

import (
	data "base/data/seeds"
	"base/internal/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env vars")
	}

	var (
		username      = os.Getenv("DB_USERNAME")
		password      = os.Getenv("DB_PASSWORD")
		host          = os.Getenv("DB_HOST")
		port          = os.Getenv("DB_PORT")
		database_name = os.Getenv("DB_NAME")
	)

	dbInstance := db.New(db.DbConfig{MinConnections: "1", MaxConnections: "4", Username: username, Password: password, Host: host, Port: port, Database_name: database_name})

	// s.Main()
	data.SeedImportedFens(dbInstance)
}
