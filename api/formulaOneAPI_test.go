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
	tests := []struct {
		name     string
		endpoint string
	}{
		{"Test One", "f1/current/next.json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawJson, err := fetchFormulaOneApi(tt.endpoint)
			if err != nil {
				t.Errorf(err.Error())
			}
			log.Println(string(rawJson))
		})
	}
}
