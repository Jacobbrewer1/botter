package bot

import (
	"fmt"
	"github.com/Jacobbrewer1/botter/api"
	"github.com/Jacobbrewer1/botter/helper"
	"github.com/bwmarrin/discordgo"
	"log"
	"sync"
	"time"
)

var (
	layout = "2006-01-02T15:04:05Z"
)

func nextF1Race(s *discordgo.Session, channelId string) {
	race, err := api.GetNextRace()
	if err != nil {
		log.Println(err)
		return
	}
	t, err := time.Parse(layout, race.GetDateTime())
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := sendMessage(s, channelId, fmt.Sprintf(nextRace.response, *race.RaceName, t.Format(time.RFC1123))); err != nil {
		log.Println(err)
		return
	}
	return
}

func getBothStandings(s *discordgo.Session, channelId string) {
	go driverStandings(s, channelId)
	go constructorStandings(s, channelId)
}

func constructorStandings(s *discordgo.Session, channelId string) {
	log.Println("constructor standings requested")
	standings, err := api.GetConstructorStandings()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("got constructor standings")
	log.Println("creating message for embed")
	text := fmt.Sprintf("Season - %v", *standings.StandingsTable.Season)
	for _, d := range standings.StandingsTable.StandingsLists[0].ConstructorStandings {
		text = text + "\n"
		tmp := fmt.Sprintf("%v - %v - Points: %v", *d.Position, *d.Constructor.Name, *d.Points)
		text = text + tmp
	}
	log.Println("message created")
	log.Println("creating embed")
	var e = discordgo.MessageEmbed{
		Title:       "Formula 1 Constructor Standings",
		Description: text,
		Color:       redEmbed,
	}
	log.Println("embed created")
	log.Println("sending message")
	msg, err := s.ChannelMessageSendEmbed(channelId, &e)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("message sent")
	if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
		log.Println(err)
		return
	}
	log.Println("reaction added")
}

func driverStandings(s *discordgo.Session, channelId string) {
	log.Println("driver standings requested")
	standings, err := api.GetDriverStandings()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("got driver standings")
	log.Println("creating message for embed")
	text := fmt.Sprintf("Season - %v", *standings.StandingsTable.Season)
	for _, d := range standings.StandingsTable.StandingsLists[0].DriverStandings {
		text = text + "\n"
		tmp := fmt.Sprintf("%v - %v - Points: %v", *d.Position, *d.Driver.FamilyName, *d.Points)
		text = text + tmp
	}
	log.Println("message created")
	log.Println("creating embed")
	var e = discordgo.MessageEmbed{
		Title:       "Formula 1 Driver Standings",
		Description: text,
		Color:       redEmbed,
	}
	log.Println("embed created")
	log.Println("sending message")
	msg, err := s.ChannelMessageSendEmbed(channelId, &e)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("message sent")
	if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
		log.Println(err)
		return
	}
	log.Println("reaction added")
}

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
		diff := helper.CalculateTimeDifference(t, time.Now().UTC())
		if diff > 0 {
			log.Printf("waiting until %v event at %v\n", *race.RaceName, race.GetFridayDate())
			time.Sleep(diff)
			log.Println("released from wait")
		} else {
			go runWeekend(s, race)
		}
		if diff < 0 {
			d, err := time.Parse(layout, race.GetDateTime())
			if err != nil {
				log.Println(err)
				continue
			}
			x := d.Add(time.Hour * 24)
			log.Printf("waiting until %v to get the next event\n", x)
			time.Sleep(helper.CalculateTimeDifference(x, d))
			log.Println("released from waiting for next event")
			go driverStandings(s, guildSportsChannel)
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
		t = t.Add(time.Hour * 9)
		diff := t.Sub(time.Now())
		if diff > 0 {
			log.Printf("waiting for practice at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1Response, practice, *r.Circuit.CircuitName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
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
		t = t.Add(time.Hour * 9)
		diff := t.Sub(time.Now())
		if diff > 0 {
			log.Printf("waiting for qualifying at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1Response, qualifying, *r.Circuit.CircuitName))
			if err != nil {
				log.Println(err)
				return
			}

			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
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
		diff := t.Sub(time.Now().UTC())
		if diff > 0 {
			log.Printf("waiting for race day at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(raceDayResponse, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}

			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
			}
			log.Println("Race day complete")
		}
	}()
	waiter.Wait()
}
