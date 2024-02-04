package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	var lat, lon float64
	var radius int
	fmt.Print("Enter latitude, longitude and radius separated by comma (e.g., 23.777176,90.399452,100): ")
	_, err := fmt.Scanf("%f,%f,%d", &lat, &lon, &radius)
	checkErr(err)

	db, err := sql.Open("sqlite", "test.db")
	checkErr(err)
	fmt.Println("ID", "Lattitude", "Longitude", "Distance")
	FindLoc(db, lat, lon, radius)
}

func FindLoc(db *sql.DB, lat float64, lon float64, radius int) {
	stmt, err := db.Prepare(`SELECT 
	id,lat,lon, 
	(
   6371 *
   acos(cos(radians(?)) * 
   cos(radians(lat)) * 
   cos(radians(lon) - 
   radians(?)) + 
   sin(radians(?)) * 
   sin(radians(lat )))
	) AS distance 
	FROM location
	GROUP BY id
	HAVING distance < ? 
	ORDER BY distance`)

	checkErr(err)

	rows, err := stmt.Query(lat, lon, lat, radius)
	checkErr(err)

	for rows.Next() {
		var id int
		var lat float64
		var lon float64
		var distance float64
		err = rows.Scan(&id, &lat, &lon, &distance)
		checkErr(err)
		fmt.Println(id, lat, lon, distance)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
