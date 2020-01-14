package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type stationBoard struct {
	XMLName      xml.Name `xml:"StationBoard"`
	Text         string   `xml:",chardata"`
	Name         string   `xml:"name,attr"`
	Tiploc       string   `xml:"tiploc,attr"`
	Crs          string   `xml:"crs,attr"`
	PlatformData string   `xml:"PlatformData,attr"`
	Timestamp    string   `xml:"Timestamp,attr"`
	TridentID    string   `xml:"TridentId"`
	Service      []struct {
		Text        string `xml:",chardata"`
		Headcode    string `xml:"Headcode,attr"`
		UID         string `xml:"Uid,attr"`
		RetailID    string `xml:"RetailID,attr"`
		TigerID     string `xml:"TigerID,attr"`
		ServiceType struct {
			Text string `xml:",chardata"`
			Type string `xml:"Type,attr"`
		} `xml:"ServiceType"`
		ArriveTime struct {
			Text      string `xml:",chardata"`
			Time      string `xml:"time,attr"`
			Arrived   string `xml:"Arrived,attr"`
			Timestamp string `xml:"timestamp,attr"`
		} `xml:"ArriveTime"`
		DepartTime struct {
			Text          string `xml:",chardata"`
			Time          string `xml:"time,attr"`
			Timestamp     string `xml:"timestamp,attr"`
			Sorttimestamp string `xml:"sorttimestamp,attr"`
		} `xml:"DepartTime"`
		Platform struct {
			Text    string `xml:",chardata"`
			Number  string `xml:"Number,attr"`
			Changed string `xml:"Changed,attr"`
			Parent  string `xml:"Parent,attr"`
		} `xml:"Platform"`
		SecondaryServiceStatus string `xml:"SecondaryServiceStatus"`
		ServiceStatus          struct {
			Text   string `xml:",chardata"`
			Status string `xml:"Status,attr"`
		} `xml:"ServiceStatus"`
		ExpectedDepartTime struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
		} `xml:"ExpectedDepartTime"`
		ExpectedArriveTime struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
		} `xml:"ExpectedArriveTime"`
		Delay struct {
			Text    string `xml:",chardata"`
			Minutes string `xml:"Minutes,attr"`
		} `xml:"Delay"`
		ExpectedDepartStatus struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
		} `xml:"ExpectedDepartStatus"`
		ExpectedArriveStatus struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
		} `xml:"ExpectedArriveStatus"`
		DelayCause string `xml:"DelayCause"`
		LastReport struct {
			Text     string `xml:",chardata"`
			Tiploc   string `xml:"tiploc,attr"`
			Time     string `xml:"time,attr"`
			Type     string `xml:"type,attr"`
			Station1 string `xml:"station1,attr"`
			Station2 string `xml:"station2,attr"`
		} `xml:"LastReport"`
		CommentLine           string `xml:"CommentLine"`
		CommentLine2          string `xml:"CommentLine2"`
		ArrivalComment1       string `xml:"ArrivalComment1"`
		ArrivalComment2       string `xml:"ArrivalComment2"`
		PlatformComment1      string `xml:"PlatformComment1"`
		PlatformComment2      string `xml:"PlatformComment2"`
		DepartureComment1     string `xml:"DepartureComment1"`
		DepartureComment2     string `xml:"DepartureComment2"`
		AssociatedPageNotices string `xml:"AssociatedPageNotices"`
		ChangeAt              string `xml:"ChangeAt"`
		Operator              struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Code  string `xml:"code,attr"`
			Brand string `xml:"brand,attr"`
		} `xml:"Operator"`
		Origin1 struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name,attr"`
			Tiploc string `xml:"tiploc,attr"`
			Crs    string `xml:"crs,attr"`
		} `xml:"Origin1"`
		Destination1 struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name,attr"`
			Tiploc string `xml:"tiploc,attr"`
			Crs    string `xml:"crs,attr"`
			Ttarr  string `xml:"ttarr,attr"`
			Etarr  string `xml:"etarr,attr"`
		} `xml:"Destination1"`
		Via                string `xml:"Via"`
		Coaches1           string `xml:"Coaches1"`
		Incident           string `xml:"Incident"`
		Dest1CallingPoints struct {
			Text             string `xml:",chardata"`
			NumCallingPoints string `xml:"NumCallingPoints,attr"`
			CallingPoint     []struct {
				Text   string `xml:",chardata"`
				Name   string `xml:"Name,attr"`
				Tiploc string `xml:"tiploc,attr"`
				Crs    string `xml:"crs,attr"`
				Ttarr  string `xml:"ttarr,attr"`
				Ttdep  string `xml:"ttdep,attr"`
				Etarr  string `xml:"etarr,attr"`
				Etdep  string `xml:"etdep,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"CallingPoint"`
		} `xml:"Dest1CallingPoints"`
	} `xml:"Service"`
	Incident struct {
		Text    string `xml:",chardata"`
		Summary string `xml:"Summary,attr"`
	} `xml:"Incident"`
}

func getStationBoard(stationCode string) (stationBoard, error) {
	url := "https://apis.opendatani.gov.uk/translink/" + stationCode + ".xml"

	var board stationBoard

	resp, err := http.Get(url)

	if err != nil {
		return board, err
	}

	defer resp.Body.Close()

	blob, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(blob, &board)

	return board, nil
}
