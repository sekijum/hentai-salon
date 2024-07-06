package resource

import "server/domain/model"

type TagResource struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewTagResource(tag *model.Tag) *TagResource {
	return &TagResource{
		Id:   tag.Id,
		Name: tag.Name,
	}
}

func NewTagResourceList(tags []*model.Tag) []*TagResource {
	var tagResources []*TagResource
	for _, tag := range tags {
		tagResources = append(tagResources, NewTagResource(tag))
	}
	return tagResources
}
