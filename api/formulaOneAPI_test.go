package api

import (
	"encoding/json"
	"github.com/Jacobbrewer1/botter/config"
	"io/ioutil"
	"log"
	"testing"
)

func TestGetNextRace(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test One"},
	}

	file, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(file, &config.JsonConfig)
	if err != nil {
		log.Println(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawJson, err := GetNextRace()
			if err != nil {
				t.Errorf(err.Error())
			}
			log.Println(rawJson)
		})
	}
}

func TestFetchFormulaOneApi(t *testing.T) {
	file, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(file, &config.JsonConfig)
	if err != nil {
		log.Println(err)
		return
	}

	tests := []struct {
		name     string
		endpoint string
	}{
		{"next race", nextF1RaceEndpoint},
		{"driver standings", currentF1DriverStandingsEndpoint},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fetchFormulaOneApi(tt.endpoint)
			if err != nil {
				t.Errorf("fetchFormulaOneApi(tt.endpoint) error = %v", err)
			}
		})
	}
}

func TestGetDriverStandings(t *testing.T) {
	file, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(file, &config.JsonConfig)
	if err != nil {
		log.Println(err)
		return
	}

	tests := []struct {
		name     string
	}{
		{"test one"},
		{"test two"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, err := GetDriverStandings()
			if err != nil {
				t.Errorf("GetDriverStandings() error = %v", err)
			}
			log.Println(x)
		})
	}
}
