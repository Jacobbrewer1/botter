package api

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDecodeTrendingResponse(t *testing.T) {
	tests := []struct {
		name         string
		jsonData     json.RawMessage
		expectedJson Trending
	}{
		{"Trending 1", testerSetupJsonRaw(testTrendingResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testTrendingResp))},
		{"Trending 2", testerSetupJsonRaw(testTrendingResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testTrendingResp))},
		{"Trending 3", testerSetupJsonRaw(testTrendingResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testTrendingResp))},
		{"Trending 4", testerSetupJsonRaw(testTrendingResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testTrendingResp))},
		{"Trending 5", testerSetupJsonRaw(testTrendingResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testTrendingResp))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotJson, err := decodeTrendingResponse(tt.jsonData)
			if err != nil {
				log.Println(err)
			}
			for pos, elm := range gotJson.Data {
				if elm != tt.expectedJson.Data[pos] {
					t.Errorf("decodeTrendingResponse(tt.jsonData) = %v, expected %v", gotJson, tt.expectedJson)
				}
			}
		})
	}
}

func TestDecodeSearchResponse(t *testing.T) {
	tests := []struct {
		name         string
		jsonData     json.RawMessage
		expectedJson Trending
	}{
		{"Search 1", testerSetupJsonRaw(testSearchResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testSearchResp))},
		{"Search 2", testerSetupJsonRaw(testSearchResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testSearchResp))},
		{"Search 3", testerSetupJsonRaw(testSearchResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testSearchResp))},
		{"Search 4", testerSetupJsonRaw(testSearchResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testSearchResp))},
		{"Search 5", testerSetupJsonRaw(testSearchResp), testerSetupJsonFormattedTrending(testerSetupJsonRaw(testSearchResp))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotJson, err := decodeSearchResponse(tt.jsonData)
			if err != nil {
				log.Println(err)
			}
			for pos, elm := range gotJson.Data {
				if elm != tt.expectedJson.Data[pos] {
					t.Errorf("decodeSearchResponse(tt.jsonData) = %v, expected %v", gotJson, tt.expectedJson)
				}
			}
		})
	}
}

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{"Build Query 1", "this is the first test", "this/is/the/first/test"},
		{"Build Query 2", "This is test number two", "this/is/test/number/two"},
		{"Build Query 3", "This test contains,    spaces and, punctuation!", "this/test/contains/spaces/and/punctuation"},
		{"Build Query 4", "This contains  a lot    of double spaces", "this/contains/a/lot/of/double/spaces"},
		{"Build Query 5", "This test contains, a lot of punctuation!", "this/test/contains/a/lot/of/punctuation"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse := buildSearchQuery(tt.query)
			if gotResponse != tt.expected {
				t.Errorf("buildSearchQuery(tt.query) = %v, expected %v", gotResponse, tt.expected)
			}
		})
	}
}

func TestDecodeResponse(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected response
	}{
		{"Response code 200", 200, Response200},
		{"Response code 400", 400, Response400},
		{"Response code 403", 403, Response403},
		{"Response code 404", 404, Response404},
		{"Response code 414", 414, Response414},
		{"Response code 429", 429, Response429},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse := DecodeResponse(tt.code)
			if gotResponse != tt.expected {
				t.Errorf("DecodeResponse(tt.code) = %v, expected %v", gotResponse, tt.expected.Description)
			}
		})
	}
}
