package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type stationList struct {
	Stations []struct {
		Code string
		Name string
	}
}

func getStationList() stationList {
	resp, err := http.Get("https://apis.opendatani.gov.uk/translink/")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	blob, _ := ioutil.ReadAll(resp.Body)

	var stations stationList
	json.Unmarshal(blob, &stations)

	return stations
}