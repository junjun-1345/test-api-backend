package main

import (
	"sample/repositories"
	"sample/service"
	"sample/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age uint
}

func main() {
	db := utils.NewDBConnection()
	optionRepository := repositories.NewOptionRepository(db)
	productRepository := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepository, optionRepository)
	
	router:= gin.Default()
	router.Use(cors.Default())

	router.POST("/products", productService.RegisterProduct)
	router.GET("products/:id", productService.GetProductAndOptionsById)
	router.Run()
}