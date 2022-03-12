package bot

import (
	"fmt"
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
	r := *result.Output
	r = r[:len(r)-1] // Removes the new line added by the response by Jdoodle
	var embed discordgo.MessageEmbed
	embed.Title = "Code Results"
	embed.Color = blueEmbed
	embed.Description = fmt.Sprintf(`Output: %v
CPU Time: %v
Memory Used %v`, r, *result.CpuTime, *result.Memory)
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed); err != nil {
		log.Println(err)
		return
	}
}

func compileHelp(s *discordgo.Session, i *discordgo.MessageReactionAdd) {
	if _, err := sendMessage(s, i.ChannelID, "By hitting the play button, you will send the code off to be executed " +
		"and tested. The result will then be returned to you.\n" +
		"One thing to note is that the code is complied as one continuous line so" +
		"take this into account when running your code"); err != nil {
		log.Println(err)
		return
	}
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
