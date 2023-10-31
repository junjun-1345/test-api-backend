package models

import (
	"github.com/google/uuid"
)

type Option struct {
	ID uuid.UUID `gorm:"primaryKey"`
	ProductID uuid.UUID `gorm:"size:36"`
	Name string
	Price int
	Product Product
}

func (option *Option) Create(ProductID uuid.UUID, name string, price int) {
	option.ID = uuid.New()
	option.ProductID = ProductID
	option.Name = name
	option.Price = price
}