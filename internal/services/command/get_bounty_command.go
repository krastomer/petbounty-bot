package command

import (
	"context"

	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/handlers/response"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const name = "Get Bounty"

type GetBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewGetBountyCommand(repo bounty.Repository) Command {
	return &GetBountyCommand{bountyRepo: repo}
}

func (c *GetBountyCommand) Name() string {
	return name
}

func (c *GetBountyCommand) Execute(ctx context.Context, event *linebot.Event) {
	bounties, err := c.bountyRepo.GetBounty(ctx, "")
	if err != nil {
		return
	}
	if len(bounties) == 0 {
		return
	}

	content := response.MapBountyToFlexMessage(*bounties[0])
	bot.BotInstance.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("test", content)).Do()
}
