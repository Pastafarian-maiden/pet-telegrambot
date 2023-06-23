package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	user    *User
	address *Address
}

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Address struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	UserId string `json:"userId,omitempty"`
}
