package datasource_admin

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
	"strconv"
)

type TagDatasource struct {
	client *ent.Client
}

func NewTagDatasource(client *ent.Client) *TagDatasource {
	return &TagDatasource{client: client}
}

type TagDatasourceGetTagCountParams struct {
	Ctx     context.Context
	Keyword *string
}

func (ds *TagDatasource) GetTagCount(params TagDatasourceGetTagCountParams) (int, error) {
	q := ds.client.Tag.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(tag.IDEQ(id))
			}
		default:
			q = q.Where(tag.Or(
				tag.NameContainsFold(*params.Keyword),
			))
		}
	}

	TagCount, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return TagCount, nil
}

type TagDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
}

func (ds *TagDatasource) FindAll(params TagDatasourceFindAllParams) ([]*model.Tag, error) {
	q := ds.client.Tag.Query().WithThreads()

	sort := tag.FieldID
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
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(tag.IDEQ(id))
			}
		default:
			q = q.Where(tag.Or(
				tag.NameContainsFold(*params.Keyword),
			))
		}
	}

	q = q.Limit(params.Limit)
	q = q.Offset(params.Offset)

	entTagList, err := q.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var tagList []*model.Tag
	for _, entTag_i := range entTagList {
		tagList = append(tagList, model.NewTag(model.NewTagParams{EntTag: entTag_i}))
	}

	return tagList, nil
}

type TagDatasourceDeleteParams struct {
	Ctx   context.Context
	TagID int
}

func (ds *TagDatasource) Delete(params TagDatasourceDeleteParams) error {
	err := ds.client.Tag.DeleteOneID(params.TagID).Exec(params.Ctx)
	if err != nil {
		return err
	}
	return nil
}
