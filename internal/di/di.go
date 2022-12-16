package di

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/handlers"
)

func InitializeContainer() *Container {
	srv := gin.Default()

	handlers := handlers.InitializeHandlers()
	handlers.RegisterRouter(srv)

	return &Container{
		srv: srv,
	}
}
