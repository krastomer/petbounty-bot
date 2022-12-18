package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/services/command"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Handlers interface {
	RegisterRouter(srv *gin.Engine)
}

type handlers struct {
	commandService command.Service
}

func InitializeHandlers(commandSvc command.Service) Handlers {
	return &handlers{
		commandService: commandSvc,
	}
}

func (h *handlers) RegisterRouter(srv *gin.Engine) {
	apiV1Group := srv.Group("/api/v1")

	apiV1Group.POST("/callback", h.callback)
}

func (h *handlers) callback(ctx *gin.Context) {
	events, err := bot.GetInstance().ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	for _, event := range events {
		h.commandService.Execute(ctx, event)
	}
}
