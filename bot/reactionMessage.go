package bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
	"sync"
)

func reactionMessage(s *discordgo.Session, m *discordgo.MessageCreate, cmd command) error {
	var waiter sync.WaitGroup

	waiter.Add(1)
	log.Println(waiter)
	message, err := sendReactionMessage(s, m.ChannelID, cmd.response, &waiter)
	waiter.Wait()
	if err != nil {
		return err
	}

	if err := deleteMessage(s, m); err != nil {
		return err
	}

	log.Println("React message sent successfully")

	if cmd == roleReactCustom {
		roleReactionIdCustom = message.ID
		log.Printf("roleReactionIdCustom now set to: %v", roleReactionIdCustom)
	}

	err = addReaction(s, m.ChannelID, message.ID, blueEmoji)
	if err != nil {
		return err
	}
	log.Printf("Emoji %v setup", blueEmoji)

	err = addReaction(s, m.ChannelID, message.ID, redEmoji)
	if err != nil {
		return err
	}
	log.Printf("Emoji %v setup", redEmoji)
	log.Println("Reactions setup")
	return nil
}

func giveReactionCustom(s *discordgo.Session, i *discordgo.MessageReactionAdd) error {
	switch i.Emoji.Name {
	case redEmoji:
		return addRole(s, i.GuildID, i.UserID, redRole.id)
	case blueEmoji:
		return addRole(s, i.GuildID, i.UserID, blueRole.id)
	}
	return errors.New("unknown emoji")
}

func removeRoleCustom(s *discordgo.Session, i *discordgo.MessageReactionRemove) error {
	switch i.Emoji.Name {
	case redEmoji:
		return voidRole(s, i.GuildID, i.UserID, redRole.id)
	case blueEmoji:
		return voidRole(s, i.GuildID, i.UserID, blueRole.id)
	}
	return errors.New("unknown emoji")
}

// Function to reset custom roles
func rcrFunc(s *discordgo.Session, m *discordgo.MessageCreate) error {
	roleReactionIdCustom = ""
	members, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		return err
	}
	for _, mbr := range members {
		if err = voidRole(s, m.GuildID, blueRole.id, mbr.User.ID); err != nil {
			return err
		}
		err = voidRole(s, m.GuildID, redRole.id, mbr.User.ID)
		if err != nil {
			return err
		}
	}
	log.Println(rcrResponse)
	// TODO : When the server grows greater that 1000 members. There will have to be an if members > 1000 then get more members
	return nil
}
