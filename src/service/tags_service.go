package service

import (
	"sample/data/request"
	"sample/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsReponce
	FindAll() []response.TagsReponce
}
