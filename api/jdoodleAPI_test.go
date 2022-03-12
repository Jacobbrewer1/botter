package api

import (
	"encoding/json"
	"github.com/Jacobbrewer1/botter/config"
	"io/ioutil"
	"log"
	"testing"
)

func TestRunCode(t *testing.T) {

	log.Println("config.json detected")
	file, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(string(file))

	err = json.Unmarshal(file, &config.JsonConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	config.ReadConfig()

	c := `
package main
import "fmt"

func main() {
fmt.Println("Hello world")
}`

	g := "go"
	v := "4"
	e := ExecuteInput{
		ClientId:     config.ApiSecrets.JdoodleApiClientId,
		ClientSecret: config.ApiSecrets.JdoodleApiClientSecret,
		Script:       &c,
		Language:     &g,
		VersionIndex: &v,
	}

	got, err := RunCode(e)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(got)
}
