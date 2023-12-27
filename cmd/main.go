package main

import (
	s "base/cmd/api"
	data "base/data/seeds"
	"base/internal/db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env vars")
	}

	var dbInstance *db.Database

	if len(os.Args) == 1 {
		fmt.Println("Starting server...")
		dbInstance = db.New(db.DbConfig{MinConnections: "3", MaxConnections: "4"})
		s.Main(dbInstance)
	}

	switch args := os.Args[1]; args {
	case "seed":
		dbInstance = db.New(db.DbConfig{MinConnections: "3", MaxConnections: "20"})
		data.SeedImportedFens(dbInstance)
	default:
	}
}
