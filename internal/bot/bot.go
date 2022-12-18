package bot

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type bot struct {
	*linebot.Client
}

var botInstance *bot

func InitializeBot() {
	token := os.Getenv("BOT_CHANNEL_TOKEN")
	secret := os.Getenv("BOT_CHANNEL_SECRET")
	client, err := linebot.New(secret, token)
	if err != nil {
		panic(err)
	}

	botInstance = &bot{client}
}

func GetInstance() *bot {
	return botInstance
}
