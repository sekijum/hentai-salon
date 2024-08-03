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
	q := ds.client.
		Tag.
		Query()

	if len(params.ThreadIDs) > 0 {
		q = q.Where(tag.HasThreadsWith(thread.IDIn(params.ThreadIDs...)))
	}

	entTagList, err := q.Select(tag.FieldID).All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var tagIDs []int
	for _, entTag_i := range entTagList {
		tagIDs = append(tagIDs, entTag_i.ID)
	}

	return tagIDs, nil
}

type TagDatasourceFindAllParams struct {
	Ctx context.Context
}

func (ds *TagDatasource) FindAll(params TagDatasourceFindAllParams) ([]*model.Tag, error) {
	entTagList, err := ds.client.
		Tag.
		Query().
		WithThreads().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var tagList []*model.Tag
	for _, entTag_i := range entTagList {
		tagList = append(tagList, model.NewTag(model.NewTagParams{EntTag: entTag_i}))
	}

	return tagList, nil
}

type TagDatasourceCreateManyTxParams struct {
	Ctx         context.Context
	Tx          *ent.Tx
	TagNameList []string
}

func (ds *TagDatasource) CreateManyTx(params TagDatasourceCreateManyTxParams) ([]*model.Tag, error) {
	var tagList []*model.Tag

	for _, tagName_i := range params.TagNameList {
		entTag, err := params.Tx.Tag.Query().Where(tag.NameEQ(tagName_i)).Only(params.Ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		if entTag == nil {
			entTag, err = params.Tx.Tag.Create().SetName(tagName_i).Save(params.Ctx)
			if err != nil {
				return nil, err
			}
		}
		tagList = append(tagList, model.NewTag(model.NewTagParams{EntTag: entTag}))
	}

	return tagList, nil
}
