package api

import (
	"encoding/json"
	"github.com/Jacobbrewer1/botter/config"
	"io/ioutil"
	"net/http"
)

const (
	nextF1RaceEndpoint               = "f1/current/next.json"
	currentF1DriverStandingsEndpoint = "f1/current/driverStandings.json"
)

func GetNextRace() (Race, error) {
	rawJson, err := fetchFormulaOneApi(nextF1RaceEndpoint)
	if err != nil {
		return Race{}, err
	}
	gp, err := decodeGrandPix(rawJson)
	if err != nil {
		return Race{}, err
	}
	return *gp.RaceTable.Races[0], nil
}

func decodeGrandPix(rawJson json.RawMessage) (GrandPixMRDataStruct, error) {
	var g GrandPixMRDataStruct
	err := json.Unmarshal(rawJson, &g)
	return g, err
}

func GetDriverStandings() (DriverStandingsStruct, error) {
	rawJson, err := fetchFormulaOneApi(currentF1DriverStandingsEndpoint)
	if err != nil {
		return DriverStandingsStruct{}, err
	}
	return decodeDriverStandings(rawJson)
}

func decodeDriverStandings(rawJson json.RawMessage) (DriverStandingsStruct, error) {
	var d DriverStandingsMRDataStruct
	err := json.Unmarshal(rawJson, &d)
	return *d.DriverStandingsStruct, err
}

func fetchFormulaOneApi(endpoint string) (json.RawMessage, error) {
	req, err := http.NewRequest(http.MethodGet, *config.JsonConfig.Endpoints.FormulaOneApiEndpoint+endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
