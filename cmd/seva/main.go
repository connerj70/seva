package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/connerj70/seva/cmd/seva/internal/handlers"
	"github.com/connerj70/seva/internal/mongo"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log := log.New(os.Stdout, "SEVA : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	// Setup config
	var cfg struct {
		API struct {
			HostPort     string
			JWTSecretKey string
		}
		DB struct {
			ConnectionString string
			ENV              string
		}
	}
	cfg.API.HostPort = os.Getenv("HOST_PORT")
	cfg.API.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	cfg.DB.ENV = os.Getenv("DB_ENV")
	cfg.DB.ConnectionString = os.Getenv("MONGO_CONNECTION_STRING")
	log.Printf("Config: %v", cfg)

	// Open connection to database
	client, err := mongo.Open(cfg.DB.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}
	// Set the database environment
	db := client.Database(cfg.DB.ENV)

	// Router
	router := httprouter.New()
	handlers.Register(router, db, cfg.API.JWTSecretKey)

	server := http.Server{
		Addr:    cfg.API.HostPort,
		Handler: router,
	}

	errors := make(chan error, 1)
	// Listen
	go func() {
		log.Printf("main : API listening on %s", server.Addr)
		errors <- server.ListenAndServe()
	}()

	select {
	case err := <-errors:
		return fmt.Errorf("server error %w", err)
	}
}
