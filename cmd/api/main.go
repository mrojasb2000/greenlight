package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/mrojasb2000/greenlight/internal/data"
)

const version = "1.0.0"

// Add a db struct field to hold the configuration settigs for our database connection
// pool. For now the only holds the DSN, which we will read in from a command-line flag.

// Add maxOpenConns, maxIdleConns and maxIdleTime fields to hold the connection
// settings for the connection pool.

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

// Add a models field to hold our new Models struct.
type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	// We default to using the port number 4000 and then environment "development" if no corresponding flags are provided
	flag.IntVar(&cfg.port, "port", 4000, "API Server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Read the DSN value from the db-dsn command-line flag into the config struct. We
	// default to using our development DSN if no flag is provided.
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

	// Read the connection pool settings from command-line flags into the config struct.
	// Notice the default values that we're using?
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "postgreSQL max open connections")

	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")

	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	// Initialize a new logger which writes messages to the standard out stream
	// prefixed with the current date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Call the openDB() helper function (see below) to create the connection pool,
	// passing in the config struct. If this returns an error, we log it and exit the
	// application immediately.
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// Defer a call to db.Close() so that the connection pool is closed before the
	// main() function exit.
	defer db.Close()

	// Also log a message to say the connection pool has been successfully
	// established.
	logger.Printf("database connection pool established")

	// Declare an instance of the application struct, containing the config struct and the logger
	// Use the data.NewModels() function to initialize a Models struct, passing in the
	// connection pool as a parameter.
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

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
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

// The openDB function returns a sql.DB connection pool.
func openDB(cfg config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	// struct.
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Use the time.ParseDuration() function to convert the idle timeout durantion type
	// to a time.Duration type.
	durantion, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	// Set the maximum idle timeout
	db.SetConnMaxIdleTime(durantion)

	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5 second deadline, then this will return an
	// err
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Return the sql.Db connection pool
	return db, nil
}
