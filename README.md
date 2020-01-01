# Trainwatch

Tiny program for pulling data from Translink NI Railways Open Data API and popping it into a database.

Build `main.exe` with `go build main.go stationBoard.go stationList.go`. It takes two command line flags:
* `db` name of the SQLite database to write to (default is `trains-db.sqlite`)
* `watch` number of days (default is a day)

There's also a SQL file for setting up tables.
