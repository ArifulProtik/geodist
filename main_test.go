package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

type DataTable struct {
	id       int
	lat      float64
	lon      float64
	rlat     float64
	rlon     float64
	distance float64
}

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

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestGetSortedLocation(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	newData := DataTable{
		id:       5,
		lat:      23.622641,
		lon:      90.499794,
		rlat:     0.412292863466606,
		rlon:     1.57951937767661,
		distance: 19.99125734738771,
	}

	query := `SELECT 
	id,lat,lon, 
	(
  6371  *
   acos(cos(radians(23.777176)) * 
   cos(rlat) * 
   cos(rlon - 
   radians(90.399452)) + 
   sin(radians(23.777176)) * 
   sin(rlat))
	) AS distance 
	FROM location2
	GROUP BY id HAVING distance < 20 
	ORDER BY distance LIMIT 1`

	rows := sqlmock.NewRows([]string{"id", "lat", "lon", "rlat", "rlon", "distance"}).AddRow(newData.id, newData.lat, newData.lon, newData.rlat, newData.rlon, newData.distance)
	mock.ExpectQuery(query).WillReturnRows(rows)
	// add sqlite db connection
	sqlDB, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	Locations, err := GetSortedLocation(sqlDB, InputData{lat: 23.777176, lon: 90.399452, radius: 20})
	assert.NotNil(t, Locations)
	assert.Equal(t, newData.id, Locations[0].id)
	assert.NoError(t, err)
}
