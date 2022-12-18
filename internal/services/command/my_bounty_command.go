package command

import (
	"context"

	"github.com/krastomer/petbounty-bot/internal/bot"
	flexmessage "github.com/krastomer/petbounty-bot/internal/bot/flex_message"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const MyBountyName = "My Bounty"

type MyBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewMyBountyCommand(repo bounty.Repository) Command {
	return &MyBountyCommand{bountyRepo: repo}
}

func (c *MyBountyCommand) Name() string {
	return MyBountyName
}

func (c *MyBountyCommand) Execute(ctx context.Context, event *linebot.Event) error {
	bounties, err := c.bountyRepo.GetBountyByUserID(ctx, event.Source.UserID)
	if err != nil {
		return err
	}

	if len(bounties) == 0 {
		err := bot.ReplyMessageWithText(event.ReplyToken, "You never register bounty.")
		return err
	}

	contents := make([]*linebot.BubbleContainer, len(bounties))
	for index, bounty := range bounties {
		contents[index] = flexmessage.MapBountyToBubbleContainer(*bounty, true)
	}

	carousel := linebot.NewFlexMessage(c.Name(), &linebot.CarouselContainer{Contents: contents})
	return bot.ReplyMessageWithFlexMessage(event.ReplyToken, carousel)
}
