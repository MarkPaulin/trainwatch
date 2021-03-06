package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPath      = flag.String("db", "trains-db.sqlite", "Database path")
	watchPeriod = flag.Int("watch", 1, "Watch period (days)")
)

const query string = "REPLACE INTO trains (station_code, timestamp, headcode, " +
	"uid, arrive_status, arrive_timestamp, depart_timestamp, date, platform, " +
	"status, expected_arrive_time, expected_depart_time, delay, origin, " +
	"destination) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

func scrapeTrains() error {
	ts := time.Now().Format("20060102 150405")

	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	stations, err := getStationList()
	if err != nil {
		return err
	}

	for k := 0; k < len(stations.Stations); k++ {
		board, err := getStationBoard(stations.Stations[k].Code)
		if err != nil {
			return err
		}

		if len(board.Service) != 0 {
			for s := 0; s < len(board.Service); s++ {
				var date string

				if len(board.Service[s].ArriveTime.Timestamp) > 0 {
					date = board.Service[s].ArriveTime.Timestamp[:8]
				} else {
					date = board.Service[s].DepartTime.Timestamp[:8]
				}

				_, err := statement.Exec(
					stations.Stations[k].Code,
					ts,
					board.Service[s].Headcode,
					board.Service[s].UID,
					board.Service[s].ArriveTime.Arrived,
					board.Service[s].ArriveTime.Time,
					board.Service[s].DepartTime.Time,
					date,
					board.Service[s].Platform.Number,
					board.Service[s].ServiceStatus.Status,
					board.Service[s].ExpectedArriveTime.Time,
					board.Service[s].ExpectedDepartTime.Time,
					board.Service[s].Delay.Minutes,
					board.Service[s].Origin1.Name,
					board.Service[s].Destination1.Name,
				)

				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func main() {
	flag.Parse()

	err := scrapeTrains()
	if err != nil {
		fmt.Println(err)
	}

	ticker := time.NewTicker(2 * time.Minute)

	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				scrapeTrains()
				if err != nil {
					fmt.Println(err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(time.Duration(*watchPeriod*24) * time.Hour)
	ticker.Stop()
	quit <- true
	fmt.Println("Scraping complete")
}
