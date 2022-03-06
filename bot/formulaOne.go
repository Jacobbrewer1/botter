package bot

import (
	"fmt"
	"github.com/Jacobbrewer1/botter/api"
	"github.com/bwmarrin/discordgo"
	"log"
	"sync"
	"time"
)

var (
	layout = "2006-01-02T15:04:05Z"
)

// Should be run as a go routine to allow it to run independently to the rest of the bot
// i.e. go runFormulaOne
func runFormulaOne(s *discordgo.Session) {
	log.Println("formula 1 go routine started")
	for {
		race, err := api.GetNextRace()
		if err != nil {
			log.Println(err)
			continue
		}
		t, err := time.Parse(layout, race.GetFridayDate())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("race weekend start date %v\n", t)
		diff := calculateTimeDifference(t)
		if diff > 0 {
			log.Printf("waiting until %v event at %v\n", *race.RaceName, race.GetFridayDate())
			time.Sleep(diff)
			log.Println("released from wait")
		} else if diff < 0 {
			log.Println("waiting until the next event to come up")
			time.Sleep(time.Hour * 6)
			log.Println("released from waiting for next event")
		} else {
			runWeekend(s, race)
		}
	}
}

func runWeekend(s *discordgo.Session, r api.Race) {
	log.Println("running race weekend process")
	var waiter sync.WaitGroup
	waiter.Add(3)
	go func() {
		defer waiter.Done()
		t, err := time.Parse(layout, r.GetFridayDate())
		if err != nil {
			log.Println(err)
			return
		}
		t.Add(time.Hour * 9)
		diff := t.Sub(time.Now())
		if diff > 0 {
			log.Println("waiting for practice")
			time.Sleep(diff)
			if _, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1Response, practice, *r.Circuit.CircuitName)); err != nil {
				log.Println(err)
			}
			log.Println("practice complete")
		}
	}()
	go func() {
		defer waiter.Done()
		t, err := time.Parse(layout, r.GetSaturdayDate())
		if err != nil {
			log.Println(err)
			return
		}
		t.Add(time.Hour * 9)
		diff := t.Sub(time.Now())
		if diff > 0 {
			log.Println("waiting for qualifying")
			time.Sleep(diff)
			if _, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1Response, qualifying, *r.Circuit.CircuitName)); err != nil {
				log.Println(err)
			}
			log.Println("qualifying complete")
		}
	}()
	go func() {
		defer waiter.Done()
		t, err := time.Parse(layout, r.GetDateTime())
		t = t.Add(-time.Hour)
		if err != nil {
			log.Println(err)
			return
		}
		t.Add(time.Hour * 9)
		diff := t.Sub(time.Now())
		if diff > 0 {
			log.Println("waiting for Race day")
			time.Sleep(diff)
			if _, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(raceDayResponse, *r.RaceName)); err != nil {
				log.Println(err)
			}
			log.Println("Race day complete")
		}
	}()
	waiter.Wait()
}

func calculateTimeDifference(t time.Time) time.Duration {
	diff := t.Sub(time.Now())
	log.Println("f1 time difference ", diff)
	if diff < time.Hour*24 {
		return 0
	}
	return diff
}
