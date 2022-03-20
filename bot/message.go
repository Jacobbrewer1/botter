package bot

import (
	"errors"
	"fmt"
	"github.com/Jacobbrewer1/botter/config"
	"github.com/Jacobbrewer1/botter/helper"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"sync"
)

type messageStruct struct {
	cmd   string
	query string
}

func (m *messageStruct) isEmpty() bool {
	return m.query == ""
}

var message = messageStruct{}

func botterProcess(s *discordgo.Session, m *discordgo.MessageCreate) {
	message.cmd, message.query = helper.FormatMessage(m.Content, *config.JsonConfig.BotPrefix)
	switch message.cmd {
	case help.trigger:
		if _, err := sendMessage(s, m.ChannelID, help.response); err != nil {
			log.Println(err)
			break
		}
		break
	case bossCommand.trigger:
		if _, err := sendMessage(s, m.ChannelID, bossCommand.response); err != nil {
			log.Println(err)
			break
		}
		break
	case mrsBossCommand.trigger:
		if _, err := sendMessage(s, m.ChannelID, mrsBossCommand.response); err != nil {
			log.Println(err)
			break
		}
		break
	case hello.trigger:
		if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(hello.response, m.Author.ID)); err != nil {
			log.Println(err)
			break
		}
		break
	case hi.trigger:
		if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(hi.response, m.Author.ID)); err != nil {
			log.Println(err)
			break
		}
		break
	case hey.trigger:
		if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(hey.response, m.Author.ID)); err != nil {
			log.Println(err)
			break
		}
		break
	case ping.trigger:
		if _, err := sendMessage(s, m.ChannelID, ping.response); err != nil {
			log.Println(err)
			break
		}
		break
	case laugh.trigger:
		if err := deleteMessage(s, m); err != nil {
			log.Println(err)
			break
		}
		if _, err := sendMessage(s, m.ChannelID, laugh.response); err != nil {
			log.Println(err)
			break
		}
		break
	case minecraftHomeCoordinates.trigger:
		var embed discordgo.MessageEmbed
		embed.Title = minecraftHomeCoordinates.name
		embed.Color = greenEmbed
		embed.Description = minecraftHomeCoordinates.response

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		if err != nil {
			log.Println(err)
			break
		}
		break
	case verification.trigger:
		if verifier(s, m, moderator) {
			if err := verificationMessage(s, m); err != nil {
				log.Println(err)
				break
			}
		}
		break
	case minecraftBrewing.trigger:
		if _, err := sendMessage(s, m.ChannelID, minecraftBrewing.response); err != nil {
			log.Println(err)
			break
		}
		break
	case mcPigRide.trigger:
		if _, err := sendTextJoke(s, m, mcPigRide.response); err != nil {
			log.Println(err)
			break
		}
		break
	case roleReactCustom.trigger:
		if verifier(s, m, moderator) {
			go func() {
				log.Println("rcr go routine started")
				if err := rcrFunc(s, m); err != nil {
					log.Println(err)
					return
				}
				log.Println("rcr go routine completed")
				return
			}()
			go func() {
				log.Println("reactionMessage go routine started")
				if err := reactionMessage(s, m, roleReactCustom); err != nil {
					log.Println(err)
					if _, err := sendReplyMessage(s, m, roleReactCustom.response); err != nil {
						log.Println(err)
						return
					}
					return
				}
				log.Println("reactionMessage go routine completed")
				return
			}()
		}
		break
	case actions.trigger:
		if err := sendActions(s, m); err != nil {
			log.Println(err)
			break
		}
		break
	case poll.trigger:
		if err := reactionMessage(s, m, poll); err != nil {
			log.Println(err)
			if _, err := sendReplyMessage(s, m, poll.response); err != nil {
				log.Println(err)
				break
			}
			break
		}
		break
	case gifCommand.trigger:
		if err := gifMessage(s, m); err != nil {
			log.Println(err)
			break
		}
		break
	case grant.trigger:
		if success, err := roleHandler(s, m, moderator, true); !success {
			if err != nil {
				log.Println(err)
			}
			if _, err := sendReplyMessage(s, m, grant.response); err != nil {
				log.Println(err)
			}
			break
		}
		break
	case void.trigger:
		if success, err := roleHandler(s, m, moderator, false); !success {
			if err != nil {
				log.Println(err)
			}
			if _, err := sendReplyMessage(s, m, void.response); err != nil {
				log.Println(err)
			}
			break
		}
		break
	case serverAdTemplateCommand.trigger:
		if !verifier(s, m, moderator, false) {
			break
		}
		invite, err := generateInvite(s)
		if err != nil {
			log.Println(err)
			break
		}
		if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(serverAdTemplateCommand.response, invite)); err != nil {
			log.Println(err)
			break
		}
		break
	case invite.trigger:
		if !verifier(s, m, member, true) {
			break
		}
		if strings.Contains(message.query, "<@") || len(strings.Split(message.query, " ")) > 1 {
			if message.isEmpty() || strings.Contains(message.query, " ") {
				if _, err := sendReplyMessage(s, m, inviteFormatResponse); err != nil {
					log.Println(err)
					break
				}
				break
			}
			query, err := helper.FormatIdFromMention(message.query)
			if err != nil {
				log.Println(err)
				break
			}
			user, err := s.GuildMember(m.GuildID, query)
			if err != nil {
				log.Println(err)
				break
			}
			guild, err := s.Guild(m.GuildID)
			if err != nil {
				log.Println(err)
				break
			}
			invMessage, err := generateInvite(s)
			if err != nil {
				log.Println(err)
				break
			}
			if _, err = sendPrivateMessage(s, user.User.ID, fmt.Sprintf(invite.response, guild.Name, m.Author.Username, invMessage)); err != nil {
				log.Println(err)
				break
			}
			if _, err = sendReplyMessage(s, m, fmt.Sprintf(invite.secondResponse, user.User.Username)); err != nil {
				log.Println(err)
				break
			}
		} else {
			invite, err := generateInvite(s)
			if err != nil {
				log.Println(err)
				break
			}
			if _, err := sendReplyMessage(s, m, invite); err != nil {
				log.Println(err)
				break
			}
		}
		break
	case resetCustomRoles.trigger:
		if !verifier(s, m, moderator) {
			break
		}
		go func() {
			if err := rcrFunc(s, m); err != nil {
				log.Println(err)
				return
			}
			if _, err := sendReplyMessage(s, m, rcrResponse); err != nil {
				log.Println(err)
				return
			}
			return
		}()
		break
	case stickerCommand.trigger:
		if err := stickerMessage(s, m); err != nil {
			log.Println(err)
			break
		}
		break
	case issue.trigger:
		if !verifier(s, m, member) {
			break
		}
		if err := createGithubIssue(s, m); err != nil {
			log.Println(err)
			if _, err := sendMessage(s, m.ChannelID, issue.secondResponse); err != nil {
				log.Println(err)
				break
			}
			break
		}
		break
	case listIssues.trigger:
		if err := getGithubIssues(s, m); err != nil {
			log.Println(err)
			break
		}
		break
	case driverStandingsCommand.trigger:
		go driverStandings(s, m.ChannelID)
		break
	default:
		log.Println(unknownResponse)
		if _, err := sendMessage(s, m.ChannelID, unknownResponse); err != nil {
			log.Println(err)
			break
		}
		break
	}
}

