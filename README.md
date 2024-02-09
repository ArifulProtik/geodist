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
go build -o geodist
``` 
### For Help

```bash
./geodist -h # for help
```
### For Running

```bash
./geofinder -lt=<latitude> -ln=<longitude> -r=<radius> 
```
## References 
- [Haversine Formula](https://en.wikipedia.org/wiki/Haversine_formula)

