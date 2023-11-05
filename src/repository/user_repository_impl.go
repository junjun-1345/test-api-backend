package repository

import (
	"sample/helper"
	"sample/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (u *UserRepositoryImpl) Save(user model.User) {
	// DBに保存できない場合はエラーが出て終了
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}
