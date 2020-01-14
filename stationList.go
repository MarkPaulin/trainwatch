package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type stationList struct {
	Stations []struct {
		Code string
		Name string
	}
}

func getStationList() (stationList, error) {
	resp, err := http.Get("https://apis.opendatani.gov.uk/translink/")
	var stations stationList

	if err != nil {
		return stations, err
	}

	defer resp.Body.Close()

	blob, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(blob, &stations)

	return stations, nil
}
