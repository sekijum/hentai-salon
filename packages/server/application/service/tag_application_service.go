package service

import (
	"context"
	"server/infrastructure/datasource"
)

type TagApplicationService struct {
	tagDatasource *datasource.TagDatasource
}

func NewTagApplicationService(tagDatasource *datasource.TagDatasource) *TagApplicationService {
	return &TagApplicationService{tagDatasource: tagDatasource}
}

type TagApplicationServiceFindAllNameParams struct {
	Ctx context.Context
}

func (svc *TagApplicationService) FindAllName(params TagApplicationServiceFindAllNameParams) ([]string, error) {
	tags, err := svc.tagDatasource.FindAll(datasource.TagDatasourceFindAllParams{
		Ctx: params.Ctx,
	})
	if err != nil {
		return nil, err
	}

	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name())
	}
	return tagNames, nil
}
