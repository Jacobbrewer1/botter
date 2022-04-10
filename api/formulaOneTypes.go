package api

import (
	"fmt"
	"github.com/Jacobbrewer1/botter/helper"
	"log"
	"strconv"
	"strings"
	"time"
)

type (
	GrandPixMRDataStruct struct {
		*GrandPixStruct `json:"MRData,omitempty"`
	}

	BaseStruct struct {
		Xmlns  *string `json:"xmlns,omitempty"`
		Series *string `json:"series,omitempty"`
		Url    *string `json:"url,omitempty"`
		Limit  *string `json:"limit,omitempty"`
		Offset *string `json:"offset,omitempty"`
		Total  *string `json:"total,omitempty"`
	}

	GrandPixStruct struct {
		*BaseStruct
		RaceTable *RaceTableStruct `json:"RaceTable,omitempty"`
	}

	RaceTableStruct struct {
		Season *string `json:"season,omitempty"`
		Round  *string `json:"round,omitempty"`
		Races  []*Race `json:"Races,omitempty"`
	}

	Race struct {
		Season         *string        `json:"season,omitempty"`
		Round          *string        `json:"round,omitempty"`
		Url            *string        `json:"url,omitempty"`
		RaceName       *string        `json:"raceName,omitempty"`
		Circuit        *CircuitStruct `json:"Circuit,omitempty"`
		FirstPractice  *Session       `json:"FirstPractice,omitempty"`
		SecondPractice *Session       `json:"SecondPractice,omitempty"`
		ThirdPractice  *Session       `json:"ThirdPractice,omitempty"`
		Qualifying     *Session       `json:"Qualifying,omitempty"`
		*Session
	}

	Session struct {
		Date *string `json:"date,omitempty"`
		Time *string `json:"time,omitempty"`
	}

	CircuitStruct struct {
		CircuitId   *string         `json:"circuitId,omitempty"`
		Url         *string         `json:"url,omitempty"`
		CircuitName *string         `json:"circuitName,omitempty"`
		Location    *LocationStruct `json:"Location,omitempty"`
	}

	LocationStruct struct {
		Lat      *string `json:"lat,omitempty"`
		Long     *string `json:"long,omitempty"`
		Locality *string `json:"locality,omitempty"`
		Country  *string `json:"country,omitempty"`
	}

	DriverStandingsMRDataStruct struct {
		*DriverStandingsStruct `json:"MRData,omitempty"`
	}

	DriverStandingsStruct struct {
		*BaseStruct
		StandingsTable *DriverStandingsTableStruct `json:"StandingsTable,omitempty"`
	}

	DriverStandingsTableStruct struct {
		Season         *string                       `json:"season,omitempty"`
		StandingsLists []*DriverStandingsListsStruct `json:"StandingsLists,omitempty"`
	}

	DriverStandingsListsStruct struct {
		Season          *string                         `json:"season,omitempty"`
		Round           *string                         `json:"round,omitempty"`
		DriverStandings []*DriverStandingPositionStruct `json:"DriverStandings,omitempty"`
	}

	DriverStandingPositionStruct struct {
		*StatsStruct
		Driver       *DriverStruct         `json:"Driver,omitempty"`
		Constructors []*ConstructorsStruct `json:"constructors,omitempty"`
	}

	StatsStruct struct {
		Position     *string `json:"position,omitempty"`
		PositionText *string `json:"positionText,omitempty"`
		Points       *string `json:"points,omitempty"`
		Wins         *string `json:"wins,omitempty"`
	}

	DriverStruct struct {
		DriverId        *string `json:"driverId,omitempty"`
		PermanentNumber *string `json:"permanentNumber,omitempty"`
		Code            *string `json:"code,omitempty"`
		Url             *string `json:"url,omitempty"`
		GivenName       *string `json:"givenName,omitempty"`
		FamilyName      *string `json:"familyName,omitempty"`
		DateOfBirth     *string `json:"dateOfBirth,omitempty"`
		Nationality     *string `json:"nationality,omitempty"`
	}

	ConstructorsStruct struct {
		ConstructorId *string `json:"constructorId,omitempty"`
		Url           *string `json:"url,omitempty"`
		Name          *string `json:"name,omitempty"`
		Nationality   *string `json:"nationality,omitempty"`
	}

	ConstructorStandingsMRDataStruct struct {
		*ConstructorStandingsStruct `json:"MRData,omitempty"`
	}

	ConstructorStandingsStruct struct {
		*BaseStruct
		StandingsTable *ConstructorSeasonTable `json:"StandingsTable,omitempty"`
	}

	ConstructorSeasonTable struct {
		Season         *string                        `json:"season,omitempty"`
		StandingsLists []*ConstructorSeasonListStruct `json:"StandingsLists,omitempty"`
	}

	ConstructorSeasonListStruct struct {
		Season               *string                               `json:"season,omitempty"`
		Round                *string                               `json:"round,omitempty"`
		ConstructorStandings []*ConstructorStandingsPositionStruct `json:"ConstructorStandings,omitempty"`
	}

	ConstructorStandingsPositionStruct struct {
		*StatsStruct
		Constructor *ConstructorsStruct `json:"Constructor"`
	}
)

func (s Session) GetSessionDateTime() string {
	zone, offset := time.Now().Zone()
	if zone == "GMT" || zone == "UTC" {
		return fmt.Sprintf("%vT%v", *s.Date, *s.Time)
	}
	u, err := time.Parse("15:04:05Z", *s.Time)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("%vT%v", *s.Date, *s.Time)
	}
	dur := time.Duration(offset) * time.Second
	u = u.Add(dur)
	t := u.Format(helper.TimeFormatLayout)
	return fmt.Sprintf("%vT%v", *s.Date, strings.Split(t, "T")[1])
}

func (s DriverStandingsListsStruct) GetPosition(p int) *DriverStandingPositionStruct {
	position := strconv.Itoa(p)
	for _, i := range s.DriverStandings {
		if *i.Position == position {
			return i
		}
	}
	return &DriverStandingPositionStruct{}
}

func (s ConstructorSeasonListStruct) GetPosition(p int) *ConstructorStandingsPositionStruct {
	position := strconv.Itoa(p)
	for _, i := range s.ConstructorStandings {
		if *i.Position == position {
			return i
		}
	}
	return &ConstructorStandingsPositionStruct{}
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

func (r Race) GetSessionDateTime() string {
	return fmt.Sprintf("%vT%v", *r.Date, *r.Time)
}
