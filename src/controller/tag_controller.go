package controller

import (
	"net/http"
	"sample/data/request"
	"sample/data/response"
	"sample/helper"
	"sample/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ルートを設定するファイル

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// データを作成する
func (controller *TagsController) Create(ctx *gin.Context) {
	// リクエストされたデータをDBに格納できるような型を作成
	createTagsRequest := request.CreateTagRequest{}
	// 引数をJSONに変更する
	err := ctx.ShouldBindJSON(&createTagsRequest)
	// できない場合はエラーが出て動作が終了
	helper.ErrorPanic(err)

	// DBに保存する
	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// データを更新する
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequset := request.UpdatesTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequset)
	helper.ErrorPanic(err)

	// データのIDをURLから取得
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	// データに収納
	updateTagsRequset.Id = id

	// 更新する
	controller.tagsService.Update(updateTagsRequset)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// データを削除する
func (controller *TagsController) Delete(ctx *gin.Context) {
	// データのIDをURLから取得
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// データを検索する
func (controller *TagsController) FindById(ctx *gin.Context) {
	// データのIDをURLから取得
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)

	// 返却するデータに取得したデータを格納
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// データを全表示する
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()

	// 返却するデータに取得したデータを格納
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tagResponse,
	}

	ctx.Header("Content-Type", "applicaton/json")
	ctx.JSON(http.StatusOK, webResponse)
}
