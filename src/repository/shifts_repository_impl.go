package repository

import (
	"sample/helper"
	"sample/model"

	"gorm.io/gorm"
)

type ShiftsRepositoryImpl struct {
	Db *gorm.DB
}

func NewShiftsRepositoryImpl(Db *gorm.DB) ShiftsRepository {
	return &ShiftsRepositoryImpl{Db: Db}
}

// 保存
func (s *ShiftsRepositoryImpl) Save(shifts model.Shifts) {
	result := s.Db.Create(&shifts)
	helper.ErrorPanic(result.Error)
}
