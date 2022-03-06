package api

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type MRDataStruct struct {
	GrandPix *GrandPixStruct `json:"MRData,omitempty"`
}

type GrandPixStruct struct {
	Series    *string          `json:"series,omitempty"`
	Url       *string          `json:"url,omitempty"`
	Limit     *string          `json:"limit,omitempty"`
	Offset    *string          `json:"offset,omitempty"`
	Total     *string          `json:"total,omitempty"`
	RaceTable *RaceTableStruct `json:"RaceTable,omitempty"`
}

type RaceTableStruct struct {
	Season *string `json:"season,omitempty"`
	Round  *string `json:"round,omitempty"`
	Races  []*Race `json:"Races,omitempty"`
}

type Race struct {
	Season   *string        `json:"season,omitempty"`
	Round    *string        `json:"round,omitempty"`
	Url      *string        `json:"url,omitempty"`
	RaceName *string        `json:"raceName,omitempty"`
	Circuit  *CircuitStruct `json:"Circuit,omitempty"`
	Date     *string        `json:"date,omitempty"`
	Time     *string        `json:"time,omitempty"`
}

func (r Race) GetFridayDate() string {
	t, err := time.Parse("2006-01-02", *r.Date)
	if err != nil {
		log.Println(err)
		return ""
	}
	date := fmt.Sprintf("%vT00:00:00Z", strings.Split(t.Add(-time.Hour*48).String(), " ")[0])
	return date
}

func (r Race) GetSaturdayDate() string {
	t, err := time.Parse("2006-01-02", *r.Date)
	if err != nil {
		log.Println(err)
		return ""
	}
	date := fmt.Sprintf("%vT00:00:00Z", strings.Split(t.Add(-time.Hour*24).String(), " ")[0])
	return date
}

func (r Race) GetDateTime() string {
	return fmt.Sprintf("%vT%v", *r.Date, *r.Time)
}

type CircuitStruct struct {
	CircuitId   *string         `json:"circuitId,omitempty"`
	Url         *string         `json:"url,omitempty"`
	CircuitName *string         `json:"circuitName,omitempty"`
	Location    *LocationStruct `json:"Location,omitempty"`
}

type LocationStruct struct {
	Lat      *string `json:"lat,omitempty"`
	Long     *string `json:"long,omitempty"`
	Locality *string `json:"locality,omitempty"`
	Country  *string `json:"country,omitempty"`
}
