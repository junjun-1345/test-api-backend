package repositories

import (
	"sample/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (ProductRepository *ProductRepository) Save(product *models.Product) error {
	return ProductRepository.DB.Create(product).Error
}

func (ProductRepository *ProductRepository) FindById(id uuid.UUID) (product *models.Product, err error) {
	if err = ProductRepository.DB.First(&product, id).Error; err != nil {
		return
	}
	return
}