package config

import (
	"sample/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbName   = "test"
// )

// func DatabaseConnection() *gorm.DB {
// 	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
// 	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
// 	helper.ErrorPanic(err)

// 	return db
// }

func NewDBConnection() *gorm.DB {
	dsn := "root:password@tcp(mysql_db)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	err := db.AutoMigrate(&model.Tags{}, &model.Lines{})
	if err != nil {
		panic("マイグレーションに失敗しました")
	}

	return db
}
