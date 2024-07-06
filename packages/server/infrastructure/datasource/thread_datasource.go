package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"

	"github.com/mitchellh/mapstructure"
)

type ThreadDatasource struct {
	client *ent.Client
}

func NewThreadDatasource(client *ent.Client) *ThreadDatasource {
	return &ThreadDatasource{client: client}
}

func (ds *ThreadDatasource) FindAll(
	ctx context.Context,
	limit int, 
	offset int,
	) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Limit(limit).
		Offset(offset).
		WithTags().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		modelThread, err := entThreadToModelThread(entThread)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, modelThread)
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByTitle(ctx context.Context, title string) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.TitleEQ(title)).
		WithTags().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		modelThread, err := entThreadToModelThread(entThread)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, modelThread)
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) Create(ctx context.Context, t *model.Thread, tagNameList []string) (*model.Thread, error) {
	tx, err := ds.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var tags []*ent.Tag
	for _, tagName := range tagNameList {
		tag, err := tx.Tag.Query().Where(tag.NameEQ(tagName)).Only(ctx)
		if err != nil {
			tag, err = tx.Tag.Create().SetName(tagName).Save(ctx)
			if err != nil {
				return nil, err
			}
		}
		tags = append(tags, tag)
	}

	tagIDs := make([]int, len(tags))
	for i, tag := range tags {
		tagIDs[i] = tag.ID
	}

	threadBuilder := tx.Thread.Create().
		SetTitle(t.Title).
		SetUserId(t.UserId).
		SetBoardId(t.BoardId).
		SetIPAddress(t.IpAddress).
		SetStatus(t.Status.ToInt()).
		AddTagIDs(tagIDs...)
	if t.Description != nil {
		threadBuilder.SetDescription(*t.Description)
	}
	if t.ThumbnailUrl != nil {
		threadBuilder.SetThumbnailUrl(*t.ThumbnailUrl)
	}

	savedThread, err := threadBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	modelThread, err := entThreadToModelThread(savedThread)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return modelThread, nil
}

func entThreadToModelThread(entThread *ent.Thread) (*model.Thread, error) {
	var modelThread model.Thread
	err := mapstructure.Decode(entThread, &modelThread)
	if err != nil {
		return nil, err
	}

	var modelTags []*model.Tag
	for _, entTag := range entThread.Edges.Tags {
		var modelTag model.Tag
		err := mapstructure.Decode(entTag, &modelTag)
		if err != nil {
			return nil, err
		}
		modelTags = append(modelTags, &modelTag)
	}
	modelThread.Tags = modelTags

	return &modelThread, nil
}
