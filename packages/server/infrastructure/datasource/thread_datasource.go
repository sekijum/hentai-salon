package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/thread"
)

type ThreadDatasource struct {
	client *ent.Client
}

func NewThreadDatasource(client *ent.Client) *ThreadDatasource {
	return &ThreadDatasource{client: client}
}

func (ds *ThreadDatasource) FindByPopularity(ctx context.Context, limit, offset int) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Limit(limit).
		Offset(offset).
		WithTags().
		WithComments().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		modelThreads = append(modelThreads, &model.Thread{
			EntThread: entThread,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByNewest(ctx context.Context, limit, offset int) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Order(ent.Desc(thread.FieldCreatedAt)).
		Limit(limit).
		Offset(offset).
		WithTags().
		WithComments().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		modelThreads = append(modelThreads, &model.Thread{
			EntThread: entThread,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByHistories(ctx context.Context, threadIds []int) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.IDIn(threadIds...)).
		WithTags().
		WithComments().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		modelThreads = append(modelThreads, &model.Thread{
			EntThread: entThread,
		})
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
		modelThreads = append(modelThreads, &model.Thread{
			EntThread: entThread,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) CreateTx(ctx context.Context, tx *ent.Tx, m *model.Thread, tagIDs []int) (*model.Thread, error) {
	threadBuilder := tx.Thread.Create().
		SetTitle(m.EntThread.Title).
		SetUserID(m.EntThread.UserID).
		SetBoardID(m.EntThread.BoardID).
		SetIPAddress(m.EntThread.IPAddress).
		SetStatus(m.EntThread.Status).
		AddTagIDs(tagIDs...)
	if m.EntThread.Description != "" {
		threadBuilder.SetDescription(m.EntThread.Description)
	}
	if m.EntThread.ThumbnailURL != "" {
		threadBuilder.SetThumbnailURL(m.EntThread.ThumbnailURL)
	}

	savedThread, err := threadBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Thread{
		EntThread: savedThread,
	}, nil
}
