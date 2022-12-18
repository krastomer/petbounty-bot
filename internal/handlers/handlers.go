package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/handlers/response"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Handlers interface {
	RegisterRouter(srv *gin.Engine)
}

type handlers struct {
	bountyRepo bounty.Repository
}

func InitializeHandlers(bountyRepo bounty.Repository) Handlers {
	return &handlers{
		bountyRepo: bountyRepo,
	}
}

func (h *handlers) RegisterRouter(srv *gin.Engine) {
	apiV1Group := srv.Group("/api/v1")

	apiV1Group.POST("/callback", h.callback)
}

func (h *handlers) callback(ctx *gin.Context) {
	events, err := bot.BotInstance.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	for _, event := range events {
		bounties, err := h.bountyRepo.GetBounty(ctx, "")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(bounties) == 0 {
			continue
		}
		// bot.BotInstance.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bounties[0].Name)).Do()
		content := response.MapBountyToFlexMessage(*bounties[0])
		bot.BotInstance.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("test", content)).Do()

	}
}
