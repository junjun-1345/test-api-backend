package repository

import (
	"sample/helper"
	"sample/model"

	"gorm.io/gorm"
)

type LinesRepositoryImpl struct {
	Db *gorm.DB
}

func NewLinesRepositoryImpl(Db *gorm.DB) LinesRepository {
	return &LinesRepositoryImpl{Db: Db}
}

// Save implements LinesRepository.
func (l *LinesRepositoryImpl) Save(lines model.Lines) {
	// DBに保存できない場合はエラーが出て終了
	result := l.Db.Create(&lines)
	helper.ErrorPanic(result.Error)
}

// FindAll implements LinesRepository.
func (l *LinesRepositoryImpl) FindAll() []model.Lines {
	var lines []model.Lines
	result := l.Db.Find(&lines)
	helper.ErrorPanic(result.Error)
	return lines
}
