package bot

import (
	"fmt"
	"github.com/Jacobbrewer1/botter/config"
	"github.com/bwmarrin/discordgo"
	"log"
)

var botId string

func Start() {
	bot, err := discordgo.New("Bot " + *config.ApiSecrets.BotToken)
	if err != nil {
		log.Println(err.Error())
		return
	}

	u, err := bot.User("@me")
	if err != nil {
		log.Println(err.Error())
		return
	}

	botId = u.ID

	bot.AddHandler(messageHandler)
	bot.AddHandler(reactionAddHandler)
	bot.AddHandler(reactionRemoveHandler)
	bot.AddHandler(memberJoinHandler)

	bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err = bot.Open()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Bot is running!")

	//err = bot.UpdateListeningStatus("Spotify")
	//err = bot.UpdateGameStatus(0, "@botter help")
	//err = bot.UpdateStreamingStatus(0, "@botter help", "https://www.bbc.com")
	err = bot.UpdateStatusComplex(setupComplexStatus())
	if err != nil {
		log.Println(err)
	}
}

func setupComplexStatus() discordgo.UpdateStatusData {
	return discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: fmt.Sprintf("@%v help", botName),
				Type: discordgo.ActivityTypeWatching,
				URL:  "",
			},
		},
	}
}
