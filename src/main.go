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

	// works
	worksRepository := repository.NewWorksRepositoryImpl(db)
	worksService := service.NewWorksServiceImpl(worksRepository, validate)
	worksController := controller.NewWorksController(worksService)

	//shifts
	shiftsRepository := repository.NewShiftsRepositoryImpl(db)
	shiftsService := service.NewShiftsServiceImpl(shiftsRepository, validate)
	shiftsController := controller.NewShiftController(shiftsService)

	// Router
	routes := router.NewRouter(tagsController, worksController, shiftsController)

	// サーバー立ち上げ
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	//サーバーの状態に関するエラーハンドリング
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
