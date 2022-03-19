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
	var w sync.WaitGroup
	for _, mbr := range members {
		w.Add(1)
		go func(mem *discordgo.Member) {
			defer w.Done()
			log.Printf("clearing blue role from %v\n", mem.User.Username)
			if err = voidRole(s, m.GuildID, blueRole.id, mem.User.ID); err != nil {
				log.Println(err)
				return
			}
			log.Printf("clearing red role from %v\n", mem.User.Username)
			if err = voidRole(s, m.GuildID, redRole.id, mem.User.ID); err != nil {
				log.Println(err)
				return
			}
		}(mbr)
	}
	w.Wait()
	log.Println(rcrResponse)
	// TODO : When the server grows greater that 1000 members. There will have to be an if members > 1000 then get more members
	return nil
}
