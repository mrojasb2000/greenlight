package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	// We default to using the port number 4000 and then environment "development" if no corresponding flags are provided
	flag.IntVar(&cfg.port, "port", 4000, "API Server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger which writes messages to the standard out stream
	// prefixed with the current date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of the application struct, containing the config struct and the logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a new servemux and add a /v1/healthcheck toute which dispatches requests
	// to the healthcheckHandler method (which we will create in a moment
	/*
		mux := http.NewServeMux()
		mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	*/

	// Declare a HTTP server with some sensible timeout settings, which lstens on the
	// port provided in the config struct and uses the servemux we created above as the handler
	/*
		srv := &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.port),
			Handler:      mux,
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		}*/

	// Use the httprouter instance returned by app.routes() as the server handler
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}