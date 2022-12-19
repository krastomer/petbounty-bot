package command

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/krastomer/petbounty-bot/internal/bot"
	flexmessage "github.com/krastomer/petbounty-bot/internal/bot/flex_message"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CreateBountyName = "Create Bounty"

type CreateBountyContext struct{}

type CreateBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewCreateBountyCommand(repo bounty.Repository) Command {
	return &CreateBountyCommand{bountyRepo: repo}
}

func (c *CreateBountyCommand) Name() string {
	return CreateBountyName
}

func (c *CreateBountyCommand) Execute(ctx context.Context, event *linebot.Event) error {
	if ctx.Value(CreateBountyContext{}) == true {
		text, _ := event.Message.(*linebot.TextMessage)
		rawData := text.Text
		lines := strings.Split(rawData, "\n")

		if len(lines) != 5 {
			return ErrBadRequest
		}
		reward, err := strconv.ParseFloat(trimTitleAndSpace(lines[1]), 64)
		if err != nil {
			return ErrBadRequest
		}

		newBounty := bounty.Bounty{
			ID:        primitive.NewObjectID(),
			UserID:    event.Source.UserID,
			Name:      trimTitleAndSpace(lines[0]),
			Reward:    reward,
			Detail:    trimTitleAndSpace(lines[2]),
			Address:   trimTitleAndSpace(lines[3]),
			Telephone: trimTitleAndSpace(lines[4]),
			CreatedAt: time.Now(),
			Status:    bounty.Missing,
		}
		err = c.bountyRepo.CreateBounty(ctx, newBounty)
		if err != nil {
			return err
		}
		_, err = bot.GetInstance().BroadcastMessage(linebot.NewTextMessage("New Bounty Active"), linebot.NewFlexMessage(c.Name(), flexmessage.MapBountyToBubbleContainer(newBounty, false)).WithQuickReplies(bot.GetQuickReplyItems())).Do()
		return err
	}

	_, err := bot.GetInstance().ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage("Please Fill you're pet information\nCopy Next Message and fill"),
		linebot.NewTextMessage("Name: \nReward: \nDetail: \nAddress: \nTelephone: "),
	).Do()
	return err
}

func trimTitleAndSpace(in string) string {
	return strings.TrimSpace(strings.SplitN(in, ":", 2)[1])
}
