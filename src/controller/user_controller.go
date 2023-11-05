package controller

import (
	"fmt"
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
	fmt.Printf("ここまで来たよ5")
	// この下を繰り返す
	// リクエストされたデータをDBに格納できるような型を作成
	createUserRequest := request.CreateUserRequest{}
	// 引数をJSONに変更する
	err := ctx.ShouldBindJSON(&createUserRequest)
	// できない場合はエラーが出て動作が終了
	helper.ErrorPanic(err)
	fmt.Printf("ここまで来たよ4")

	// DBに保存する
	controller.userService.Create(createUserRequest)
	fmt.Printf("ここまで来たよ3")
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}
