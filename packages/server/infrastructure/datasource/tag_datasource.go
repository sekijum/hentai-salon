package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
)

type TagDatasource struct {
	client *ent.Client
}

func NewTagDatasource(client *ent.Client) *TagDatasource {
	return &TagDatasource{client: client}
}

type TagDatasourceFindAllIDsParams struct {
	Ctx       context.Context
	ThreadIDs []int
}

func (ds *TagDatasource) FindAllIDs(params TagDatasourceFindAllIDsParams) ([]int, error) {
	query := ds.client.Tag.Query()

	if len(params.ThreadIDs) > 0 {
		query = query.Where(tag.HasThreadsWith(thread.IDIn(params.ThreadIDs...)))
	}

	tags, err := query.Select(tag.FieldID).All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var tagIDs []int
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	return tagIDs, nil
}

type TagDatasourceFindAllParams struct {
	Ctx context.Context
}

func (ds *TagDatasource) FindAll(params TagDatasourceFindAllParams) ([]*model.Tag, error) {
	tags, err := ds.client.Tag.Query().All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelTags []*model.Tag
	for _, entTag := range tags {
		modelTags = append(modelTags, &model.Tag{EntTag: entTag})
	}

	return modelTags, nil
}

type TagDatasourceCreateManyTxParams struct {
	Ctx         context.Context
	Tx          *ent.Tx
	TagNameList []string
}

func (ds *TagDatasource) CreateManyTx(params TagDatasourceCreateManyTxParams) ([]*model.Tag, error) {
	var modelTags []*model.Tag

	for _, tagName := range params.TagNameList {
		entTag, err := params.Tx.Tag.Query().Where(tag.NameEQ(tagName)).Only(params.Ctx)
		if entTag == nil {
			entTag, err = params.Tx.Tag.Create().SetName(tagName).Save(params.Ctx)
			if err != nil {
				return nil, err
			}
		}
		modelTags = append(modelTags, &model.Tag{EntTag: entTag})
	}

	return modelTags, nil
}
