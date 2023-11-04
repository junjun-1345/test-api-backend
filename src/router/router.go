package router

import (
	"net/http"
	"sample/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController, worksController *controller.WorksController) *gin.Engine {
	//ルーター関連・初期化
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	tagRouter := router.Group("/tags")
	tagRouter.GET("", tagsController.FindAll)
	tagRouter.GET("/:tagId", tagsController.FindById)
	tagRouter.POST("", tagsController.Create)
	tagRouter.PATCH("/:tagId", tagsController.Update)
	tagRouter.DELETE("/:tagId", tagsController.Delete)

	workRouter := router.Group("/works")
	workRouter.POST("/callback", worksController.Create)
	workRouter.GET("", worksController.FindAll)

	return router
}
