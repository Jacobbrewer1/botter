package config

import (
	"encoding/json"
	"fmt"
	"github.com/Jacobbrewer1/botter/helper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	IgnoreVerification bool
	IgnoreBadWords     bool
	IgnoreInvites      bool

	JsonConfig *JsonConfigStruct
	ApiSecrets *ApiSecretsStruct
	override   *overrideStruct

	BannedWordsArray []string
)

func ReadConfig() error {
	if findFile("./config/override.json") {
		log.Println("Override detected - Reading file")

		file, err := ioutil.ReadFile("./config/override.json")
		if err != nil {
			return err
		}

		log.Println(string(file))

		err = json.Unmarshal(file, &override)
		if err != nil {
			return err
		}

		ApiSecrets = &ApiSecretsStruct{
			BotToken:       override.Secrets.BotToken,
			GithubApiToken: override.Secrets.GithubApiToken,
			GiphyApiToken:  override.Secrets.GiphyApiToken,
		}
		JsonConfig = new(JsonConfigStruct)

		if findFile("./config/config.json") {
			log.Println("config.json detected")
			file, err := ioutil.ReadFile("./config/config.json")
			if err != nil {
				return err
			}

			log.Println(string(file))

			err = json.Unmarshal(file, &JsonConfig)
			if err != nil {
				return err
			}
		}

		JsonConfig.BotPrefix = override.BotPrefix
		IgnoreVerification = *override.IgnoreVerification
		IgnoreInvites = *override.IgnoreInvites
		IgnoreBadWords = *override.IgnoreBadWords
	} else {
		log.Println("No override detected. Using production config")

		if findFile("./config/config.json") {
			log.Println("config.json detected")
			file, err := ioutil.ReadFile("./config/config.json")
			if err != nil {
				return err
			}

			log.Println(string(file))

			err = json.Unmarshal(file, &JsonConfig)
			if err != nil {
				return err
			}
		}

		a, err := getConfig()
		if err != nil {
			return err
		}
		ApiSecrets = &a
	}

	log.Printf("Bot prefix: %v", *JsonConfig.BotPrefix)

	if exists := findFile("./config/badWordList.txt"); exists {
		file, err := ioutil.ReadFile("./config/badWordList.txt")
		log.Println("Bad words file detected")
		if err != nil {
			return err
		}
		log.Println(string(file))
		BannedWordsArray = strings.Split(helper.FormatStrings(string(file)), ",")
	}
	return nil
}

func findFile(path string) bool {
	abs, err := filepath.Abs(path)
	if err != nil {
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

func getConfig() (ApiSecretsStruct, error) {
	rawJson, err := requestConfig()
	if err != nil {
		return ApiSecretsStruct{BotToken: nil}, err
	}
	log.Printf("Secrets json recieved:\n%v", string(rawJson))
	return decodeApiSecrets(rawJson)
}

func decodeApiSecrets(rawJson json.RawMessage) (ApiSecretsStruct, error) {
	var c ApiSecretsStruct
	err := json.Unmarshal(rawJson, &c)
	return c, err
}

func requestConfig() (json.RawMessage, error) {
	resp, err := http.Get(fmt.Sprintf("http://%v/botterconfig", *JsonConfig.ConfigIpAddress))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	log.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
