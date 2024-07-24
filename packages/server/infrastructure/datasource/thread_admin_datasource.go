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
	Ctx       context.Context
	Limit     int
	Offset    int
	SortKey   *string
	SortOrder *string
	Keyword   *string
	Status    *int
}

func (ds *ThreadAdminDatasource) FindAll(params ThreadAdminDatasourceFindAllParams) ([]*model.Thread, error) {
	query := ds.client.Thread.Query().WithBoard()

	sortKey := thread.FieldID
	if params.SortKey != nil && *params.SortKey != "" {
		sortKey = *params.SortKey
	}

	if params.SortOrder != nil && *params.SortOrder == "asc" {
		query = query.Order(ent.Asc(sortKey))
	} else {
		query = query.Order(ent.Desc(sortKey))
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

	entThreads, err := query.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range entThreads {
		modelThreads = append(modelThreads, &model.Thread{
			EntThread: entThread,
		})
	}

	return modelThreads, nil
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
	return &model.Thread{EntThread: entThread}, nil
}

type ThreadAdminDatasourceUpdateParams struct {
	Ctx    context.Context
	Thread model.Thread
}

func (ds *ThreadAdminDatasource) Update(params ThreadAdminDatasourceUpdateParams) (*model.Thread, error) {
	t, err := ds.client.Thread.Get(params.Ctx, params.Thread.EntThread.ID)
	if err != nil {
		return nil, err
	}

	update := t.Update()

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

	update.SetUpdatedAt(time.Now())

	t, err = update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	return &model.Thread{EntThread: t}, nil
}
