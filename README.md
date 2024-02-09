# geodist

This is a simple Golang Project to Calculate And Sort the distance Between Several Geo Locations. 
## Problem Statement 
Given a list of geolocation coordinates and a current location, determine the coordinate closest to the
current location from the list. The search radius can be adjusted as needed. 
### Motivation 
A lot of current apps now provide the option to show nearby data, which could include user profiles, physical locations, or any information associated with a specific place. This feature is undoubtedly beneficial, and there's a significant possibility that you might find yourself needing to incorporate it, be it for a client or your personal project.

## Prerequisites 
- Golang 1.18 or later
- Sqlite3.31+ 

## Usage 
### To Build
```bash 
$ go mod tidy
$ go build -o geodist
``` 
### For Help

```bash
./geodist -h # for help
```
### For Running

```bash
./geofinder -lt=<latitude> -ln=<longitude> -r=<radius> 
```

### Adding More Data 
- The data is stored in a sqlite database. You can add more data by adding more rows to the database.
- The database is located at `./test.db` 
- The table name is `location2`
- The columns are `id`, `lat`, `lon` `rlat` and `rlon` (rlat and rlon are the radians of lat and lon Generative Columns)
For example: [test.db]  


```sql
INSERT INTO location2(id,lat,lon) VALUES(2,23.643999,88.855637);

```
### Recreate The Database Table 
```sql
CREATE TABLE location2 (
    id INTEGER PRIMARY KEY,
    lat REAL, -- latitude and longtitude as FLOAT( 10, 8 ) NB: see reference for more info
    lon REAL,
    rlat REAL GENERATED ALWAYS AS (RADIANS(lat)) STORED,
    rlon REAL GENERATED ALWAYS AS (RADIANS(lon)) STORED
);
```


## References 
- [Haversine Formula](https://en.wikipedia.org/wiki/Haversine_formula)

