package resource

import "server/domain/model"

type TagResource struct {
	Name string `json:"name"`
}

func NewTagResource(tag *model.Tag) *TagResource {
	return &TagResource{
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

func GetTagNames(tags []*model.Tag) []string {
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	return tagNames
}
