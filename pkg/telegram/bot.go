package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Tbot struct {
	API *tgbotapi.BotAPI
}

// NewBot initialisation of the Telegram bot
// token - is telegram bot token
// certificate - .pem file for telegram webhook
// webHookUrl - url for telegram webhook (domain + handler URL)
func NewBot(token, certificate, webHookUrl string) (*Tbot, error) {
	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	// Set webhook
	wh, err := tgbotapi.NewWebhookWithCert(webHookUrl, tgbotapi.FilePath(certificate))
	if err != nil {
		return nil, err
	}
	resp, err := b.Request(wh)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", resp)
	return &Tbot{
		API: b,
	}, nil
}
