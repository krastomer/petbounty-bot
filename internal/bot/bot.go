package bot

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var BotInstance *linebot.Client

func init() {
	token := os.Getenv("BOT_CHANNEL_TOKEN")
	secret := os.Getenv("BOT_CHANNEL_SECRET")
	client, err := linebot.New(secret, token)
	if err != nil {
		panic(err)
	}

	BotInstance = client
}
