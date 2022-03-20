package bot

import (
	"fmt"
	"github.com/Jacobbrewer1/botter/config"
	"github.com/Jacobbrewer1/botter/helper"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

func messageFilter(s *discordgo.Session, m *discordgo.Message) bool {
	messageUser := func(s *discordgo.Session, m *discordgo.Message, response string, timeout ...bool) error {
		guild, err := s.Guild(m.GuildID)
		if err != nil {
			return err
		}
		if timeout != nil {
			if timeout[0] {
				until := time.Now().Add(time.Minute * time.Duration(5))
				err := s.GuildMemberTimeout(m.GuildID, m.Author.ID, &until)
				if err != nil {
					if _, err := sendPrivateMessage(s, m.Author.ID, fmt.Sprintf(failedTimeoutDm, guild.Name)); err != nil {
						return err
					}
					return err
				}
				if _, err := sendPrivateMessage(s, m.Author.ID, fmt.Sprintf(response, guild.Name, until.Format("15:04 - 2 January, 2006"))); err != nil {
					return err
				}
			}
		} else {
			if _, err := sendPrivateMessage(s, m.Author.ID, fmt.Sprintf(response, guild.Name)); err != nil {
				return err
			}
		}
		log.Printf("Messaged %v about breach", m.Author.Username)
		return nil
	}
	result := make(chan bool, 2)
	defer close(result)
	go func() {
		log.Println("Running invite check")
		if containsServerInvite(m.Content) && !config.IgnoreInvites {
			// If the message is a DM do not proceed!
			if m.GuildID == "" {
				result <- true
				return
			}
			log.Println("Server invite detected")
			if err := deleteMessage(s, m.ChannelID, m.ID); err != nil {
				log.Println(err)
				result <- false
				return
			}
			log.Println("Server invite deleted")
			if err := messageUser(s, m, messageIsServerInviteDm); err != nil {
				log.Println(err)
				result <- true
				return
			}
			log.Println("invite check failed")
			result <- true
			return
		}
		log.Println("invite check passed")
		result <- false
		return
	}()
	go func() {
		log.Println("Running Banned word check")
		if bannedWordsFilter(m.Content) && !config.IgnoreBadWords {
			// If the message is a DM do not proceed!
			if m.GuildID == "" {
				if _, err := sendPrivateMessage(s, m.Author.ID, "That's rude!"); err != nil {
					log.Println(err)
				}
				result <- true
				return
			}
			log.Println("banned word detected")
			if err := deleteMessage(s, m.ChannelID, m.ID); err != nil {
				log.Println(err)
				result <- false
				return
			}
			if err := messageUser(s, m, messageContainedBannedWordDm, true); err != nil {
				log.Printf("Timout error: %v\n", err)
			}
			log.Println("banned word check failed")
			result <- true
			return
		}
		log.Println("banned word check passed")
		result <- false
		return
	}()
	if <-result {
		return true
	} else {
		if <-result {
			return true
		}
	}
	return false
}

func containsServerInvite(message string) bool {
	if strings.Contains(message, "https://discord.gg/") {
		return true
	}
	return false
}

func bannedWordsFilter(message string) bool {
	compareIndividualWords := func(word, text string) bool {
		array := strings.Split(message, " ")
		for _, elm := range array {
			if word == elm {
				return true
			}
			if (word + "er") == elm {
				return true
			}
			if (word + "ing") == elm {
				return true
			}
		}
		return false
	}
	specificWordDetermine := func(word string) bool {
		// add here for words like ass where they can be used in words like grass
		if word == "ass" {
			return true
		}
		return false
	}
	message = strings.ToLower(helper.RemoveNonAlphaChars(message))
	for _, banWord := range config.BannedWordsArray {
		if strings.Contains(message, banWord) {
			if specificWordDetermine(banWord) {
				// For performance, only run the loop over the string if it contains suspect text
				if compareIndividualWords(banWord, message) {
					return true
				}
			} else {
				return true
			}
		}
	}
	return false
}

func reactionAddFilter(i *discordgo.MessageReactionAdd) (bool, bool) {
	if i.MessageID == verificationMessageId {
		log.Println("Reaction to verification message")
		return false, true
	}
	if roleReactionIdCustom == "" {
		log.Println("roleReactionIdCustom is empty")
		return false, false
	}
	if i.MessageID != roleReactionIdCustom {
		log.Println("Poll message, do not give reaction")
		return false, false
	}
	return true, false
}

func reactionRemoveFilter(i *discordgo.MessageReactionRemove) (bool, bool) {
	if i.UserID == botId {
		log.Println("Botter reaction")
		return false, false
	}
	if i.MessageID == verificationMessageId {
		log.Println("Reaction removed from verification message")
		return false, true
	}
	if roleReactionIdCustom == "" {
		log.Println("roleReactionIdCustom is empty")
		return false, false
	}
	return true, false
}
