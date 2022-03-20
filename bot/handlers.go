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

	log.Printf("message input: %v: %v", m.Author, m.Content)
	if memeChannel(m) {
		laughAtMessage(s, m)
	}

	if m.Author.ID == botId {
		log.Println("message is from botter")
		return
	}

	if messageFilter(s, m) {
		return
	}

	if containsCode(m.Content) {
		reactToCodeMessage(s, m)
	}

	if m.Content[0:1] == *config.JsonConfig.BotPrefix {
		log.Printf("command received: %v", m.Content)
		botterProcess(s, m)
	}
	return
}

func reactionAddHandler(s *discordgo.Session, i *discordgo.MessageReactionAdd) {
	log.Println("reaction received")

	if i.UserID == botId {
		log.Println("botter reaction")
		return
	}

	runCustom, staticMessage := reactionAddFilter(i)
	if staticMessage {
		giveMemberRole(s, i)
	}
	if runCustom {
		giveReactionCustom(s, i)
	}
	m, err := s.ChannelMessage(i.ChannelID, i.MessageID)
	if err != nil {
		log.Println(err)
		return
	}
	if containsCode(m.Content) {
		if i.Emoji.APIName() == runReactionEmoji {
			if err := removeSingleEmojiReaction(s, i.ChannelID, i.MessageID, runReactionEmoji); err != nil {
				log.Println(err)
				return
			}
			runCompile(s, m)
		} else if i.Emoji.APIName() == helpReactionEmoji {
			if err := removeSingleEmojiReaction(s, i.ChannelID, i.MessageID, helpReactionEmoji); err != nil {
				log.Println(err)
				return
			}
			compileHelp(s, i)
		}
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
