package repository

import (
	"errors"
	"sample/data/request"
	"sample/helper"
	"sample/model"

	"gorm.io/gorm"
)

// DBの操作をする（SQL）ファイル

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// 削除機能
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	// 特定のデータを検索して削除
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// すべて表示
func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// 検索
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// 保存
func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	// DBに保存できない場合はエラーが出て終了
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// アップデート
func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
