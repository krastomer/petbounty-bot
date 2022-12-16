package main

import (
	"github.com/krastomer/petbounty-bot/internal/di"
)

func main() {
	contianer := di.InitializeContainer()
	contianer.Run()
}
