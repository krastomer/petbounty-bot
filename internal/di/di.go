package di

import (
	"github.com/gin-gonic/gin"
	"github.com/krastomer/petbounty-bot/internal/bot"
	"github.com/krastomer/petbounty-bot/internal/handlers"
	"github.com/krastomer/petbounty-bot/internal/repositories"
	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/krastomer/petbounty-bot/internal/services/command"
)

func InitializeContainer() *Container {
	bot.InitializeBot()

	mongodb := repositories.InitializeRepositories()
	database := mongodb.GetDatabase()
	bountyRepo := bounty.InitializeRepository(database)

	srv := gin.Default()

	commandSvc := command.InitializeService(bountyRepo)

	handlers := handlers.InitializeHandlers(commandSvc)
	handlers.RegisterRouter(srv)

	return &Container{
		srv: srv,
	}
}
