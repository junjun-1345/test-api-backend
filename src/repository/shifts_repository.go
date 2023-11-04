package repository

import "sample/model"

type ShiftsRepository interface {
	Save(shifts model.Shifts)
}
