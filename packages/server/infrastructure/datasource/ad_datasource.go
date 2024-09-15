package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/ad"

	"entgo.io/ent/dialect/sql"
)

type AdDatasource struct {
	client *ent.Client
}

func NewAdDatasource(client *ent.Client) *AdDatasource {
	return &AdDatasource{client: client}
}

type AdDatasourceFindAllParams struct {
	Ctx context.Context
}

func (ds *AdDatasource) FindAll(params AdDatasourceFindAllParams) ([]*model.Ad, error) {
	entAdList, err := ds.client.
		Ad.
		Query().
		Where(ad.IsActiveEQ(1)).
		Order(sql.OrderByRand()).
		All(params.Ctx)

	if err != nil {
		return nil, err
	}

	var AdList []*model.Ad
	for _, entAd_i := range entAdList {
		AdList = append(AdList, model.NewAd(model.NewAdParams{EntAd: entAd_i}))
	}

	return AdList, nil
}
