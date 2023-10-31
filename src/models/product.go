package models

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID `gorm:"primaryKey"`
	Name string
	Src string
	Price int
}

func (product *Product) Create(name string, src string, price int) {
	product.ID = uuid.New()
	product.Name = name
	product.Src = src
	product.Price = price
}