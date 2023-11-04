package service

import (
	"sample/data/response"
	"sample/model"
	"sample/repository"

	"github.com/go-playground/validator/v10"
)

type LinesServiceImpl struct {
	LinesRepository repository.LinesRepository
	Validate        *validator.Validate
}

func NewLinesServiceImpl(linesRepository repository.LinesRepository,
	validate *validator.Validate) LinesService {
	return &LinesServiceImpl{
		LinesRepository: linesRepository,
		Validate:        validate,
	}
}

// Create implements LinesService.
func (l *LinesServiceImpl) Create(userId string) {
	// IDが付与されるように変更
	lineModel := model.Lines{
		UserID:  userId,
		Content: "attendance",
	}
	// DBに保存
	l.LinesRepository.Save(lineModel)
}

// FindAll implements LinesService.
func (l *LinesServiceImpl) FindAll() []response.LinesReponce {
	result := l.LinesRepository.FindAll()

	var lines []response.LinesReponce
	for _, value := range result {
		line := response.LinesReponce{
			Id:        int(value.ID),
			CreatedAt: value.CreatedAt.String(),
			UserId:    value.UserID,
			Content:   value.Content,
		}
		lines = append(lines, line)
	}

	return lines
}
