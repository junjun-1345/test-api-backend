package utils

import (
	"sample/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() *gorm.DB {
	dsn := "root:password@tcp(mysql_db)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		panic("マイグレーションに失敗しました")
	}

	return db
}