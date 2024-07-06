package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"

	"github.com/mitchellh/mapstructure"
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
		modelTag, err := entTagToModelTag(entTag)
		if err != nil {
			return nil, err
		}
		modelTags = append(modelTags, modelTag)
	}

	return modelTags, nil
}

func entTagToModelTag(entTag *ent.Tag) (*model.Tag, error) {
	var modelTag model.Tag
	err := mapstructure.Decode(entTag, &modelTag)
	if err != nil {
		return nil, err
	}
	return &modelTag, nil
}
