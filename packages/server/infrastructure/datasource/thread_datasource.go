package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
	"server/infrastructure/ent/userthreadlike"
	"time"

	"entgo.io/ent/dialect/sql"
)

type ThreadDatasource struct {
	client *ent.Client
}

func NewThreadDatasource(client *ent.Client) *ThreadDatasource {
	return &ThreadDatasource{client: client}
}

type ThreadDatasourceGetThreadCountParams struct {
	Ctx      context.Context
	UserID   *int
	ThreadID *int
}

func (ds *ThreadDatasource) GetThreadCount(params ThreadDatasourceGetCommentCountParams) (int, error) {
	q := ds.client.Thread.Query()

	if params.UserID != nil {
		q = q.Where(thread.UserID(*params.UserID))
	}

	count, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type ThreadDatasourceFindAllParams struct {
	Ctx       context.Context
	UserID    int
	BoardID   int
	Keyword   string
	ThreadIDs []int
	SortOrder string
	Limit     int
	Offset    int
}

func (ds *ThreadDatasource) FindAll(params ThreadDatasourceFindAllParams) ([]*model.Thread, error) {
	q := ds.client.Thread.Query()

	if params.UserID != 0 {
		q = q.Where(thread.UserIDEQ(params.UserID))
	}

	if params.BoardID != 0 {
		q = q.Where(thread.BoardIDEQ(params.BoardID))
	}

	if params.Keyword != "" {
		q = q.Where(
			thread.Or(
				thread.TitleContainsFold(params.Keyword),
				thread.DescriptionContainsFold(params.Keyword),
			),
		)
	}

	if len(params.ThreadIDs) > 0 {
		q = q.Where(thread.IDIn(params.ThreadIDs...))
	}

	if params.SortOrder != "" {
		orderFunc := ent.Desc
		if params.SortOrder == "asc" {
			orderFunc = ent.Asc
		}
		q = q.Order(orderFunc(thread.FieldCreatedAt))
	}

	q = q.
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().WithBoard().
		WithComments(func(rq *ent.ThreadCommentQuery) {
			rq.Select(threadcomment.FieldID)
		})

	entThreadList, err := q.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threads []*model.Thread
	for _, entThread_i := range entThreadList {
		threads = append(threads, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threads, nil
}

type ThreadDatasourceFindByRelatedTagParams struct {
	Ctx       context.Context
	ThreadIDs []int
	Limit     int
	Offset    int
	TagIDs    []int
}

func (ds *ThreadDatasource) FindByRelatedTag(params ThreadDatasourceFindByRelatedTagParams) ([]*model.Thread, error) {
	entThreadList, err := ds.client.Thread.Query().
		Where(
			thread.And(
				thread.HasTagsWith(tag.IDIn(params.TagIDs...)),
				thread.Not(thread.IDIn(params.ThreadIDs...)),
			),
		).
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().
		WithBoard().
		WithComments(func(rq *ent.ThreadCommentQuery) {
			rq.Select(threadcomment.FieldID)
		}).
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threads []*model.Thread
	for _, entThread_i := range entThreadList {
		threads = append(threads, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threads, nil
}

type ThreadDatasourceFindByPopularityParams struct {
	Ctx    context.Context
	Limit  int
	Offset int
}

func (ds *ThreadDatasource) FindByPopularity(params ThreadDatasourceFindByPopularityParams) ([]*model.Thread, error) {
	entThreadList, err := ds.client.Thread.Query().
		WithTags().
		WithBoard().
		WithComments(func(rq *ent.ThreadCommentQuery) {
			rq.Select(threadcomment.FieldID)
		}).
		Order(thread.ByCommentsCount(sql.OrderDesc())).
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threads []*model.Thread
	for _, entThread_i := range entThreadList {
		threads = append(threads, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threads, nil
}

type ThreadDatasourceFindByTitleParams struct {
	Ctx   context.Context
	Title string
}

func (ds *ThreadDatasource) FindByTitle(params ThreadDatasourceFindByTitleParams) ([]*model.Thread, error) {
	entThreadList, err := ds.client.Thread.Query().
		Where(thread.TitleEQ(params.Title)).
		WithTags().
		WithComments().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var threads []*model.Thread
	for _, entThread_i := range entThreadList {
		threads = append(threads, model.NewThread(model.NewThreadParams{EntThread: entThread_i}))
	}

	return threads, nil
}

type ThreadDatasourceFindByIDParams struct {
	Ctx           context.Context
	ThreadID      int
	SortOrder     string
	Limit, Offset int
}

func (ds *ThreadDatasource) FindById(params ThreadDatasourceFindByIDParams) (*model.Thread, error) {
	orderFunc := ent.Desc
	if params.SortOrder == "asc" {
		orderFunc = ent.Asc
	}

	entThread, err := ds.client.Thread.Query().
		Where(thread.IDEQ(params.ThreadID)).
		WithTags().
		WithBoard().
		WithComments(func(q *ent.ThreadCommentQuery) {
			q.Order(orderFunc(threadcomment.FieldCreatedAt)).
				Limit(params.Limit).
				Offset(params.Offset).
				WithAuthor().
				WithAttachments(func(aq *ent.ThreadCommentAttachmentQuery) {
					aq.Order(ent.Asc(threadcommentattachment.FieldDisplayOrder))
				}).
				WithReplies(func(rq *ent.ThreadCommentQuery) {
					rq.Select(thread.FieldID)
				}).
				WithLikedUsers()
		}).
		WithLikedUsers().
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	thread := model.NewThread(model.NewThreadParams{EntThread: entThread})

	return thread, nil
}

type ThreadDatasourceCreateTxParams struct {
	Ctx    context.Context
	Tx     *ent.Tx
	Thread *model.Thread
	TagIDs []int
}

func (ds *ThreadDatasource) CreateTx(params ThreadDatasourceCreateTxParams) (*model.Thread, error) {
	q := params.Tx.Thread.Create().
		SetTitle(params.Thread.EntThread.Title).
		SetUserID(params.Thread.EntThread.UserID).
		SetBoardID(params.Thread.EntThread.BoardID).
		SetIPAddress(params.Thread.EntThread.IPAddress).
		SetStatus(params.Thread.EntThread.Status).
		AddTagIDs(params.TagIDs...)

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

	thread := model.NewThread(model.NewThreadParams{EntThread: entThread})

	return thread, nil
}

type ThreadDatasourceLikeParams struct {
	Ctx      context.Context
	UserID   int
	ThreadID int
}

func (ds *ThreadDatasource) Like(params ThreadDatasourceLikeParams) error {
	_, err := ds.client.UserThreadLike.Create().
		SetUserID(params.UserID).
		SetThreadID(params.ThreadID).
		SetLikedAt(time.Now()).
		Save(params.Ctx)
	return err

}

type ThreadDatasourceUnlikeParams struct {
	Ctx      context.Context
	UserID   int
	ThreadID int
}

func (ds *ThreadDatasource) Unlike(params ThreadDatasourceUnlikeParams) (int, error) {
	return ds.client.UserThreadLike.Delete().
		Where(userthreadlike.UserIDEQ(params.UserID), userthreadlike.ThreadIDEQ(params.ThreadID)).
		Exec(params.Ctx)
}
