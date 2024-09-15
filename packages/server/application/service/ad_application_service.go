package service

import (
	"context"
	"server/infrastructure/datasource"
)

type AdApplicationService struct {
	AdDatasource *datasource.AdDatasource
}

func NewAdApplicationService(AdDatasource *datasource.AdDatasource) *AdApplicationService {
	return &AdApplicationService{AdDatasource: AdDatasource}
}

type AdApplicationServiceFindAllParams struct {
	Ctx context.Context
}

func (svc *AdApplicationService) FindAll(params AdApplicationServiceFindAllParams) ([]string, error) {
	adList, err := svc.AdDatasource.FindAll(datasource.AdDatasourceFindAllParams{
		Ctx: params.Ctx,
	})
	if err != nil {
		return nil, err
	}

	var adContentList []string
	for _, ad_i := range adList {
		adContentList = append(adContentList, ad_i.EntAd.Content)
	}
	return adContentList, nil

}
