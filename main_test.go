package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

// TestMain is a special function that runs before any tests are run in the package
func TestMain(m *testing.M) {
	// os.Exit skips defer calls
	// so we need to call another function
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

// run is a helper function that sets up the database and runs the tests for Database Integration
func run(m *testing.M) (code int, err error) {
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		return -1, fmt.Errorf("could not connect to database: %w", err)
	}
	defer db.Close()

	return m.Run(), nil
}
