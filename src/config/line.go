package config

import (
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func NewLineBot() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		panic("Line Botの作成に失敗しました")
	}

	return bot
}
