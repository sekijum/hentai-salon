package datasource_admin

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"strconv"
	"time"
)

type ThreadDatasource struct {
	client *ent.Client
}

func NewThreadDatasource(client *ent.Client) *ThreadDatasource {
	return &ThreadDatasource{client: client}
}

type ThreadDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
}

func (ds *ThreadDatasource) FindAll(params ThreadDatasourceFindAllParams) ([]*model.Thread, error) {
	q := ds.client.
		Thread.
		Query().
		Where(thread.HasBoardWith(board.StatusEQ(0))).
		WithBoard()

	sort := thread.FieldID
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
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(thread.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(thread.IDEQ(id))
			}
		default:
			q = q.Where(thread.Or(
				thread.TitleContainsFold(*params.Keyword),
				thread.DescriptionContainsFold(*params.Keyword),
			))
		}
	}

	q = q.Limit(params.Limit)
	q = q.Offset(params.Offset)

	entThreadList, err := q.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threadList []*model.Thread
	for _, entThread_i := range entThreadList {
		threadList = append(threadList, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threadList, nil
}

type ThreadDatasourceGetThreadCountParams struct {
	Ctx     context.Context
	Keyword *string
}

func (ds *ThreadDatasource) GetThreadCount(params ThreadDatasourceGetThreadCountParams) (int, error) {
	q := ds.client.Thread.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(thread.IDEQ(id))
			}
		default:
			q = q.Where(thread.Or(
				thread.TitleContainsFold(*params.Keyword),
				thread.DescriptionContainsFold(*params.Keyword),
			))
		}
	}

	threadCount, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return threadCount, nil
}

type ThreadDatasourceGetThreadCommentCountParams struct {
	Ctx      context.Context
	ThreadID int
	Sort     *string
	Order    *string
	Keyword  *string
}

func (ds *ThreadDatasource) GetThreadCommentCount(params ThreadDatasourceGetThreadCommentCountParams) (int, error) {
	q := ds.client.
		Thread.
		Query().
		Where(thread.ID(params.ThreadID)).
		QueryComments()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.ParseUint((*params.Keyword)[3:], 10, 64); err == nil {
				q = q.Where(threadcomment.IDEQ(id))
			}
		default:
			q = q.Where(threadcomment.ContentContainsFold(*params.Keyword))
		}
	}

	threadCommentCount, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return threadCommentCount, nil
}

type ThreadDatasourceFindByIDParams struct {
	Ctx      context.Context
	ThreadID int
	Limit    int
	Offset   int
	Sort     *string
	Order    *string
	Keyword  *string
}

func (ds *ThreadDatasource) FindByID(params ThreadDatasourceFindByIDParams) (*model.Thread, error) {
	entThread, err := ds.client.
		Thread.
		Query().
		Where(thread.ID(params.ThreadID)).
		Where(thread.HasBoardWith(board.StatusEQ(0))).
		WithComments(func(q *ent.ThreadCommentQuery) {
			sort := user.FieldID
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
					if id, err := strconv.ParseUint((*params.Keyword)[3:], 10, 64); err == nil {
						q = q.Where(threadcomment.IDEQ(id))
					}
				default:
					q = q.Where(threadcomment.Or(
						threadcomment.ContentContainsFold(*params.Keyword),
					))
				}
			}

			q = q.Limit(params.Limit)
			q = q.Offset(params.Offset)

			q.All(params.Ctx)
		}).
		WithBoard().
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	thread := model.NewThread(model.NewThreadParams{EntThread: entThread})

	return thread, nil
}

type ThreadDatasourceUpdateParams struct {
	Ctx    context.Context
	Thread model.Thread
}

func (ds *ThreadDatasource) Update(params ThreadDatasourceUpdateParams) (*model.Thread, error) {
	q := ds.client.Thread.
		UpdateOneID(params.Thread.EntThread.ID)

	q.SetTitle(params.Thread.EntThread.Title).
		SetStatus(params.Thread.EntThread.Status).
		SetUpdatedAt(time.Now())

	if params.Thread.EntThread.Description != nil {
		q.SetDescription(*params.Thread.EntThread.Description)
	}

	if params.Thread.EntThread.ThumbnailURL != nil {
		q.SetThumbnailURL(*params.Thread.EntThread.ThumbnailURL)
	}

	entThread, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	thread := model.NewThread(model.NewThreadParams{
		EntThread: entThread,
	})

	return thread, nil
}

type ThreadDatasourceUpdateStatusParams struct {
	Ctx    context.Context
	Thread model.Thread
}

func (ds *ThreadDatasource) UpdateStatus(params ThreadDatasourceUpdateStatusParams) (*model.Thread, error) {
	update := ds.client.
		Thread.
		UpdateOneID(params.Thread.EntThread.ID)

	update.
		SetStatus(params.Thread.EntThread.Status).
		SetUpdatedAt(time.Now())

	entThread, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	thread := model.NewThread(model.NewThreadParams{
		EntThread: entThread,
	})

	return thread, nil
}

type ThreadDatasourceDeleteParams struct {
	Ctx      context.Context
	ThreadId int
}

func (ds *ThreadDatasource) Delete(params ThreadDatasourceDeleteParams) error {
	err := ds.client.
		Thread.
		DeleteOneID(params.ThreadId).
		Exec(params.Ctx)
	if err != nil {
		return err
	}

	return nil
}
