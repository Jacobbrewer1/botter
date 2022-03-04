package main

import (
	"github.com/Jacobbrewer1/botter/bot"
	"github.com/Jacobbrewer1/botter/config"
	"log"
)

func init() {
	log.Println("Initializing logging")
	//log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("Logging initialized")
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
