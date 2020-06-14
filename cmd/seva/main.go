package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/ardanlabs/conf"
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
			HostPort     string `conf:"default:0.0.0.0:8080"`
			JWTSecretKey string `conf:"default:aaa,noprint"`
		}
		DB struct {
			ConnectionString string `conf:"default:mongodb+srv://prod-machine:JxgnpPEQuNFTJxel@cluster0-bwnad.mongodb.net/dev?retryWrites=true&w=majority"`
			ENV              string `conf:"default:dev"`
		}
	}

	if err := conf.Parse(os.Args[1:], "SEVA", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("SEVA", &cfg)
			if err != nil {
				return fmt.Errorf("generating config usage %w", err)
			}
			fmt.Println(usage)
			return nil
		}
		return fmt.Errorf("parsing config %w", err)
	}

	// Open connection to database
	client, err := mongo.Open(cfg.DB.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}
	// Set the database environment
	db := client.Database(cfg.DB.ENV)

	// Router
	router := httprouter.New()
	handlers.Register(router, db, cfg.API.JWTSecretKey, log)

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
