package resource

import "server/domain/model"

func GetTagNames(tags []*model.Tag) []string {
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name())
	}
	return tagNames
}
