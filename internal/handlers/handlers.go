package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Handlers interface {
	RegisterRouter(srv *gin.Engine)
}

type handlers struct{}

func InitializeHandlers() Handlers {
	return &handlers{}
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
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err := bot.BotInstance.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					fmt.Println("hehe")
				}
			default:
				fmt.Println("hehe")
			}
		}
	}
}
