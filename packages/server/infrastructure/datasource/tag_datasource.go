package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
)

type TagDatasource struct {
	client *ent.Client
}

func NewTagDatasource(client *ent.Client) *TagDatasource {
	return &TagDatasource{client: client}
}

func (ds *TagDatasource) FindAll(ctx context.Context) ([]*model.Tag, error) {
	tags, err := ds.client.Tag.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var modelTags []*model.Tag
	for _, entTag := range tags {
		modelTags = append(modelTags, &model.Tag{
			EntTag: entTag,
		})
	}

	return modelTags, nil
}

func (ds *TagDatasource) CreateManyTx(ctx context.Context, tx *ent.Tx, tagNames []string) ([]*model.Tag, error) {
	var modelTags []*model.Tag

	for _, tagName := range tagNames {
		entTag, err := tx.Tag.Query().Where(tag.NameEQ(tagName)).Only(ctx)
		if entTag == nil {
			entTag, err = tx.Tag.Create().SetName(tagName).Save(ctx)
			if err != nil {
				return nil, err
			}
		}
		modelTags = append(modelTags, &model.Tag{
			EntTag: entTag,
		})
	}

	return modelTags, nil
}
