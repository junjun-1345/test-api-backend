package main

import (
	"fmt"
	"net/http"
	"sample/config"
	"sample/controller"
	"sample/helper"
	"sample/repository"
	"sample/router"
	"sample/service"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")

	e := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if e != nil {
		fmt.Printf("読み込み出来ませんでした: %v", e)
	}

	// データベース設定
	db := config.NewDBConnection()
	validate := validator.New()

	// Tags
	tagsRepository := repository.NewTagsRepositoryImpl(db)
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	tagsController := controller.NewTagsController(tagsService)

	// Lines
	linesRepository := repository.NewLinesRepositoryImpl(db)
	linesService := service.NewLinesServiceImpl(linesRepository, validate)
	linesController := controller.NewLinesController(linesService)
	// Router
	routes := router.NewRouter(tagsController, linesController)

	// サーバー立ち上げ
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	//サーバーの状態に関するエラーハンドリング
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
