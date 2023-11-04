package service

import (
	"sample/data/response"
	"sample/helper"
	"sample/model"
	"sample/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type WorksServiceImpl struct {
	WorksRepository repository.WorksRepository
	Validate        *validator.Validate
}

func NewWorksServiceImpl(worksRepository repository.WorksRepository,
	validate *validator.Validate) WorksService {
	return &WorksServiceImpl{
		WorksRepository: worksRepository,
		Validate:        validate,
	}
}

// // Create implements WorksService.
// func (l *WorksServiceImpl) Create(userId string) {
// 	// IDが付与されるように変更
// 	lineModel := model.Works{
// 		UserID:  userId,
// 		Content: "attendance",
// 	}
// 	// DBに保存
// 	l.WorksRepository.Save(lineModel)
// }

// 全表示
func (l *WorksServiceImpl) FindAll() []response.WorksReponce {
	result := l.WorksRepository.FindAll()

	var works []response.WorksReponce
	for _, value := range result {
		work := response.WorksReponce{
			ID:          value.ID,
			UserID:      value.UserID,
			Date:        value.Date,
			ClockIn:     value.ClockIn,
			ClockOut:    value.ClockOut,
			WorkingTime: value.WorkingTime,
		}
		works = append(works, work)
	}

	return works
}

// 出勤
func (l *WorksServiceImpl) ClockIn(userId string) {
	//　シフトがあるか確認
	now := nowLocTime()
	dateInt := nowLocDateInt()

	lineModel := model.Works{
		UserID:      userId,
		Date:        dateInt,
		ClockIn:     now,
		ClockOut:    now,
		WorkingTime: 0,
	}
	// DBに保存
	l.WorksRepository.Save(lineModel)
}

// 退勤
func (l *WorksServiceImpl) ClockOut(userId string) {
	now := nowLocTime()

	workData, err := l.WorksRepository.FindByUserIdAndWorkingTime(userId, 0)
	helper.ErrorPanic(err)

	workData.ClockOut = now
	workData.WorkingTime = int64(now.Sub(workData.ClockIn))

	// DBに保存
	l.WorksRepository.Update(workData)
}

// 日時計算
func nowLocTime() time.Time {
	// 日本のロケールを設定します。
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load location")
	}

	now := time.Now().In(loc)

	return now
}

func nowLocDateInt() int {
	// 日本のロケールを設定します。
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load location")
	}

	now := time.Now().In(loc)

	year, month, day := now.Date()

	monthInt := int(month)
	dayInt := day

	// int型の月と他の値を文字列でフォーマット
	dateInt := year*10000 + monthInt*100 + dayInt

	return dateInt
}
