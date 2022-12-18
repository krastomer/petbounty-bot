package command

import (
	"context"

	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	flexmessage "github.com/krastomer/petbounty-bot/internal/bot/flex_message"
)

const GetBountyName = "Get Bounty"

type GetBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewGetBountyCommand(repo bounty.Repository) Command {
	return &GetBountyCommand{bountyRepo: repo}
}

func (c *GetBountyCommand) Name() string {
	return GetBountyName
}

func (c *GetBountyCommand) Execute(ctx context.Context, event *linebot.Event) error {
	bounties, err := c.bountyRepo.GetBounties(ctx)
	if err != nil {
		return err
	}

	if len(bounties) == 0 {
		return bot.ReplyMessageWithText(event.ReplyToken, "No pet missing. Maybe next time.")
	}

	contents := make([]*linebot.BubbleContainer, len(bounties))
	for index, bounty := range bounties {
		contents[index] = flexmessage.MapBountyToBubbleContainer(*bounty, false)
	}

	carousel := linebot.NewFlexMessage(c.Name(), &linebot.CarouselContainer{Contents: contents})
	return bot.ReplyMessageWithFlexMessage(event.ReplyToken, carousel)
}
