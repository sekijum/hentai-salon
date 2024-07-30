package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/thread"
	"time"
)

type ThreadAdminDatasource struct {
	client *ent.Client
}

func NewThreadAdminDatasource(client *ent.Client) *ThreadAdminDatasource {
	return &ThreadAdminDatasource{client: client}
}

type ThreadAdminDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
	Status  *int
}

func (ds *ThreadAdminDatasource) FindAll(params ThreadAdminDatasourceFindAllParams) ([]*model.Thread, error) {
	query := ds.client.Thread.Query().WithBoard()

	sort := thread.FieldID
	if params.Sort != nil && *params.Sort != "" {
		sort = *params.Sort
	}

	if params.Order != nil && *params.Order == "asc" {
		query = query.Order(ent.Asc(sort))
	} else {
		query = query.Order(ent.Desc(sort))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(thread.Or(
			thread.TitleContains(*params.Keyword),
			thread.DescriptionContains(*params.Keyword),
		))
	}

	if params.Status != nil && *params.Status != 0 {
		query = query.Where(thread.StatusEQ(*params.Status))
	}

	query = query.Limit(params.Limit)
	query = query.Offset(params.Offset)

	entThreadList, err := query.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threadList []*model.Thread
	for _, entThread_i := range entThreadList {
		threadList = append(threadList, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threadList, nil
}

type ThreadAdminDatasourceGetThreadCountParams struct {
	Ctx     context.Context
	Keyword *string
	Status  *int
}

func (ds *ThreadAdminDatasource) GetThreadCount(params ThreadAdminDatasourceGetThreadCountParams) (int, error) {
	query := ds.client.Thread.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(thread.Or(
			thread.TitleContains(*params.Keyword),
			thread.DescriptionContains(*params.Keyword),
		))
	}

	if params.Status != nil && *params.Status != 0 {
		query = query.Where(thread.StatusEQ(*params.Status))
	}

	threadCount, err := query.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return threadCount, nil
}

type ThreadAdminDatasourceFindByIDParams struct {
	Ctx      context.Context
	ThreadID int
	Limit    int
	Offset   int
}

func (ds *ThreadAdminDatasource) FindByID(params ThreadAdminDatasourceFindByIDParams) (*model.Thread, error) {
	entThread, err := ds.client.Thread.Query().
		Where(thread.ID(params.ThreadID)).
		WithComments(func(q *ent.ThreadCommentQuery) {
			q.Limit(params.Limit).Offset(params.Offset)
		}).
		WithBoard().
		WithComments().
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	thread := model.NewThread(model.NewThreadParams{EntThread: entThread})

	return thread, nil
}

type ThreadAdminDatasourceUpdateParams struct {
	Ctx    context.Context
	Thread model.Thread
}

func (ds *ThreadAdminDatasource) Update(params ThreadAdminDatasourceUpdateParams) (*model.Thread, error) {
	update := ds.client.Thread.UpdateOneID(params.Thread.EntThread.ID)

	if params.Thread.EntThread.Title != "" {
		update = update.SetTitle(params.Thread.EntThread.Title)
	}
	if params.Thread.EntThread.Description != nil {
		update = update.SetDescription(*params.Thread.EntThread.Description)
	}
	if params.Thread.EntThread.Status != 0 {
		update = update.SetStatus(params.Thread.EntThread.Status)
	}
	if params.Thread.EntThread.ThumbnailURL != nil {
		update = update.SetThumbnailURL(*params.Thread.EntThread.ThumbnailURL)
	}

	update = update.SetUpdatedAt(time.Now())

	entThread, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	thread := model.NewThread(model.NewThreadParams{
		EntThread: entThread,
	})

	return thread, nil
}
