package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	switch update.Message.Command() {
	case "getUsers":
		c.GetUsers(update.Message)
	case "getAddress":
		c.GetAddresses(update.Message)
	case "addUser":
		c.AddUser(update.Message)
	case "addAddress":
		c.AddAddresses(update.Message)
	default:
		c.Default(update.Message)
	}
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot: bot,
	}
}
