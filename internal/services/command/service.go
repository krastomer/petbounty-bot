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
	commands           map[string]Command
	errQuickReplyItems linebot.QuickReplyItems
	previousState      map[string]string
}

func InitializeService(bountyRepo bounty.Repository) Service {
	commands := make(map[string]Command)

	getBountyCmd := NewGetBountyCommand(bountyRepo)
	createBountyCmd := NewCreateBountyCommand(bountyRepo)
	myBountyCmd := NewMyBountyCommand(bountyRepo)

	commands[getBountyCmd.Name()] = getBountyCmd
	commands[createBountyCmd.Name()] = createBountyCmd
	commands[myBountyCmd.Name()] = myBountyCmd

	quickReplyButtons := make([]*linebot.QuickReplyButton, 0, len(commands))
	for command := range commands {
		quickReplyButtons = append(quickReplyButtons, linebot.NewQuickReplyButton("", linebot.NewMessageAction(command, command)))
	}

	errQuickReplyItems := *linebot.NewQuickReplyItems(quickReplyButtons...)

	return &service{commands: commands, errQuickReplyItems: errQuickReplyItems, previousState: make(map[string]string)}
}

func (s *service) Execute(ctx context.Context, event *linebot.Event) {
	if event.Type != linebot.EventTypeMessage {
		return
	}

	text, ok := event.Message.(*linebot.TextMessage)
	if !ok {
		return
	}

	if val, ok := s.previousState[event.Source.UserID]; ok {
		if val == "Create Bounty" {
			command := s.commands["Create Bounty"]
			ctx = context.WithValue(ctx, CreateBountyContext{}, true)
			command.Execute(ctx, event)
			s.previousState[event.Source.UserID] = "Created Bounty"
			return
		}
	}

	command := s.commands[text.Text]
	if command == nil {
		bot.BotInstance.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("I don't understand what you say. Please type again").WithQuickReplies(&s.errQuickReplyItems)).Do()
		return
	}

	command.Execute(ctx, event)
	s.previousState[event.Source.UserID] = text.Text
}
