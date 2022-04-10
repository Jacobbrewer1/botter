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

func nextF1Race(s *discordgo.Session, channelId string) {
	race, err := api.GetNextRace()
	if err != nil {
		log.Println(err)
		return
	}
	t, err := time.Parse(helper.TimeFormatLayout, race.Session.GetSessionDateTime())
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := sendMessage(s, channelId, fmt.Sprintf(nextRace.response, *race.RaceName, t.Format(time.Stamp))); err != nil {
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
		t, err := time.Parse(helper.TimeFormatLayout, race.FirstPractice.GetSessionDateTime())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("first practice start %v\n", t)
		diff := helper.CalculateTimeDifferenceNow(t)
		if diff > 0 {
			log.Printf("waiting until %v event at %v\n", *race.RaceName, race.Session.GetSessionDateTime())
			time.Sleep(diff)
			log.Println("released from wait")
		} else {
			go runWeekend(s, race)
		}
		if diff < 0 {
			d, err := time.Parse(helper.TimeFormatLayout, race.Session.GetSessionDateTime())
			if err != nil {
				log.Println(err)
				continue
			}
			x := d.Add(time.Hour * 24)
			log.Printf("waiting until %v to get the next event\n", x)
			time.Sleep(helper.CalculateTimeDifferenceNow(x))
			log.Println("released from waiting for next event")
			go driverStandings(s, guildSportsChannel)
		}
	}
}

func runWeekend(session *discordgo.Session, raceStruct api.Race) {
	log.Println("running race weekend process")
	var waiter sync.WaitGroup
	waiter.Add(6)

	// First practice
	go func(s *discordgo.Session, r api.Race) {
		defer waiter.Done()

		if r.FirstPractice == nil {
			log.Printf("no first practice at %v\n", *r.RaceName)
			return
		}

		t, err := time.Parse(helper.TimeFormatLayout, r.FirstPractice.GetSessionDateTime())
		if err != nil {
			log.Println(err)
			return
		}
		t = t.Add(-time.Hour)
		diff := helper.CalculateTimeDifferenceNow(t)
		if diff > 0 {
			log.Printf("waiting for first practice at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseOneHourTime, firstPractice, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			t = t.Add(time.Hour)
			diff = helper.CalculateTimeDifferenceNow(t)
			time.Sleep(diff)
			msg, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseNow, firstPractice, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			log.Println("practice complete")
		}
	}(session, raceStruct)

	// Second practice
	go func(s *discordgo.Session, r api.Race) {
		defer waiter.Done()

		if r.SecondPractice == nil {
			log.Printf("no second practice at %v\n", *r.RaceName)
			return
		}

		t, err := time.Parse(helper.TimeFormatLayout, r.SecondPractice.GetSessionDateTime())
		if err != nil {
			log.Println(err)
			return
		}
		t = t.Add(-time.Hour)
		diff := helper.CalculateTimeDifferenceNow(t)
		if diff > 0 {
			log.Printf("waiting for second practice at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseOneHourTime, secondPractice, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			t = t.Add(time.Hour)
			diff = helper.CalculateTimeDifferenceNow(t)
			time.Sleep(diff)
			msg, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseNow, secondPractice, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			log.Println("practice complete")
		}
	}(session, raceStruct)

	// Third practice method
	go func(s *discordgo.Session, r api.Race) {
		defer waiter.Done()

		if r.ThirdPractice == nil {
			log.Printf("no third practice at %v\n", *r.RaceName)
			return
		}

		t, err := time.Parse(helper.TimeFormatLayout, r.ThirdPractice.GetSessionDateTime())
		if err != nil {
			log.Println(err)
			return
		}
		t = t.Add(-time.Hour)
		diff := helper.CalculateTimeDifferenceNow(t)
		if diff > 0 {
			log.Printf("waiting for third practice at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseOneHourTime, thirdPractice, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			t = t.Add(time.Hour)
			diff = helper.CalculateTimeDifferenceNow(t)
			time.Sleep(diff)
			msg, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseNow, thirdPractice, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			log.Println("practice complete")
		}
	}(session, raceStruct)

	// Qualifying method
	go func(s *discordgo.Session, r api.Race) {
		defer waiter.Done()

		if r.FirstPractice == nil {
			log.Printf("no qualifying at %v\n", *r.RaceName)
			return
		}

		t, err := time.Parse(helper.TimeFormatLayout, r.Qualifying.GetSessionDateTime())
		if err != nil {
			log.Println(err)
			return
		}
		t = t.Add(-time.Hour)
		diff := helper.CalculateTimeDifferenceNow(t)
		if diff > 0 {
			log.Printf("waiting for qualifying at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseOneHourTime, qualifying, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			t = t.Add(time.Hour)
			diff = helper.CalculateTimeDifferenceNow(t)
			time.Sleep(diff)
			msg, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseNow, qualifying, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}
			log.Println("qualifying complete")
		}
	}(session, raceStruct)

	// Sprint method
	go func(s *discordgo.Session, r api.Race) {
		defer waiter.Done()

		if r.Sprint == nil {
			log.Printf("no sprint race at %v\n", *r.RaceName)
			return
		}

		t, err := time.Parse(helper.TimeFormatLayout, r.Sprint.GetSessionDateTime())
		t = t.Add(-time.Hour)
		if err != nil {
			log.Println(err)
			return
		}

		diff := helper.CalculateTimeDifferenceNow(t)
		if diff > 0 {
			log.Printf("waiting for sprint race at %v\n", t)
			time.Sleep(diff)
			m, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseOneHourTime, sprint, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, m.ChannelID, m.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}

			t = t.Add(time.Hour)
			diff = helper.CalculateTimeDifferenceNow(t)
			time.Sleep(diff)
			msg, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(genericF1ResponseNow, sprint, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}
			log.Println("Race day complete")
		}
	}(session, raceStruct)

	// Racing method
	go func(s *discordgo.Session, r api.Race) {
		defer waiter.Done()

		if r.FirstPractice == nil {
			log.Printf("no race at %v\n", *r.RaceName)
			return
		}

		t, err := time.Parse(helper.TimeFormatLayout, r.Session.GetSessionDateTime())
		t = t.Add(-time.Hour)
		if err != nil {
			log.Println(err)
			return
		}
		diff := t.Sub(time.Now())
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
				return
			}

			t = t.Add(time.Hour)
			diff = helper.CalculateTimeDifferenceNow(t)
			time.Sleep(diff)
			msg, err := sendMessage(s, guildSportsChannel, fmt.Sprintf(raceDayResponseNow, *r.RaceName))
			if err != nil {
				log.Println(err)
				return
			}
			if err := addReaction(s, msg.ChannelID, msg.ID, racingCarEmoji); err != nil {
				log.Println(err)
				return
			}
			log.Println("Race day complete")
		}
	}(session, raceStruct)
	waiter.Wait()
}
