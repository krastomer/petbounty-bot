package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}

	// // Setup HTTP Server for receiving requests from LINE platform
	// http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
	// 	events, err := bot.ParseRequest(req)
	// 	if err != nil {
	// 		if err == linebot.ErrInvalidSignature {
	// 			w.WriteHeader(400)
	// 		} else {
	// 			w.WriteHeader(500)
	// 		}
	// 		return
	// 	}
	// 	for _, event := range events {
	// 		if event.Type == linebot.EventTypeMessage {
	// 			switch message := event.Message.(type) {
	// 			case *linebot.TextMessage:
	// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
	// 					log.Print(err)
	// 				}
	// 			case *linebot.StickerMessage:
	// 				replyMessage := fmt.Sprintf(
	// 					"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
	// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
	// 					log.Print(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// })
	// // This is just sample code.
	// // For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	// if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
	// 	log.Fatal(err)
	// }
}
