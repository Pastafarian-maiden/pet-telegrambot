package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) GetUsers(inputMessage update.Message) {
	var msg tgbotapi.MessageConfig

	request := "localhost:8080/api/names"

	if response, err := http.Get(request); err != nil {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "API is not responding")
	}

	defer response.Body.Close()

	//Reading answer
	contents, _ := ioutil.ReadAll(response.Body)

	//Unmarshal answer and set it to struct
	user := &User{}
	if err = json.Unmarshal([]byte(contents), user); err != nil {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "API is not responding")
	}

	//Check if struct is not empty
	if user.Name == "" {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "There is no data")
	}

	//Loop through Results struct and assigning data to s slice

	c.bot.Send(user)
}
func (c *Commander) GetAddresses(inputMessage update.Message) {

}
func (c *Commander) AddUser(inputMessage update.Message) {

}
func (c *Commander) AddAddresses(inputMessage update.Message) {

}
