package service

import "sample/data/response"

type LinesService interface {
	Create(userId string)
	FindAll() []response.LinesReponce
}
