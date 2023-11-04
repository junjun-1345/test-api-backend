package service

import (
	"sample/data/request"
	"sample/helper"
	"sample/model"
	"sample/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type ShiftsServiceImpl struct {
	shiftsRepository repository.ShiftsRepository
	Validate         *validator.Validate
}

func NewShiftsServiceImpl(shiftsRepository repository.ShiftsRepository, validate *validator.Validate) ShiftsService {
	return &ShiftsServiceImpl{
		shiftsRepository: shiftsRepository,
		Validate:         validate,
	}
}

// Create implements ShiftsService.
func (s *ShiftsServiceImpl) Create(reqShifts request.CreateShiftsRequest) {
	// Structの型にできない場合はエラーを出して終了する
	err := s.Validate.Struct(reqShifts)
	helper.ErrorPanic(err)

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load location")
	}

	now := time.Now().In(loc)

	itemsToSave := make([]model.Shift, len(reqShifts.Shifts))

	// DBに保存
	for i, shiftData := range reqShifts.Shifts {
		itemsToSave[i] = model.Shift{
			UserID:   shiftData.UserID,
			Date:     shiftData.Date,
			ClockIn:  now,
			ClockOut: now,
			Decide:   false,
		}

	}

	// shifts := RequestToItems(reqShifts)

	s.shiftsRepository.Save(itemsToSave)
}

// func RequestToItems(requests []request.CreateShiftsRequest) []model.Shift {
// 	items := make([]model.Shift, len(requests))
// 	for i, req := range requests {
// 		items[i] = model.Shift{
// 			UserID: req.UserID,
// 			Date:   req.Date,
// 			Decide: false,
// 		}
// 	}
// 	return items
// }
