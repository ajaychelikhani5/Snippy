package main

import (
	"database/sql" // Go's standard SQL interface package
	"flag"         // Parses command-line flags (like -addr=:4000)
	"log/slog"     // Structured, fast logger introduced in Go 1.21
	"net/http"     // Standard web server package
	"os"           // Used to exit the program on critical errors (os.Exit)

	// Imports your internal database model package from within this project
	"snippy.aj.net/internal/models"

	// The MySQL driver. The "_" blank identifier means: "Run this package's init()
	// function to register the 'mysql' driver name with database/sql, but don't
	// use any functions from it directly in our code."
	_ "github.com/go-sql-driver/mysql"
)

// application acts as our dependency backpack. Any resource that HTTP handlers
// need (logger, DB models, template cache later in the book) is stored here.
type application struct {
	logger   *slog.Logger         // Pointer to our structured logger
	snippets *models.SnippetModel // Pointer to our MySQL queries model
}

func main() {
	// Define CLI flags with default values. Both return memory addresses (*string).
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@web:pass@tcp(127.0.0.1:3306)/snippetbox?parseTime=true", "MySQL data source name")

	// Reads terminal flags passed by the user (e.g. go run ./cmd/web -addr=:8080)
	// and writes values directly to the memory locations of addr and dsn.
	flag.Parse()

	// Create a structured logger that writes text logs directly to os.Stdout (the terminal).
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Dereference dsn (*dsn) to pass the raw connection string into openDB().
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error()) // Log the connection failure
		os.Exit(1)                // Immediately kill the program with exit code 1
	}

	// defer guarantees db.Close() runs right before main() exits, releasing all pool connections.
	defer db.Close()

	// Instantiate our application backpack.
	// We use & to share this single struct across all route handlers without copying.
	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{
			DB: db,
		}, // Pass our shared db connection into the model
	}

	// Dereference addr (*addr) to read the port text (e.g. ":4000").
	logger.Info("starting server", "addr", *addr)

	// http.ListenAndServe starts the web server on the given port and attaches our router.
	// app.routes() comes from routes.go.
	err = http.ListenAndServe(*addr, app.routes())

	// http.ListenAndServe only returns if an error occurs.
	logger.Error(err.Error())
	os.Exit(1)
}

// openDB wraps sql.Open and tests the connection with a Ping.
// It returns a pointer (*sql.DB) so the caller shares the active connection pool on the heap.
func openDB(dsn string) (*sql.DB, error) {
	// sql.Open initializes the pool but does NOT actually connect to MySQL yet.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// db.Ping() makes the first actual network call to MySQL to confirm credentials work.
	err = db.Ping()
	if err != nil {
		db.Close() // If ping fails, clean up the pool before returning the error
		return nil, err
	}

	return db, nil // Return the active connection pool pointer
}
