package main

import (
	"net/http"
	"sample/config"
	"sample/controller"
	"sample/helper"
	"sample/repository"
	"sample/router"
	"sample/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")
	// データベース設定
	db := config.NewDBConnection()
	validate := validator.New()

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)

	//Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	// サーバー立ち上げ
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	//サーバーの状態に関するエラーハンドリング
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
