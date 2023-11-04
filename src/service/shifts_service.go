package service

import "sample/data/request"

type ShiftsService interface {
	Create(shifts request.CreateShiftsRequest)
}
