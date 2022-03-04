package bot

import (
	"fmt"
	"github.com/Jacobbrewer1/botter/config"
	"github.com/bwmarrin/discordgo"
	"log"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "" {
		return
	}

	log.Printf("Message input: %v: %v", m.Author, m.Content)
	if memeChannel(m) {
		laughAtMessage(s, m)
	}

	if m.Author.ID == botId {
		return
	}

	if messageFilter(s, m) {
		return
	}

	if m.Content[0:1] == *config.JsonConfig.BotPrefix {
		log.Printf("Command received: %v", m.Content)
		botterProcess(s, m)
	}
	return
}

func reactionAddHandler(s *discordgo.Session, i *discordgo.MessageReactionAdd) {
	log.Println("Reaction received")
	runCustom, staticMessage := reactionAddFilter(i)
	if staticMessage {
		giveMemberRole(s, i)
	}
	if runCustom {
		giveReactionCustom(s, i)
	}
	return
}

func reactionRemoveHandler(s *discordgo.Session, i *discordgo.MessageReactionRemove) {
	log.Println("Reaction remove received")
	runCustom, staticMessage := reactionRemoveFilter(i)
	if staticMessage {
		err := removeMemberRole(s, i)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if runCustom {
		err := removeRoleCustom(s, i)
		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}

func memberJoinHandler(s *discordgo.Session, j *discordgo.GuildMemberAdd) {
	guild, err := s.Guild(j.GuildID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v just joined the server %v\n", j.User.Username, guild.Name)
	if _, err := sendPrivateMessage(s, j.User.ID, fmt.Sprintf(verificationJoinDm, j.User.Username, guild.Name)); err != nil {
		log.Println(err)
		return
	}
}
