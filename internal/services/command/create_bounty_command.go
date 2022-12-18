package command

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBountyContext struct{}

type CreateBountyCommand struct {
	bountyRepo bounty.Repository
}

func NewCreateBountyCommand(repo bounty.Repository) Command {
	return &CreateBountyCommand{bountyRepo: repo}
}

func (c *CreateBountyCommand) Name() string {
	return "Create Bounty"
}

func (c *CreateBountyCommand) Execute(ctx context.Context, event *linebot.Event) {
	if ctx.Value(CreateBountyContext{}) == true {
		text, _ := event.Message.(*linebot.TextMessage)
		rawData := text.Text
		lines := strings.Split(rawData, "\n")

		if len(lines) != 5 {
			return
		}
		reward, err := strconv.ParseFloat(trimTitleAndSpace(lines[1]), 64)
		if err != nil {
			return
		}

		newBounty := bounty.Bounty{}
		newBounty.ID = primitive.NewObjectID()
		newBounty.UserID = event.Source.UserID
		newBounty.Name = trimTitleAndSpace(lines[0])
		newBounty.Reward = reward
		newBounty.Detail = trimTitleAndSpace(lines[2])
		newBounty.Address = trimTitleAndSpace(lines[3])
		newBounty.Telephone = trimTitleAndSpace(lines[4])
		newBounty.CreatedAt = time.Now()
		newBounty.Status = bounty.Missing
		err = c.bountyRepo.CreateBounty(ctx, newBounty)
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := bot.BotInstance.BroadcastMessage(linebot.NewTextMessage("แมวบักเจดหาย")).Do()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		return
	}

	bot.BotInstance.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage("Please Fill you're pet information\nCopy Next Message and fill"),
		linebot.NewTextMessage("Name: \nReward: \nDetail: \nAddress: \nTelephone: "),
	).Do()

}

func trimTitleAndSpace(in string) string {
	return strings.TrimSpace(strings.Split(in, ":")[1])
}
