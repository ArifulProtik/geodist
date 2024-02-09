package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

const (
	Unit = 6371 // 6371 for KM and 3959 for Miles
	Limt = 10   // Limit of the number of locations to be returned
)

type (
	InputData struct {
		lat    float64
		lon    float64
		radius int
	}
	OutputData struct {
		id       int
		lat      float64
		lon      float64
		distance float64
	}
)

func main() {
	inputData := ReciveFlags()

	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // Since no Defer Statement is used
	}
	sortedList, err := GetSortedLocation(db, inputData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if len(sortedList) == 0 {
		fmt.Println("No Location Found within the given radius")
		os.Exit(3)
	}
	fmt.Println("The Neerest location is: \t", sortedList[0].lat, sortedList[0].lon, "Distance(KM):", sortedList[0].distance)

	// Table Header
	fmt.Println("The Sorted List(Near to Far):")
	fmt.Println("Latitude\tLongitude\tDistance(KM)")
	// Pretty Print Sorted List of Locations
	for _, loc := range sortedList {
		fmt.Println(loc.lat, "\t", loc.lon, "\t", loc.distance)
	}
}

// GetSortedLocation returns a sorted list(near to far ) of locations based
// on the input Geo Location and Radius from the database
func GetSortedLocation(db *sql.DB, inp InputData) ([]OutputData, error) {
	stmt, err := db.Prepare(`SELECT 
	id,lat,lon, 
	(
   ? *
   acos(cos(radians(?)) * 
   cos(rlat) * 
   cos(rlon - 
   radians(?)) + 
   sin(radians(?)) * 
   sin(rlat))
	) AS distance 
	FROM location2
	GROUP BY id HAVING distance < ? 
	ORDER BY distance LIMIT ? `)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(Unit, inp.lat, inp.lon, inp.lat, inp.radius, Limt)
	if err != nil {
		return nil, err
	}
	var sortedList []OutputData
	for rows.Next() {

		var loc OutputData
		err = rows.Scan(&loc.id, &loc.lat, &loc.lon, &loc.distance)
		if err != nil {
			return nil, err
		}
		sortedList = append(sortedList, loc)
	}

	return sortedList, nil
}

// ReciveFlags receives the latitude, longitude and radius from the command line flags
func ReciveFlags() InputData {
	lt := flag.Float64("lt", 23.777176, "lt command receives latitude of a geolocation")
	ln := flag.Float64("ln", 90.399452, "ln command receives longitude of a geolocation")
	r := flag.Int("r", 100, "r command receives radius of search area")
	flag.Parse()
	return InputData{
		lat:    *lt,
		lon:    *ln,
		radius: *r,
	}
}
