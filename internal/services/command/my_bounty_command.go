package command

import (
	"context"

	"github.com/krastomer/petbounty-bot/internal/bot"
	flexmessage "github.com/krastomer/petbounty-bot/internal/bot/flex_message"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MyBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewMyBountyCommand(repo bounty.Repository) Command {
	return &MyBountyCommand{bountyRepo: repo}
}

func (c *MyBountyCommand) Name() string {
	return "My Bounty"
}

func (c *MyBountyCommand) Execute(ctx context.Context, event *linebot.Event) {
	bounties, err := c.bountyRepo.GetBountyByUserID(ctx, event.Source.UserID)
	if err != nil {
		return
	}
	if len(bounties) == 0 {
		return
	}

	contents := make([]*linebot.BubbleContainer, len(bounties))
	for index, bounty := range bounties {
		contents[index] = flexmessage.MapBountyToBubbleContainer(*bounty, true)
	}

	carousel := linebot.NewFlexMessage("test", &linebot.CarouselContainer{Contents: contents})
	bot.BotInstance.ReplyMessage(event.ReplyToken, carousel).Do()
}
