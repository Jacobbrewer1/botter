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
		"java":   java,
		"c":      c,
		"c99":    c99,
		"cpp":    cpp,
		"cpp14":  cpp14,
		"cpp17":  cpp17,
		"php":    php,
		"perl":   perl,
	}

	golang = langStruct{
		languageCode: "go",
		index:        "4",
	}

	java = langStruct{
		languageCode: "java",
		index:        "4",
	}

	c = langStruct{
		languageCode: "c",
		index:        "5",
	}

	c99 = langStruct{
		languageCode: "c99",
		index:        "4",
	}

	cpp = langStruct{
		languageCode: "cpp",
		index:        "5",
	}

	cpp14 = langStruct{
		languageCode: "cpp14",
		index:        "1",
	}

	cpp17 = langStruct{
		languageCode: "cpp17",
		index:        "1",
	}

	php = langStruct{
		languageCode: "php",
		index:        "4",
	}

	perl = langStruct{
		languageCode: "perl",
		index:        "4",
	}
)

func containsCode(msg string) bool {
	_, l := getCode(msg)
	return strings.Count(msg, "```") == 2 && langMap[l].languageCode != ""
}

func getCode(input string) (string, string) {
	n := helper.RemoveNewLines(helper.RemoveTab(input))
	n = n[strings.Index(n, "```")+3:]
	n = n[:strings.Index(n, "```")]
	slice := strings.Split(n, " ")
	x := strings.Join(slice[1:], " ")
	return x[:len(x)-1], slice[0]
}

func runCompile(s *discordgo.Session, m *discordgo.Message) {
	n, language := getCode(m.Content)
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
	if _, err := sendMessage(s, i.ChannelID, "By hitting the play button, you will send the code off to be executed "+
		"and tested. The result will then be returned to you.\n"+
		"One thing to note is that the code is complied as one continuous line so"+
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
