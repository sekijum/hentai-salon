package service

import (
	"context"
	"server/infrastructure/datasource"
	resource "server/presentation/resource"
)

type TagApplicationService struct {
	tagDatasource *datasource.TagDatasource
}

func NewTagApplicationService(tagDatasource *datasource.TagDatasource) *TagApplicationService {
	return &TagApplicationService{tagDatasource: tagDatasource}
}

func (svc *TagApplicationService) FindAllName(ctx context.Context) ([]string, error) {
	tags, err := svc.tagDatasource.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	tagNames := resource.GetTagNames(tags)
	return tagNames, nil
}
