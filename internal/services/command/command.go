package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Command interface {
	Name() string
	Execute(ctx context.Context, event *linebot.Event)
}
