package service

import (
	"sample/data/response"
)

type WorksService interface {
	// Create(userId string)
	FindAll() []response.WorksReponce
	ClockIn(userId string)
	ClockOut(userId string)
}
