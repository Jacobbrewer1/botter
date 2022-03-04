package bot

import "github.com/bwmarrin/discordgo"

func memeChannel(m *discordgo.MessageCreate) bool {
	if m.ChannelID == memeChannelId {
		return true
	}
	return false
}

func laughAtMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	err := s.MessageReactionAdd(m.ChannelID, m.ID, joyEmoji)
	if err != nil {
		return err
	}
	return nil
}
