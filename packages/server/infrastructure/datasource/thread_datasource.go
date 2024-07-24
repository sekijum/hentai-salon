package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"

	"entgo.io/ent/dialect/sql"
)

type ThreadDatasource struct {
	client *ent.Client
}

func NewThreadDatasource(client *ent.Client) *ThreadDatasource {
	return &ThreadDatasource{client: client}
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
	query := ds.client.Thread.Query()

	if params.UserID != 0 {
		query = query.Where(thread.UserIDEQ(params.UserID))
	}

	if params.BoardID != 0 {
		query = query.Where(thread.BoardIDEQ(params.BoardID))
	}

	if params.Keyword != "" {
		query = query.Where(
			thread.Or(
				thread.TitleContainsFold(params.Keyword),
				thread.DescriptionContainsFold(params.Keyword),
			),
		)
	}

	if len(params.ThreadIDs) > 0 {
		query = query.Where(thread.IDIn(params.ThreadIDs...))
	}

	if params.SortOrder != "" {
		orderFunc := ent.Desc
		if params.SortOrder == "asc" {
			orderFunc = ent.Asc
		}
		query = query.Order(orderFunc(thread.FieldCreatedAt))
	}

	query = query.Limit(params.Limit).Offset(params.Offset).WithTags().WithBoard().WithComments()

	threads, err := query.All(params.Ctx)
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

type ThreadDatasourceFindByRelatedTagParams struct {
	Ctx       context.Context
	ThreadIDs []int
	Limit     int
	Offset    int
	TagIds    []int
}

func (ds *ThreadDatasource) FindByRelatedTag(params ThreadDatasourceFindByRelatedTagParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(
			thread.And(
				thread.HasTagsWith(tag.IDIn(params.TagIds...)),
				thread.Not(thread.IDIn(params.ThreadIDs...)),
			),
		).
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().
		WithBoard().
		WithComments().
		All(params.Ctx)
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

type ThreadDatasourceFindByPopularityParams struct {
	Ctx    context.Context
	Limit  int
	Offset int
}

func (ds *ThreadDatasource) FindByPopularity(params ThreadDatasourceFindByPopularityParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		WithTags().
		WithBoard().
		Order(thread.ByCommentsCount(sql.OrderDesc())).
		WithComments().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, thread := range threads {
		modelThreads = append(modelThreads, &model.Thread{
			EntThread: thread,
		})
	}

	return modelThreads, nil
}

type ThreadDatasourceFindByTitleParams struct {
	Ctx   context.Context
	Title string
}

func (ds *ThreadDatasource) FindByTitle(params ThreadDatasourceFindByTitleParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.TitleEQ(params.Title)).
		WithTags().
		WithComments().
		All(params.Ctx)
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

type ThreadDatasourceFindByIDParams struct {
	Ctx           context.Context
	ThreadID      int
	SortOrder     string
	Limit, Offset int
}

func (ds *ThreadDatasource) FindById(params ThreadDatasourceFindByIDParams) (*model.Thread, int, error) {
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
				})
		}).
		Only(params.Ctx)
	if err != nil {
		return nil, 0, err
	}

	commentCount, err := ds.client.ThreadComment.Query().
		Where(threadcomment.ThreadID(params.ThreadID)).
		Count(params.Ctx)

	if err != nil {
		return nil, 0, err
	}

	modelThread := &model.Thread{
		EntThread: entThread,
	}

	return modelThread, commentCount, nil
}

type ThreadDatasourceCreateTxParams struct {
	Ctx    context.Context
	Tx     *ent.Tx
	Thread *model.Thread
	TagIDs []int
}

func (ds *ThreadDatasource) CreateTx(params ThreadDatasourceCreateTxParams) (*model.Thread, error) {
	threadBuilder := params.Tx.Thread.Create().
		SetTitle(params.Thread.EntThread.Title).
		SetUserID(params.Thread.EntThread.UserID).
		SetBoardID(params.Thread.EntThread.BoardID).
		SetIPAddress(params.Thread.EntThread.IPAddress).
		SetStatus(params.Thread.EntThread.Status).
		AddTagIDs(params.TagIDs...)

	if params.Thread.EntThread.Description != nil {
		threadBuilder.SetDescription(*params.Thread.EntThread.Description)
	}
	if params.Thread.EntThread.ThumbnailURL != nil {
		threadBuilder.SetThumbnailURL(*params.Thread.EntThread.ThumbnailURL)
	}

	savedThread, err := threadBuilder.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	return &model.Thread{EntThread: savedThread}, nil
}
