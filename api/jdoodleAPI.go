package api

import (
	"bytes"
	"encoding/json"
	"github.com/Jacobbrewer1/botter/config"
	"io/ioutil"
	"net/http"
)

func RunCode(exec ExecuteInput) (ExecuteOutput, error) {
	rawJson, err := requestCompile(exec)
	if err != nil {
		return ExecuteOutput{}, err
	}
	return decodeCompile(rawJson)
}

func decodeCompile(rawJson json.RawMessage) (ExecuteOutput, error) {
	var c ExecuteOutput
	err := json.Unmarshal(rawJson, &c)
	return c, err
}

func requestCompile(exec ExecuteInput) (json.RawMessage, error) {
	body, err := json.Marshal(exec)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, *config.JsonConfig.Endpoints.JdoodleApiEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
