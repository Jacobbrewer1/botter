package bot

import (
	"github.com/Jacobbrewer1/botter/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func TestingSetupMemberRoles(level int) []string {
	switch level {
	case moderator.level:
		var r = []string{everyone.id, member.id, moderator.id}
		return r
	case squad.level:
		var r = []string{everyone.id, member.id, squad.id}
		return r
	case minecraft.level:
		var r = []string{everyone.id, member.id, minecraft.id}
		return r
	case member.level:
		var r = []string{everyone.id, member.id}
		return r
	default:
		var r = []string{"811180171830886410"}
		return r
	}
}

func TestingSetupBadWords() error {
	if exists := TestingFindFile("../config/badWordList.txt"); exists {
		file, err := ioutil.ReadFile("../config/badWordList.txt")
		if err != nil {
			log.Println(err.Error())
			return err
		}
		log.Println(file)
		config.BannedWordsArray = strings.Split(string(file), ",")

		for arrayPos, word := range config.BannedWordsArray {
			newWord := word
			if strings.Contains(word, "\r") {
				newWord = strings.TrimLeft(newWord, "\r")
			}
			if strings.Contains(word, "\n") {
				newWord = strings.TrimLeft(newWord, "\n")
			}
			config.BannedWordsArray[arrayPos] = newWord
		}
	}
	return nil
}

func TestingFindFile(path string) bool {
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(abs)

	file, err := os.Open(abs)
	if err != nil {
		return false
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return true
}
