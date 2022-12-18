package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/logger"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	ginzap "github.com/gin-contrib/zap"
)

type Handlers interface {
	RegisterRouter(srv *gin.Engine)
}

type handlers struct {
	// commandService
}

func InitializeHandlers() Handlers {
	return &handlers{}
}

func (h *handlers) RegisterRouter(srv *gin.Engine) {
	srv.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))
	srv.Use(ginzap.RecoveryWithZap(logger.Logger, true))

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
			quickReply := linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton("HEHE", nil),
			)
			bot.BotInstance.ReplyMessage(event.ReplyToken)
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				response := bot.BotInstance.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text).WithQuickReplies(quickReply))
				if _, err := response.Do(); err != nil {
					fmt.Println("hehe")
				}
			default:
				fmt.Println("hehe")
			}
		}
	}
}
