package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"unicode/utf8"
)


func main() {
	bot, err := tgbotapi.NewBotAPI(GetString(BOT_TOKEN))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			text := MessageBuilderInit().NormalText("Hello ").BoldText("UGS competitor ").
				NormalText("just send me the text,and i'll tell you the length of the text").
					ItalicText("\nCreated by Nikita Maliarenko =)").Text

			msg := 	tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ParseMode = "HTML"

			bot.Send(msg)
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Length is " + strconv.Itoa(utf8.RuneCountInString(update.Message.Text))))
		}
	}
}
