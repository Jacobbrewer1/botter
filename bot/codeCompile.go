package bot

import (
	"github.com/Jacobbrewer1/botter/api"
	"github.com/Jacobbrewer1/botter/config"
	"github.com/Jacobbrewer1/botter/helper"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type langStruct struct {
	languageCode string
	index        string
}

var (
	langMap = map[string]langStruct{
		"golang": golang,
	}

	golang = langStruct{
		languageCode: "go",
		index:        "4",
	}
)

func containsCode(msg string) bool {
	return strings.Count(msg, "```") == 2
}

func runCompile(s *discordgo.Session, m *discordgo.Message) {
	n := helper.RemoveNewLines(helper.RemoveTab(m.Content))
	n = n[strings.Index(n, "```")+3:]
	n = n[:strings.Index(n, "```")]
	slice := strings.Split(n, " ")
	language := slice[0]
	n = strings.Join(slice[1:], " ")
	log.Println("code extracted from message: " + n)
	result, err := api.RunCode(createCompileStruct(n, language))
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result recieved: %v\n", result)
	if _, err := sendMessage(s, m.ChannelID, *result.Output); err != nil {
		log.Println(err)
		return
	}
}

func compileHelp(s *discordgo.Session, m *discordgo.Message) {

}

func createCompileStruct(code, language string) api.ExecuteInput {
	s := langMap[language]
	return api.ExecuteInput{
		ClientId:     config.ApiSecrets.JdoodleApiClientId,
		ClientSecret: config.ApiSecrets.JdoodleApiClientSecret,
		Script:       &code,
		Language:     &s.languageCode,
		VersionIndex: &s.index,
	}
}

func reactToCodeMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if err := addReaction(s, m.ChannelID, m.ID, runReactionEmoji); err != nil {
		log.Println(err)
		return
	}
	if err := addReaction(s, m.ChannelID, m.ID, helpReactionEmoji); err != nil {
		log.Println(err)
		return
	}
}
