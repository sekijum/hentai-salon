package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"strconv"
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
}

func (ds *ThreadAdminDatasource) FindAll(params ThreadAdminDatasourceFindAllParams) ([]*model.Thread, error) {
	q := ds.client.Thread.Query().WithBoard()

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

type ThreadAdminDatasourceGetThreadCountParams struct {
	Ctx     context.Context
	Keyword *string
}

func (ds *ThreadAdminDatasource) GetThreadCount(params ThreadAdminDatasourceGetThreadCountParams) (int, error) {
	q := ds.client.Thread.Query()

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

	threadCount, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return threadCount, nil
}

type ThreadAdminDatasourceGetThreadCommentCountParams struct {
	Ctx      context.Context
	ThreadID int
	Sort     *string
	Order    *string
	Keyword  *string
}

func (ds *ThreadAdminDatasource) GetThreadCommentCount(params ThreadAdminDatasourceGetThreadCommentCountParams) (int, error) {
	q := ds.client.Thread.Query().
		Where(thread.ID(params.ThreadID)).
		QueryComments()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(threadcomment.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
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

type ThreadAdminDatasourceFindByIDParams struct {
	Ctx      context.Context
	ThreadID int
	Limit    int
	Offset   int
	Sort     *string
	Order    *string
	Keyword  *string
}

func (ds *ThreadAdminDatasource) FindByID(params ThreadAdminDatasourceFindByIDParams) (*model.Thread, error) {
	entThread, err := ds.client.Thread.Query().
		Where(thread.ID(params.ThreadID)).
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
				case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
					if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
						q = q.Where(threadcomment.StatusEQ(status))
					}
				case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
					if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
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

type ThreadAdminDatasourceUpdateParams struct {
	Ctx    context.Context
	Thread model.Thread
}

func (ds *ThreadAdminDatasource) Update(params ThreadAdminDatasourceUpdateParams) (*model.Thread, error) {
	update := ds.client.Thread.UpdateOneID(params.Thread.EntThread.ID)

	update.
		SetTitle(params.Thread.EntThread.Title).
		SetStatus(params.Thread.EntThread.Status).
		SetUpdatedAt(time.Now())

	if params.Thread.EntThread.Description != nil {
		update.SetDescription(*params.Thread.EntThread.Description)
	}

	if params.Thread.EntThread.ThumbnailURL != nil {
		update.SetThumbnailURL(*params.Thread.EntThread.ThumbnailURL)
	}

	entThread, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	thread := model.NewThread(model.NewThreadParams{
		EntThread: entThread,
	})

	return thread, nil
}
