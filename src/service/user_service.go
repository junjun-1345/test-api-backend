package service

import "sample/data/request"

type UserService interface {
	Create(user request.CreateUserRequest)
}
