package repository

import "sample/model"

type UserRepository interface {
	Save(user model.User)
}
