package di

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/handlers"
	"github.com/krastomer/petbounty-bot/internal/repositories"
)

func InitializeContainer() *Container {
	srv := gin.New()

	handlers := handlers.InitializeHandlers()
	handlers.RegisterRouter(srv)

	mongodb := repositories.InitializeRepositories()
	_ = mongodb

	return &Container{
		srv: srv,
	}
}
