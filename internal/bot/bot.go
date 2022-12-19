package bot

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type bot struct {
	*linebot.Client
	quickReplyItems linebot.QuickReplyItems
}

var botInstance *bot

func InitializeBot() {
	token := os.Getenv("BOT_CHANNEL_TOKEN")
	secret := os.Getenv("BOT_CHANNEL_SECRET")
	client, err := linebot.New(secret, token)
	if err != nil {
		panic(err)
	}

	botInstance = &bot{client, linebot.QuickReplyItems{}}
}

func GetInstance() *bot {
	return botInstance
}

func SetQuickReplyItems(items linebot.QuickReplyItems) {
	botInstance.quickReplyItems = items
}

func GetQuickReplyItems() *linebot.QuickReplyItems {
	return &botInstance.quickReplyItems
}

func ReplyMessageWithText(replyToken string, text string) error {
	return replyMessageWithQuickReply(replyToken, linebot.NewTextMessage(text))
}

func ReplyMessageWithFlexMessage(replyToken string, message *linebot.FlexMessage) error {
	return replyMessageWithQuickReply(replyToken, message)
}

func replyMessageWithQuickReply(replyToken string, message ...linebot.SendingMessage) error {
	message[len(message)-1] = message[len(message)-1].WithQuickReplies(&botInstance.quickReplyItems)
	_, err := botInstance.ReplyMessage(replyToken, message...).Do()
	return err
}

func ReplyErrorMessage(replyToken string) error {
	return ReplyMessageWithText(replyToken, "I don't understand what you say. Please try again")
}
