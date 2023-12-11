package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	c "github.com/brianqian/go-http-server/cmd/webapp-api/controllers"
	"github.com/brianqian/go-http-server/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	envport := os.Getenv("PORT")
	port, err := strconv.Atoi(envport)
	if err != nil {
		log.Fatal("Error converting port", err)
	}
	if port == 0 {
		port = 8080
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	svr := server.New(
		server.WithHost(host),
		server.WithPort(port),
		server.WithTimeout(time.Minute),
	)

	r := svr.GetServer()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", c.GetUsers)
	})

	err = http.ListenAndServe(":"+strconv.Itoa(port), r)

	if err != nil {
		log.Fatal(err)
	}
}
