package service

import (
	"context"
	"server/infrastructure/datasource"
	resource "server/presentation/resource/tag"
)

type TagApplicationService struct {
	tagDatasource *datasource.TagDatasource
}

func NewTagApplicationService(tagDatasource *datasource.TagDatasource) *TagApplicationService {
	return &TagApplicationService{tagDatasource: tagDatasource}
}

func (svc *TagApplicationService) FindAll(ctx context.Context) ([]*resource.TagResource, error) {
	tags, err := svc.tagDatasource.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return resource.NewTagResourceList(tags), nil
}
