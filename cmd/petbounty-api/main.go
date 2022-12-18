package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/krastomer/petbounty-bot/internal/di"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	contianer := di.InitializeContainer()
	contianer.Run()
}
