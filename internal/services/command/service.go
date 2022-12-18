package command

import (
	"context"

	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Service interface {
	Execute(ctx context.Context, event *linebot.Event)
}

type service struct {
	commands map[string]Command
	// auditLog
}

func InitializeService(bountyRepo bounty.Repository) Service {
	commands := make(map[string]Command)

	bountyCmd := NewGetBountyCommand(bountyRepo)
	commands[bountyCmd.Name()] = bountyCmd

	return &service{commands: commands}
}

func (s *service) Execute(ctx context.Context, event *linebot.Event) {
	if event.Type != linebot.EventTypeMessage {
		return
	}

	text, ok := event.Message.(*linebot.TextMessage)
	if !ok {
		return
	}

	command := s.commands[text.Text]
	if command == nil {
		quickReplyItems := linebot.QuickReplyItems{
			Items: []*linebot.QuickReplyButton{
				linebot.NewQuickReplyButton("", &linebot.MessageAction{
					Label: "Get Bounty",
					Text:  "Get Bounty",
				}),
			},
		}
		bot.BotInstance.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("I don't understand what you say. Please type again").WithQuickReplies(&quickReplyItems)).Do()
		return
	}

	command.Execute(ctx, event)
}
