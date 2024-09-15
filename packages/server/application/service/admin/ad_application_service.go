package service_admin

import (
	"context"
	"server/domain/model"
	datasource_admin "server/infrastructure/datasource/admin"
	"server/infrastructure/ent"
	request_admin "server/presentation/request/admin"
	response_admin "server/presentation/response/admin"
)

type AdApplicationService struct {
	AdDatasource *datasource_admin.AdDatasource
}

func NewAdApplicationService(AdDatasource *datasource_admin.AdDatasource) *AdApplicationService {
	return &AdApplicationService{AdDatasource: AdDatasource}
}

type AdApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request_admin.AdFindAllRequest
}

func (svc *AdApplicationService) FindAll(params AdApplicationServiceFindAllParams) (*response_admin.Collection[*response_admin.AdResponse], error) {
	AdList, err := svc.AdDatasource.FindAll(datasource_admin.AdDatasourceFindAllParams{
		Ctx:     params.Ctx,
		Limit:   params.Qs.Limit,
		Offset:  params.Qs.Offset,
		Sort:    params.Qs.Sort,
		Order:   params.Qs.Order,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	AdCount, err := svc.AdDatasource.GetAdCount(datasource_admin.AdDatasourceGetAdCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	var AdResponseList []*response_admin.AdResponse
	for _, Ad_i := range AdList {
		AdResponse := response_admin.NewAdResponse(response_admin.NewAdResponseParams{Ad: Ad_i})
		AdResponseList = append(AdResponseList, AdResponse)
	}

	dto := response_admin.NewCollection(response_admin.NewCollectionParams[*response_admin.AdResponse]{
		Data:       AdResponseList,
		TotalCount: AdCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type AdApplicationServiceFindByIDParams struct {
	Ctx  context.Context
	AdID int
}

func (svc *AdApplicationService) FindByID(params AdApplicationServiceFindByIDParams) (*response_admin.AdResponse, error) {
	Ad, err := svc.AdDatasource.FindByID(datasource_admin.AdDatasourceFindByIDParams{
		Ctx:  params.Ctx,
		AdID: params.AdID,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewAdResponse(response_admin.NewAdResponseParams{Ad: Ad})

	return dto, nil
}

type AdApplicationServiceCreateParams struct {
	Ctx  context.Context
	Body request_admin.AdCreateRequest
}

func (svc *AdApplicationService) Create(params AdApplicationServiceCreateParams) error {
	Ad := model.NewAd(model.NewAdParams{
		EntAd: &ent.Ad{
			Content:  params.Body.Content,
			IsActive: params.Body.IsActive,
		},
	})

	err := svc.AdDatasource.Create(datasource_admin.AdDatasourceCreateParams{
		Ctx: params.Ctx,
		Ad:  *Ad,
	})
	if err != nil {
		return err
	}

	return nil
}

type AdApplicationServiceUpdateParams struct {
	Ctx  context.Context
	AdID int
	Body request_admin.AdUpdateRequest
}

func (svc *AdApplicationService) Update(params AdApplicationServiceUpdateParams) (*response_admin.AdResponse, error) {
	Ad := model.NewAd(model.NewAdParams{
		EntAd: &ent.Ad{
			ID:       params.AdID,
			Content:  params.Body.Content,
			IsActive: params.Body.IsActive,
		},
	})

	Ad, err := svc.AdDatasource.Update(datasource_admin.AdDatasourceUpdateParams{
		Ctx: params.Ctx,
		Ad:  *Ad,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewAdResponse(response_admin.NewAdResponseParams{Ad: Ad})

	return dto, nil
}

type AdApplicationServiceDeleteParams struct {
	Ctx  context.Context
	AdID int
}

func (svc *AdApplicationService) Delete(params AdApplicationServiceDeleteParams) error {
	err := svc.AdDatasource.Delete(datasource_admin.AdDatasourceDeleteParams{
		Ctx:  params.Ctx,
		AdId: params.AdID,
	})
	if err != nil {
		return err
	}

	return nil
}

type AdApplicationServiceUpdateIsActiveParams struct {
	Ctx  context.Context
	AdID int
	Body request_admin.AdUpdateIsActiveRequest
}

func (svc *AdApplicationService) UpdateIsActive(params AdApplicationServiceUpdateIsActiveParams) error {
	ad := model.NewAd(model.NewAdParams{
		EntAd: &ent.Ad{
			ID:       params.AdID,
			IsActive: params.Body.IsActive,
		},
	})

	_, err := svc.AdDatasource.UpdateIsActive(datasource_admin.AdDatasourceUpdateIsActiveParams{
		Ctx: params.Ctx,
		Ad:  ad,
	})
	if err != nil {
		return err
	}

	return nil
}
