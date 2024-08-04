package service_admin

import (
	"context"
	datasource_admin "server/infrastructure/datasource/admin"
	request_admin "server/presentation/request/admin"
	response_admin "server/presentation/response/admin"
)

type TagApplicationService struct {
	tagDatasource *datasource_admin.TagDatasource
}

func NewTagApplicationService(tagDatasource *datasource_admin.TagDatasource) *TagApplicationService {
	return &TagApplicationService{tagDatasource: tagDatasource}
}

type TagApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request_admin.TagFindAllRequest
}

func (svc *TagApplicationService) FindAll(params TagApplicationServiceFindAllParams) (*response_admin.Collection[*response_admin.TagResponse], error) {
	tagList, err := svc.tagDatasource.FindAll(datasource_admin.TagDatasourceFindAllParams{
		Ctx:     params.Ctx,
		Limit:   params.Qs.Limit,
		Offset:  params.Qs.Offset,
		Sort:    params.Qs.Sort,
		Order:   params.Qs.Order,
		Keyword: params.Qs.Keyword})
	if err != nil {
		return nil, err
	}

	threadCount, err := svc.tagDatasource.GetTagCount(datasource_admin.TagDatasourceGetTagCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	var tagResponseList []*response_admin.TagResponse
	for _, tag_i := range tagList {
		threadResponse := response_admin.NewTagResponse(response_admin.NewTagResponseParams{
			Tag: tag_i,
		})
		tagResponseList = append(tagResponseList, threadResponse)
	}

	dto := response_admin.NewCollection(response_admin.NewCollectionParams[*response_admin.TagResponse]{
		Data:       tagResponseList,
		TotalCount: threadCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type TagApplicationServiceDeleteParams struct {
	Ctx   context.Context
	TagID int
}

func (svc *TagApplicationService) Delete(params TagApplicationServiceDeleteParams) error {
	err := svc.tagDatasource.Delete(datasource_admin.TagDatasourceDeleteParams{
		Ctx:   params.Ctx,
		TagID: params.TagID,
	})
	if err != nil {
		return err
	}

	return nil
}