func sendMessage(s *discordgo.Session, channelId, message string, trailer ...string) (*discordgo.Message, error) {
	botMessage, err := s.ChannelMessageSend(channelId, message)
	if trailer != nil {
		log.Printf("Trailer message: %v\n", trailer)
		_, err = s.ChannelMessageSend(channelId, trailer[0])
	}
	return botMessage, err
}

func sendWaiterMessage(s *discordgo.Session, m *discordgo.MessageCreate, message string, waiter *sync.WaitGroup) (*discordgo.Message, error) {
	defer waiter.Done()
	return sendMessage(s, m.ChannelID, message)
}

func sendReactionMessage(s *discordgo.Session, channelId, response string, waiter *sync.WaitGroup) (*discordgo.Message, error) {
	defer waiter.Done()

	var embed discordgo.MessageEmbed
	embed.Color = 15158332

	mess := strings.Split(message.query, "-")

	if len(mess) > 2 {
		return nil, errors.New(response)
	}
	embed.Title = strings.TrimSpace(mess[0])
	if len(mess) == 2 {
		embed.Description = strings.TrimSpace(mess[1])
	}

	log.Printf("Embedded message: \n%v", embed)

	return s.ChannelMessageSendEmbed(channelId, &embed)
}

func sendPrivateMessage(s *discordgo.Session, userId string, message string) (*discordgo.Message, error) {
	channel, err := s.UserChannelCreate(userId)
	if err != nil {
		return nil, err
	}
	return s.ChannelMessageSend(channel.ID, message)
}

func sendReplyMessage(s *discordgo.Session, m *discordgo.MessageCreate, message string) (*discordgo.Message, error) {
	var reply discordgo.MessageReference
	reply.MessageID = m.ID
	reply.GuildID = m.GuildID
	reply.ChannelID = m.ChannelID
	return s.ChannelMessageSendReply(m.ChannelID, message, &reply)
}

func sendTextJoke(s *discordgo.Session, m *discordgo.MessageCreate, text string) (*discordgo.Message, error) {
	if err := deleteMessage(s, m); err != nil {
		return nil, err
	}
	message, err := s.ChannelMessageSend(m.ChannelID, text)
	if err != nil {
		return message, err
	}
	return message, s.MessageReactionAdd(message.ChannelID, message.ID, joyEmoji)
}

func sendActions(s *discordgo.Session, m *discordgo.MessageCreate) error {
	setupActions := func() string {
		var actionString string
		for _, cmd := range commands {
			actionString = actionString + cmd.name + fmt.Sprintf(" = '%v'\n", cmd.trigger)
		}
		return actionString
	}
	var embed discordgo.MessageEmbed
	embed.Title = actions.name
	embed.Color = blueEmbed
	embed.Description = setupActions()

	message, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	if err != nil {
		return err
	}
	log.Println("Actions sent")
	err = s.MessageReactionAdd(m.ChannelID, message.ID, faceWithTongueEmoji)
	if err != nil {
		return err
	}
	log.Println("Reaction to actions message complete")
	return err
}

func deleteMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		return err
	}
	log.Println("Message deleted")
	return err
}

func generateInvite(s *discordgo.Session) (string, error) {
	var i discordgo.Invite
	i.MaxAge = 604800
	i.MaxUses = 100
	i.Temporary = false
	i.Unique = false
	invite, err := s.ChannelInviteCreate(regulationsChannel, i)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://discord.gg/%v", invite.Code), err
}

func addReaction(s *discordgo.Session, channelId, msgId, emoji string) error {
	log.Printf("adding reaction %v to message %v\n", emoji, msgId)
	return s.MessageReactionAdd(channelId, msgId, emoji)
}

func removeSingleEmojiReaction(s *discordgo.Session, channelId, msgId, emoji string) error {
	log.Printf("removing emoji %v from message %v\n", emoji, msgId)
	return s.MessageReactionsRemoveEmoji(channelId, msgId, emoji)
}

func removeAllMessageReactions(s *discordgo.Session, channelId, msgId string) error {
	log.Println("removing all emojis from message " + msgId)
	return s.MessageReactionsRemoveAll(channelId, msgId)
}
