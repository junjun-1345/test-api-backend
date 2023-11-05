package controller

import (
	"net/http"
	"sample/data/request"
	"sample/data/response"
	"sample/helper"
	"sample/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// データを作成する
func (controller *UserController) Create(ctx *gin.Context) {
	// この下を繰り返す
	// リクエストされたデータをDBに格納できるような型を作成
	createUserRequest := request.CreateUserRequest{}
	// 引数をJSONに変更する
	err := ctx.ShouldBindJSON(&createUserRequest)
	// できない場合はエラーが出て動作が終了
	helper.ErrorPanic(err)

	// DBに保存する
	controller.userService.Create(createUserRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}
