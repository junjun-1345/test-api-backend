package service

import (
	"sample/data/request"
	"sample/data/response"
	"sample/helper"
	"sample/model"
	"sample/repository"

	"github.com/go-playground/validator/v10"
)

// DBの操作をCRUD毎にまとめたファイル

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

// データをDBに保存する
func (t *TagsServiceImpl) Create(tags request.CreateTagRequest) {
	// Structの型にできない場合はエラーを出して終了する
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)

	// IDが付与されるように変更
	tagModel := model.Tags{
		Name: tags.Name,
	}
	// DBに保存
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []response.TagsReponce {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsReponce
	for _, value := range result {
		tag := response.TagsReponce{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService.
func (t TagsServiceImpl) FindById(tagsId int) response.TagsReponce {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsReponce{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.UpdatesTagsRequest) {
	// 更新するデータを取得
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	// 更新するデータを上書き
	tagData.Name = tags.Name
	// DBを更新
	t.TagsRepository.Update(tagData)
}
