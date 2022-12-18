package di

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/handlers"
	"github.com/krastomer/petbounty-bot/internal/repositories"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
)

func InitializeContainer() *Container {
	bot.InitializeBot()

	mongodb := repositories.InitializeRepositories()
	database := mongodb.GetDatabase()
	bountyRepo := bounty.ProvideRepository(database)

	srv := gin.Default()

	handlers := handlers.InitializeHandlers(bountyRepo)
	handlers.RegisterRouter(srv)

	return &Container{
		srv: srv,
	}
}
