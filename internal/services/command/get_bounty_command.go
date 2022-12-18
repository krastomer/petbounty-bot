package command

import (
	"context"

	"github.com/krastomer/petbounty-bot/internal/bot"
	flexmessage "github.com/krastomer/petbounty-bot/internal/bot/flex_message"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type GetBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewGetBountyCommand(repo bounty.Repository) Command {
	return &GetBountyCommand{bountyRepo: repo}
}

func (c *GetBountyCommand) Name() string {
	return "Get Bounty"
}

func (c *GetBountyCommand) Execute(ctx context.Context, event *linebot.Event) {
	bounties, err := c.bountyRepo.GetBounty(ctx)
	if err != nil {
		return
	}
	if len(bounties) == 0 {
		return
	}

	contents := make([]*linebot.BubbleContainer, len(bounties))
	for index, bounty := range bounties {
		contents[index] = flexmessage.MapBountyToBubbleContainer(*bounty, false)
	}

	carousel := linebot.NewFlexMessage("test", &linebot.CarouselContainer{Contents: contents})
	bot.BotInstance.ReplyMessage(event.ReplyToken, carousel).Do()
}
