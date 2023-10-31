package repositories

import (
	"sample/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OptionRepository struct {
	DB *gorm.DB
}

func NewOptionRepository(db *gorm.DB) *OptionRepository {
	return &OptionRepository{DB: db}
}

func (OptionRepository *OptionRepository) Save(option *models.Option) error {
	return OptionRepository.DB.Create(option).Error
}

func (OptionRepository *OptionRepository) FindByProductId(productId uuid.UUID) (options []*models.Option, err error) {
	if err = OptionRepository.DB.Where("product_id = ?", productId).Find(&options).Error; err != nil {
		return
	}
	return
}