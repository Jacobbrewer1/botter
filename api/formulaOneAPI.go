package api

import (
	"encoding/json"
	"github.com/Jacobbrewer1/botter/config"
	"io/ioutil"
	"net/http"
)

const nextF1Endpoint = "f1/current/next.json"

func GetNextRace() (Race, error) {
	rawJson, err := fetchFormulaOneApi(nextF1Endpoint)
	if err != nil {
		return Race{}, err
	}
	gp, err := decodeGrandPix(rawJson)
	if err != nil {
		return Race{}, err
	}
	return *gp.GrandPix.RaceTable.Races[0], nil
}

func decodeGrandPix(rawJson json.RawMessage) (MRDataStruct, error) {
	var g MRDataStruct
	err := json.Unmarshal(rawJson, &g)
	return g, err
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
