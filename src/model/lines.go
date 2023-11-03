package model

import "gorm.io/gorm"

type Lines struct {
	gorm.Model
	UserID  string `gorm:"type:varchar(255)"`
	Content string `gorm:"type:varchar(255)"`
}
