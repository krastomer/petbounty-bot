package main

import (
	"github.com/joho/godotenv"
	"github.com/krastomer/petbounty-bot/internal/di"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	contianer := di.InitializeContainer()
	contianer.Run()
}
