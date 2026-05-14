package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	TG_TOKEN = os.Getenv("TG_TOKEN")

	// numericKeyboard = tgbotapi.NewReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("1"),
	// 		tgbotapi.NewKeyboardButton("2"),
	// 		tgbotapi.NewKeyboardButton("3"),
	// 	),
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("4"),
	// 		tgbotapi.NewKeyboardButton("5"),
	// 		tgbotapi.NewKeyboardButton("6"),
	// 	),
	// )

	numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
			tgbotapi.NewInlineKeyboardButtonData("2", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		),
	)
)

func main() {
	if TG_TOKEN == "" {
		log.Fatalf("TG_TOKEN is not set")
		return
	}

	bot, err := tgbotapi.NewBotAPI(TG_TOKEN)
	if err != nil {
		log.Fatalln("tgbotapi init failed", err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// if update.Message == nil {
		// 	continue
		// }

		// echo example
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//
		// msg.ReplyToMessageID = update.Message.MessageID
		// echo example

		// command handling
		// if !update.Message.IsCommand() {
		// 	continue
		// }
		//
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		//
		// switch update.Message.Command() {
		// case "help":
		// 	msg.Text = "/sayhi or /status"
		// case "sayhi":
		// 	msg.Text = "Hiiii"
		// case "status":
		// 	msg.Text = "OK here"
		// default:
		// 	msg.Text = "use /help"
		// }
		// command handling

		// keyboard
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//
		// switch update.Message.Text {
		// case "open":
		// 	msg.ReplyMarkup = numericKeyboard
		// case "close":
		// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		// }
		// keyboard

		// if _, err := bot.Send(msg); err != nil {
		// 	log.Fatalln("send message failed:", err)
		// }

		// inline keyboard
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			}

			if _, err := bot.Send(msg); err != nil {
				log.Fatalln("send message failed:", err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				log.Fatalln("request callback err:", err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Fatalln("send message failed:", err)
			}
		}
		// inline keyboard
	}
}
