package command

import (
	"context"
	"strings"

	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const FoundBountyName = "Found Bounty"
const FoundBountyPostback = "found"

type FoundBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewFoundBountyCommand(repo bounty.Repository) Command {
	return &FoundBountyCommand{bountyRepo: repo}
}

func (c *FoundBountyCommand) Name() string {
	return FoundBountyName
}

func (c *FoundBountyCommand) Execute(ctx context.Context, event *linebot.Event) error {
	id := strings.Split(event.Postback.Data, "=")[1]
	err := c.bountyRepo.UpdateStatusBountyByID(ctx, id, bounty.Founded)
	if err != nil {
		return err
	}

	return bot.ReplyMessageWithText(event.ReplyToken, "Update Status: Founded")
}
