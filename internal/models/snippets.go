package models

import (
    "database/sql" // Required to hold the database pool pointer
    "time"         // Represents SQL DATETIME / TIMESTAMP fields
)

// Snippet matches the columns of your 'snippets' table in MySQL.
type Snippet struct {
    ID      int
    Title   string
    Content string
    Created time.Time
    Expires time.Time
}

// SnippetModel wraps our database connection pool.
// Any method attached to SnippetModel can run SQL queries via m.DB.
type SnippetModel struct {
    DB *sql.DB // Shared MySQL connection pool pointer created in main.go
}

// Insert will execute an INSERT SQL query and return the new snippet ID.
// (Currently a stub returning 0, nil as a placeholder for upcoming chapters).
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
    return 0, nil
}

// Get will execute a SELECT query for a single snippet by its ID.
func (m *SnippetModel) Get(id int) (Snippet, error) {
    return Snippet{}, nil
}

// Latest will execute a SELECT query returning the 10 most recent snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
    return nil, nil
}