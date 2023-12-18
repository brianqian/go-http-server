package main

import (
	s "base/cmd/api"
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

	var dbInstance *db.Database

	switch args := os.Args[1]; args {
	case "seed":
		dbInstance = db.New(db.DbConfig{MinConnections: "3", MaxConnections: "8"})
		data.SeedImportedFens(dbInstance)
	default:
		dbInstance = db.New(db.DbConfig{MinConnections: "3", MaxConnections: "4"})
		s.Main(dbInstance)
	}
}
