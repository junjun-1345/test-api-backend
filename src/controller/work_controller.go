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

type WorksController struct {
	worksService service.WorksService
}

func NewWorksController(service service.WorksService) *WorksController {
	return &WorksController{
		worksService: service,
	}
}

// データを作成する
func (controller *WorksController) Create(ctx *gin.Context) {
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
					controller.worksService.ClockIn(userId)
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
				} else if message.Text == "退勤" {
					controller.worksService.ClockOut(userId)
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
						linebot.NewTextMessage("退勤しました!おつかれさまでした！"),
					).Do()
					if err != nil {
						fmt.Println(err.Error())
					}

				} else if message.Text == "はじめまして" {
					userIdUri := "https://jobkai.vercel.app/user/" + userId
					// 返信する
					container := &linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeVertical,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:   linebot.FlexComponentTypeText,
									Text:   "従業員登録",
									Weight: linebot.FlexTextWeightTypeBold,
									Size:   linebot.FlexTextSizeTypeXl,
									Align:  linebot.FlexComponentAlignTypeCenter,
								},
							},
						},
						Footer: &linebot.BoxComponent{
							Type:    linebot.FlexComponentTypeBox,
							Layout:  linebot.FlexBoxLayoutTypeVertical,
							Spacing: linebot.FlexComponentSpacingTypeSm,
							Contents: []linebot.FlexComponent{
								&linebot.ButtonComponent{
									Type:   linebot.FlexComponentTypeButton,
									Style:  linebot.FlexButtonStyleTypeLink,
									Height: linebot.FlexButtonHeightTypeSm,
									Action: linebot.NewURIAction("新規登録をする", userIdUri),
								},
								&linebot.BoxComponent{
									Type:     linebot.FlexComponentTypeBox,
									Layout:   linebot.FlexBoxLayoutTypeVertical,
									Contents: []linebot.FlexComponent{},
									Margin:   linebot.FlexComponentMarginTypeSm},
							},
						},
					}
					_, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewFlexMessage("alt text", container),
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
func (controller *WorksController) FindAll(ctx *gin.Context) {
	lineResponse := controller.worksService.FindAll()

	// 返却するデータに取得したデータを格納
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   lineResponse,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}
