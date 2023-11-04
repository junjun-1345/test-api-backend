package controller

import (
	"net/http"
	"sample/data/request"
	"sample/data/response"
	"sample/helper"
	"sample/service"

	"github.com/gin-gonic/gin"
)

// ルートを設定するファイル

type ShiftController struct {
	shiftsService service.ShiftsService
}

func NewShiftController(service service.ShiftsService) *ShiftController {
	return &ShiftController{
		shiftsService: service,
	}
}

// データを作成する
func (controller *ShiftController) Create(ctx *gin.Context) {
	// この下を繰り返す
	// リクエストされたデータをDBに格納できるような型を作成
	createShiftsRequest := request.CreateShiftsRequest{}
	// 引数をJSONに変更する
	err := ctx.ShouldBindJSON(&createShiftsRequest)
	// できない場合はエラーが出て動作が終了
	helper.ErrorPanic(err)

	// DBに保存する
	controller.shiftsService.Create(createShiftsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}
