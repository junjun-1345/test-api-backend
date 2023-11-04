package repository

import (
	"errors"
	"sample/helper"
	"sample/model"

	"gorm.io/gorm"
)

type WorksRepositoryImpl struct {
	Db *gorm.DB
}

func NewWorksRepositoryImpl(Db *gorm.DB) WorksRepository {
	return &WorksRepositoryImpl{Db: Db}
}

func (l *WorksRepositoryImpl) Save(works model.Works) {
	// DBに保存できない場合はエラーが出て終了
	result := l.Db.Create(&works)
	helper.ErrorPanic(result.Error)
}

func (l *WorksRepositoryImpl) FindAll() []model.Works {
	var works []model.Works
	result := l.Db.Find(&works)
	helper.ErrorPanic(result.Error)
	return works
}

// 退勤検索
func (l *WorksRepositoryImpl) FindByUserIdAndWorkingTime(userId string, workingTime int64) (works model.Works, err error) {
	var work model.Works
	result := l.Db.Where("working_time = ?", 0).Find(&work)
	if result != nil {
		return work, nil
	} else {
		return work, errors.New("tag is not found")
	}
}

// アップデート
func (l *WorksRepositoryImpl) Update(works model.Works) {
	result := l.Db.Model(&works).Updates(works)
	helper.ErrorPanic(result.Error)
}
