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

type TagApplicationServiceFindNameListParams struct {
	Ctx context.Context
}

func (svc *TagApplicationService) FindNameList(params TagApplicationServiceFindNameListParams) ([]string, error) {
	tagList, err := svc.tagDatasource.FindAll(datasource.TagDatasourceFindAllParams{Ctx: params.Ctx})
	if err != nil {
		return nil, err
	}

	var tagNameList []string
	for _, tag_i := range tagList {
		tagNameList = append(tagNameList, tag_i.EntTag.Name)
	}
	return tagNameList, nil
}
