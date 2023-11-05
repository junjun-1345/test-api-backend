package service

import (
	"sample/data/request"
	"sample/helper"
	"sample/model"
	"sample/repository"

	"github.com/go-playground/validator/v10"
)

// DBの操作をCRUD毎にまとめたファイル

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) {

	// Structの型にできない場合はエラーを出して終了する
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)

	// IDが付与されるように変更
	userModel := model.User{
		UserID:         user.UserId,
		Name:           user.Name,
		WorkInWeekDay:  user.WorkInWeekDay,
		WorkInWeekTime: user.WorkInWeekTime,
		Rank:           1,
		Vacation:       0,
		Admin:          false,
	}

	// DBに保存
	u.UserRepository.Save(userModel)
}
