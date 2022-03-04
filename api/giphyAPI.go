package api

import (
	"encoding/json"
	"github.com/Jacobbrewer1/botter/config"
	"github.com/Jacobbrewer1/botter/helper"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	GifText     = "gifs"
	StickerText = "stickers"
)

// Gif Endpoints
// https://api.giphy.com/v1/gifs/search?api_key=WmpMJmw9T5XxeCM2C2HQSllEEAi6wlA7&q=test
// https://api.giphy.com/v1/gifs/trending?api_key=WmpMJmw9T5XxeCM2C2HQSllEEAi6wlA7

// Sticker Endpoints
// https://api.giphy.com/v1/stickers/search?api_key=WmpMJmw9T5XxeCM2C2HQSllEEAi6wlA7&q=test
// https://api.giphy.com/v1/stickers/trending?api_key=WmpMJmw9T5XxeCM2C2HQSllEEAi6wlA7

// https://giphy.com/embed/dyp02B1LtyGWF6rgsv - Powered by giphy gif

func TrendingSearch(endpoint string) (Trending, error) {
	url := "https://api.giphy.com/v1/" + endpoint + "/trending?api_key=" + *config.ApiSecrets.GiphyApiToken
	rawJson, err := getResponse(url)
	if err != nil {
		log.Println(err)
	}
	return decodeTrendingResponse(rawJson)
}

func SearchQuery(endpoint, tempSearch string) (Search, error) {
	var search string
	if strings.Contains(tempSearch, " ") {
		search = buildSearchQuery(tempSearch)
	} else {
		search = tempSearch
	}
	log.Printf("Search query: %v", search)
	url := "https://api.giphy.com/v1/" + endpoint + "/search?api_key=" + *config.ApiSecrets.GiphyApiToken + "&q=" + search
	log.Printf("API url: %v\n", url)
	rawJson, err := getResponse(url)
	if err != nil {
		log.Println(err)
	}
	return decodeSearchResponse(rawJson)
}

func decodeTrendingResponse(rawJson json.RawMessage) (Trending, error) {
	var d Trending
	err := json.Unmarshal(rawJson, &d)
	if err != nil {
		return d, err
	}
	log.Printf("Length of data response: %v", len(d.Data))
	return d, nil
}

func buildSearchQuery(search string) string {
	return strings.ToLower(strings.Join(strings.Fields(helper.RemoveNonAlphaChars(search)), "/"))
}

func decodeSearchResponse(rawJson json.RawMessage) (Search, error) {
	var d Search
	err := json.Unmarshal(rawJson, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func getResponse(url string) (json.RawMessage, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	log.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
