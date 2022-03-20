package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"sync"
)

const (
	verificationMessageId = "936701233186623499"
)

func verificationMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	var waiter sync.WaitGroup

	waiter.Add(1)
	log.Println(waiter)
	message, err := sendVerifyMessage(s, m, &waiter)
	waiter.Wait()
	if err != nil {
		return err
	}
	log.Println("React message sent successfully")

	err = addReaction(s, m.ChannelID, message.ID, blueEmoji)
	if err != nil {
		return err
	}
	log.Printf("Emoji %v setup", blueEmoji)
	log.Println("Reaction setup")
	return nil
}

func sendVerifyMessage(s *discordgo.Session, m *discordgo.MessageCreate, waiter *sync.WaitGroup) (*discordgo.Message, error) {
	defer waiter.Done()

	var embed discordgo.MessageEmbed
	embed.Color = redEmbed
	embed.Title = verification.name
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		return nil, err
	}
	embed.Description = fmt.Sprintf(verification.response, guild.Name)

	log.Printf("Embedded message: \n%v", embed)

	message, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	if err != nil {
		return nil, err
	}
	log.Println("Embed message sent")

	if err := deleteMessage(s, m.ChannelID, m.ID); err != nil {
		return nil, err
	}
	log.Println("Command message deleted")
	return message, nil
}

func giveMemberRole(s *discordgo.Session, i *discordgo.MessageReactionAdd) error {
	var err error
	switch i.Emoji.Name {
	case blueEmoji:
		if err := addRole(s, i.GuildID, member.id, i.UserID); err != nil {
			return err
		}
	}

	log.Printf("%v given role", i.UserID)
	guild, err := s.Guild(i.GuildID)
	if err != nil {
		return err
	}
	if _, err := sendPrivateMessage(s, i.UserID, fmt.Sprintf(verificationsPassedDm, guild.Name)); err != nil {
		return err
	}
	return nil
}

func removeMemberRole(s *discordgo.Session, i *discordgo.MessageReactionRemove) error {
	switch i.Emoji.Name {
	case blueEmoji:
		if err := s.GuildMemberRoleRemove(i.GuildID, i.UserID, member.id); err != nil {
			return err
		}
	}
	log.Printf("%v role removed", i.UserID)
	guild, err := s.Guild(i.GuildID)
	if err != nil {
		return err
	}
	if _, err := sendPrivateMessage(s, i.UserID, fmt.Sprintf(verificationRemovedByUser, guild.Name)); err != nil {
		return err
	}
	return nil
}
