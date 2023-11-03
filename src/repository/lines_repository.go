package repository

import "sample/model"

type LinesRepository interface {
	Save(lines model.Lines)
	FindAll() []model.Lines
}
