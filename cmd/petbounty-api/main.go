package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	secret := "awZJHEob/3reaim0dW+QoOB7c6HWx+lSVJ6ghZQ4Cre/p1oS6vTW4ghLE+g0vH6o979YWi/CXZ3Ibs09G3BQFWUKLKLqI9R3vFpjam3PpEXWPqrWfpXE1AhKZz1Gbb69plzywOiOAZL4IAmWChwbIgdB04t89/1O/w1cDnyilFU="
	token := "28b1e02cb20c19dd1e7376564338816b"
	bot, err := linebot.New(secret, token)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/callback", func(ctx *gin.Context) {
		events, err := bot.ParseRequest(ctx.Request)
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
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				default:
					fmt.Println("hehe")
				}
			}
		}
	})

	if err := router.Run(); err != nil {
		panic(err)
	}
}
