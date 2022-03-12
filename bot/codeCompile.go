package bot

import (
	"github.com/Jacobbrewer1/botter/helper"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func containsCode(msg string) bool {
	return strings.Count(msg, "```") == 2
}

func runCompile(s *discordgo.Session, m *discordgo.Message) {
	n := helper.RemoveNewLines(helper.RemoveTab(m.Content))
	n = n[strings.Index(n, "```")+3:]
	n = n[:strings.Index(n, "```")]
	log.Println(n)
}

func reactToCodeMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := s.MessageReactionAdd(m.ChannelID, m.ID, runReactionEmoji)
	if err != nil {
		log.Println(err)
	}
}
