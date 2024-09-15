package datasource_admin

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/ad"
	"strconv"
	"time"
)

type AdDatasource struct {
	client *ent.Client
}

func NewAdDatasource(client *ent.Client) *AdDatasource {
	return &AdDatasource{client: client}
}

type AdDatasourceGetAdCountParams struct {
	Ctx     context.Context
	Keyword *string
}

func (ds *AdDatasource) GetAdCount(params AdDatasourceGetAdCountParams) (int, error) {
	q := ds.client.Ad.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "active:":
			if active, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(ad.IsActive(active))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(ad.IDEQ(id))
			}
		}
	}

	adCount, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return adCount, nil
}

type AdDatasourceFindByIDParams struct {
	Ctx  context.Context
	AdID int
}

func (ds *AdDatasource) FindByID(params AdDatasourceFindByIDParams) (*model.Ad, error) {
	entAd, err := ds.client.Ad.Get(params.Ctx, params.AdID)
	if err != nil {
		return nil, err
	}

	ad := model.NewAd(model.NewAdParams{EntAd: entAd})

	return ad, nil
}

type AdDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
}

func (ds *AdDatasource) FindAll(params AdDatasourceFindAllParams) ([]*model.Ad, error) {
	q := ds.client.Ad.Query()

	sort := ad.FieldID
	order := "desc"

	if params.Sort != nil && *params.Sort != "" {
		sort = *params.Sort
	}
	if params.Order != nil && *params.Order != "" {
		order = *params.Order
	}

	if order == "asc" {
		q = q.Order(ent.Asc(sort))
	} else {
		q = q.Order(ent.Desc(sort))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "active:":
			if active, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(ad.IsActive(active))
			}
		}
	}

	entAdList, err := q.
		Limit(params.Limit).
		Offset(params.Offset).
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelAds []*model.Ad
	for _, entAd := range entAdList {
		modelAds = append(modelAds, model.NewAd(model.NewAdParams{EntAd: entAd}))
	}

	return modelAds, nil
}

type AdDatasourceCreateParams struct {
	Ctx context.Context
	Ad  model.Ad
}

func (ds *AdDatasource) Create(params AdDatasourceCreateParams) error {
	err := ds.client.
		Ad.
		Create().
		SetContent(params.Ad.EntAd.Content).
		SetIsActive(params.Ad.EntAd.IsActive).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Exec(params.Ctx)
	if err != nil {
		return err
	}

	return nil
}

type AdDatasourceUpdateParams struct {
	Ctx context.Context
	Ad  model.Ad
}

func (ds *AdDatasource) Update(params AdDatasourceUpdateParams) (*model.Ad, error) {
	q := ds.client.
		Ad.
		UpdateOneID(params.Ad.EntAd.ID)

	q = q.
		SetContent(params.Ad.EntAd.Content).
		SetIsActive(params.Ad.EntAd.IsActive).
		SetUpdatedAt(time.Now())

	entAd, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	ad := model.NewAd(model.NewAdParams{EntAd: entAd})

	return ad, nil
}

type AdDatasourceDeleteParams struct {
	Ctx  context.Context
	AdId int
}

func (ds *AdDatasource) Delete(params AdDatasourceDeleteParams) error {
	err := ds.client.
		Ad.
		DeleteOneID(params.AdId).
		Exec(params.Ctx)
	if err != nil {
		return err
	}

	return nil
}

type AdDatasourceUpdateIsActiveParams struct {
	Ctx context.Context
	Ad  *model.Ad
}

func (ds *AdDatasource) UpdateIsActive(params AdDatasourceUpdateIsActiveParams) (*model.Ad, error) {
	q := ds.client.
		Ad.
		UpdateOneID(params.Ad.EntAd.ID)

	q = q.
		SetIsActive(params.Ad.EntAd.IsActive).
		SetUpdatedAt(time.Now())

	entAd, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	ad := model.NewAd(model.NewAdParams{
		EntAd: entAd,
	})

	return ad, nil
}
