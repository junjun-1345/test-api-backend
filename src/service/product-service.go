package service

import (
	"sample/models"
	"sample/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterProductRequestData struct {
	Name string `json:"name" binding:"required"`
	Src string `json:"src"`
	Price int `json:"price" binding:"required"`
}

type productService struct {
	productRepository *repositories.ProductRepository
	optionRepository *repositories.OptionRepository
}


func NewProductService(productRepository *repositories.ProductRepository,optionRepository *repositories.OptionRepository) *productService {
	return &productService{productRepository: productRepository, optionRepository: optionRepository}
}

func (productService *productService) RegisterProduct(c *gin.Context) {
	var requestData RegisterProductRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	product.Create(requestData.Name, requestData.Src, requestData.Price)
	productService.productRepository.Save(&product)

	c.JSON(201, gin.H{
		"message": "Successfully Created",
	})
}

func (productService *productService) GetProductAndOptionsById(c *gin.Context) {
	productId, parseErr := uuid.Parse(c.Param("id"))
	if parseErr != nil {
		c.JSON(400, gin.H{"error": parseErr.Error()})
		return
	}
	foundProduct, productErr := productService.productRepository.FindById(productId)
	if productErr != nil {
		c.JSON(500, gin.H{"error": productErr.Error()})
		return
	}

	foundOpitons, optionsErr := productService.optionRepository.FindByProductId(foundProduct.ID)
	if optionsErr != nil {
		c.JSON(500, gin.H{"error": optionsErr.Error()})
		return
	}

	var optionResponseData []map[string]interface{}
	for _, option := range foundOpitons {
		responseOption := map[string]interface{}{
			"id": option.ID,
			"productId": option.ProductID,
			"name": option.Name,
			"price": option.Price,
		}
		optionResponseData = append(optionResponseData, responseOption)
	}
	
	c.JSON(200, gin.H{
	"message": "OK",
	"product": map[string]interface{}{
		"id": foundProduct.ID,
		"name": foundProduct.Name,
		"src": foundProduct.Src,
		"price": foundProduct.Price,
		"options":optionResponseData,
		},
	})

}