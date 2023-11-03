package controller

import (
	"fmt"
	"net/http"
	"sample/config"
	"sample/data/response"
	"sample/service"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ルートを設定するファイル

type LinesController struct {
	linesService service.LinesService
}

func NewLinesController(service service.LinesService) *LinesController {
	return &LinesController{
		linesService: service,
	}
}

// データを作成する
func (controller *LinesController) Create(ctx *gin.Context) {
	// LineBotの設定
	bot := config.NewLineBot()
	// リクエスト処理
	events, berr := bot.ParseRequest(ctx.Request)
	if berr != nil {
		fmt.Println(berr.Error())
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			userId := event.Source.UserID
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "出勤" {
					// DBに保存する
					controller.linesService.Create(userId)
					webResponse := response.Response{
						Code:   http.StatusOK,
						Status: "ok",
						Data:   nil,
					}
					ctx.Header("Content-Type", "applicaton/json")
					ctx.JSON(http.StatusOK, webResponse)
					// 返信する
					_, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTextMessage("出勤しました!"),
					).Do()
					if err != nil {
						fmt.Println(err.Error())
					}
				} else {
					_, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTextMessage(getResMessage(message.Text, userId)),
					).Do()
					if err != nil {
						fmt.Println(err.Error())
					}
				}

			}
		}
	}
}

func getResMessage(message string, userId string) string {
	return userId + "さんが" + message + "と言いました。"
}

// データを全表示する
func (controller *LinesController) FindAll(ctx *gin.Context) {
	lineResponse := controller.linesService.FindAll()

	// 返却するデータに取得したデータを格納
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   lineResponse,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}
