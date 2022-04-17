package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Сайт", "https://postupi.s-vfu.ru/"),
		tgbotapi.NewInlineKeyboardButtonData("Заявки", "Заявки"),
		tgbotapi.NewInlineKeyboardButtonData("Отмена", "Отмена"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5304070353:AAFaZbnSqmuER8t0Sylm6FGyKDWDnhaiKw8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "Помощь":
				msg.Text = "1)Официальный Сайт " +
					"2)Подать заявку " +
					"3)Отмена "
				msg.ReplyMarkup = numericKeyboard
			case "Привет":
				msg.Text = "Привет, я бот. Чтобы получить список напиши: Помощь"
			default:
				msg.Text = "Оператор не может ответить на данное сообщение"

			}
			switch update.Message.Command() {
			case "help":
				msg.Text = "Вот список команд /PersonalArea , /status , /Website и отвечает на сообщения: Привет, Помощь"
			case "PersonalArea":
				msg.Text = "https://priem.s-vfu.ru/lka2022/"
			case "status":
				msg.Text = "Оператор онлайн"
			case "Website":
				msg.Text = "https://postupi.s-vfu.ru/"
			}
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
