package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"

	"server/infrastructure/ent/thread"

	"github.com/mitchellh/mapstructure"
)

type ThreadDatasource struct {
	client *ent.Client
}

func NewThreadDatasource(client *ent.Client) *ThreadDatasource {
	return &ThreadDatasource{client: client}
}

func (ds *ThreadDatasource) FindAll(ctx context.Context) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		var modelThread model.Thread
		err := mapstructure.Decode(entThread, &modelThread)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &modelThread)
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByTitle(ctx context.Context, title string) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().Where(thread.TitleEQ(title)).All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		var modelThread model.Thread
		err := mapstructure.Decode(entThread, &modelThread)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &modelThread)
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) Create(ctx context.Context, t *model.Thread) (*model.Thread, error) {
	threadBuilder := ds.client.Thread.Create().
		SetTitle(t.Title).
		SetUserId(t.UserId).
		SetBoardId(t.BoardId).
		SetIPAddress(t.IpAddress).
		SetStatus(t.Status.ToInt())
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

	var modelThread model.Thread
	err = mapstructure.Decode(savedThread, &modelThread)
	if err != nil {
		return nil, err
	}

	return &modelThread, nil
}
