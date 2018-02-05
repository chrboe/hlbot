package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	btoken, err := ioutil.ReadFile("telegram.token") // just pass the file name
    if err != nil {
        log.Fatal("Error while reading token file: ", err)
    }


    token := string(btoken)
    token = strings.Map(func(r rune) rune {
  		if unicode.IsSpace(r) {
    		return -1
  		}
  		return r
	}, token)

    log.Print(token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Error while creating bot: ", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://bots.cnot.es/hlbot"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:8890", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
