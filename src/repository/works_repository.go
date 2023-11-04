package repository

import "sample/model"

type WorksRepository interface {
	Save(works model.Works)
	FindAll() []model.Works
	FindByUserIdAndWorkingTime(userId string, workingTime int64) (Works model.Works, err error)
	Update(works model.Works)
}
