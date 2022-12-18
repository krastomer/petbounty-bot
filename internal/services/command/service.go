package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Service interface {
	Execute(ctx context.Context, event *linebot.Event)
}

type service struct {
	commands      map[string]Command
	postbacks     map[string]Command
	previousState map[string]string
}

type Command interface {
	Name() string
	Execute(ctx context.Context, event *linebot.Event) error
}

func InitializeService(bountyRepo bounty.Repository) Service {
	commands := make(map[string]Command)
	postbacks := make(map[string]Command)

	getBountyCmd := NewGetBountyCommand(bountyRepo)
	createBountyCmd := NewCreateBountyCommand(bountyRepo)
	myBountyCmd := NewMyBountyCommand(bountyRepo)
	foundBountyCmd := NewFoundBountyCommand(bountyRepo)

	commands[getBountyCmd.Name()] = getBountyCmd
	commands[createBountyCmd.Name()] = createBountyCmd
	commands[myBountyCmd.Name()] = myBountyCmd

	postbacks[FoundBountyPostback] = foundBountyCmd

	quickReplyButtons := make([]*linebot.QuickReplyButton, 0, len(commands))
	for command := range commands {
		quickReplyButtons = append(quickReplyButtons, linebot.NewQuickReplyButton("", linebot.NewMessageAction(command, command)))
	}

	bot.SetQuickReplyItems(*linebot.NewQuickReplyItems(quickReplyButtons...))

	return &service{commands: commands, postbacks: postbacks, previousState: make(map[string]string)}
}

func (s *service) Execute(ctx context.Context, event *linebot.Event) {
	switch event.Type {
	case linebot.EventTypeMessage:
		s.hitTextMessage(ctx, event)
	case linebot.EventTypePostback:
		s.hitPostback(ctx, event)
	}

}

func (s *service) hitTextMessage(ctx context.Context, event *linebot.Event) {
	text, ok := event.Message.(*linebot.TextMessage)
	if !ok {
		return
	}

	if val, ok := s.previousState[event.Source.UserID]; ok {
		if val == CreateBountyName {
			command := s.commands[CreateBountyName]
			ctx = context.WithValue(ctx, CreateBountyContext{}, true)
			command.Execute(ctx, event)
			s.previousState[event.Source.UserID] = "Created Bounty"
			return
		}
	}

	command := s.commands[text.Text]
	if command == nil {
		err := bot.ReplyMessageWithText(event.ReplyToken, "I don't understand what you say. Please try again")
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	err := command.Execute(ctx, event)
	if err != nil {
		fmt.Println(err)
	}

	s.previousState[event.Source.UserID] = command.Name()
}

func (s *service) hitPostback(ctx context.Context, event *linebot.Event) {
	postback := strings.Split(event.Postback.Data, "=")[0]
	command := s.postbacks[postback]

	if command == nil {
		err := bot.ReplyMessageWithText(event.ReplyToken, "I don't understand what you say. Please try again")
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	err := command.Execute(ctx, event)
	if err != nil {
		fmt.Println(err)
	}

	s.previousState[event.Source.UserID] = command.Name()
}
