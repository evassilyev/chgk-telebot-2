package messages

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Text message handlers

func (mh *Handler) undefined(msg *tgbotapi.Message) error {
	fmt.Printf("Undefined message from %s with text %s\n", msg.From, msg.Text)
	_, err := mh.bot.API.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Неизвестная команда/ответ:\n%s", msg.Text)))
	if err != nil {
		return err
	}
	return printTg(msg)
}

func (mh *Handler) newGroup(msg *tgbotapi.Message) error {
	return mh.db.AddGroup(msg.Chat.ID)
}

func printTg(msg interface{}) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println("======================================")
	fmt.Println(string(b))
	fmt.Println("======================================")
	return nil
}
