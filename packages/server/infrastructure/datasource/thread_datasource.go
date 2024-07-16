package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
	"sort"
)

type ThreadDatasource struct {
	client *ent.Client
}

func NewThreadDatasource(client *ent.Client) *ThreadDatasource {
	return &ThreadDatasource{client: client}
}

type ThreadDatasourceGetCommentCountParams struct {
	Ctx      context.Context
	ThreadID int
}

func (ds *ThreadDatasource) getCommentCount(params ThreadDatasourceGetCommentCountParams) (int, error) {
	commentCount, err := ds.client.ThreadComment.Query().
		Where(threadcomment.HasThreadWith(thread.IDEQ(params.ThreadID))).
		Count(params.Ctx)

	if err != nil {
		return 0, err
	}

	return commentCount, nil
}

type ThreadDatasourceFindByBoardIDParams struct {
	Ctx                    context.Context
	BoardID, Limit, Offset int
}

func (ds *ThreadDatasource) FindByBoardID(params ThreadDatasourceFindByBoardIDParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.BoardIDEQ(params.BoardID)).
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().
		WithBoard().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
			Ctx:      params.Ctx,
			ThreadID: entThread.ID,
		})
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:    entThread,
			CommentCount: commentCount,
		})
	}

	return modelThreads, nil
}

type ThreadDatasourceFindByRelatedTagParams struct {
	Ctx           context.Context
	ThreadIDs     []int
	Limit, Offset int
}

func (ds *ThreadDatasource) FindByRelatedTag(params ThreadDatasourceFindByRelatedTagParams) ([]*model.Thread, error) {
	tags, err := ds.client.Tag.Query().
		Where(tag.HasThreadsWith(thread.IDIn(params.ThreadIDs...))).
		All(params.Ctx)

	if err != nil {
		return nil, err
	}

	var tagIds []int
	for _, t := range tags {
		tagIds = append(tagIds, t.ID)
	}

	threads, err := ds.client.Thread.Query().
		Where(
			thread.And(
				thread.HasTagsWith(tag.IDIn(tagIds...)),
				thread.Not(thread.IDIn(params.ThreadIDs...)),
			),
		).
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().
		WithBoard().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
			Ctx:      params.Ctx,
			ThreadID: entThread.ID,
		})
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:    entThread,
			CommentCount: commentCount,
		})
	}

	return modelThreads, nil
}

type ThreadDatasourceFindByKeywordParams struct {
	Ctx           context.Context
	Keyword       string
	Limit, Offset int
}

func (ds *ThreadDatasource) FindByKeyword(params ThreadDatasourceFindByKeywordParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(
			thread.Or(
				thread.TitleContainsFold(params.Keyword),
				thread.DescriptionContainsFold(params.Keyword),
			),
		).
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().
		WithBoard().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
			Ctx:      params.Ctx,
			ThreadID: entThread.ID,
		})
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:    entThread,
			CommentCount: commentCount,
		})
	}

	return modelThreads, nil
}

type ThreadDatasourceFindByPopularityParams struct {
	Ctx           context.Context
	Limit, Offset int
}

func (ds *ThreadDatasource) FindByPopularity(params ThreadDatasourceFindByPopularityParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		WithTags().
		WithBoard().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	threadWithComments := make([]struct {
		Thread       *ent.Thread
		CommentCount int
	}, len(threads))

	for i, entThread := range threads {
		commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
			Ctx:      params.Ctx,
			ThreadID: entThread.ID,
		})
		if err != nil {
			return nil, err
		}
		threadWithComments[i] = struct {
			Thread       *ent.Thread
			CommentCount int
		}{
			Thread:       entThread,
			CommentCount: commentCount,
		}
	}

	sort.Slice(threadWithComments, func(i, j int) bool {
		return threadWithComments[i].CommentCount > threadWithComments[j].CommentCount
	})

	start := params.Offset
	if start > len(threadWithComments) {
		start = len(threadWithComments)
	}
	end := params.Offset + params.Limit
	if end > len(threadWithComments) {
		end = len(threadWithComments)
	}
	threadWithComments = threadWithComments[start:end]

	var modelThreads []*model.Thread
	for _, twc := range threadWithComments {
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:    twc.Thread,
			CommentCount: twc.CommentCount,
		})
	}

	return modelThreads, nil
}

type ThreadDatasourceFindByNewestParams struct {
	Ctx           context.Context
	Limit, Offset int
}

func (ds *ThreadDatasource) FindByNewest(params ThreadDatasourceFindByNewestParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Order(ent.Desc(thread.FieldCreatedAt)).
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().
		WithBoard().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
			Ctx:      params.Ctx,
			ThreadID: entThread.ID,
		})
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:    entThread,
			CommentCount: commentCount,
		})
	}

	return modelThreads, nil
}

type ThreadDatasourceFindByHistoryParams struct {
	Ctx           context.Context
	ThreadIDs     []int
	Limit, Offset int
}

func (ds *ThreadDatasource) FindByHistory(params ThreadDatasourceFindByHistoryParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.IDIn(params.ThreadIDs...)).
		Limit(params.Limit).
		Offset(params.Offset).
		WithTags().
		WithBoard().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
			Ctx:      params.Ctx,
			ThreadID: entThread.ID,
		})
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:    entThread,
			CommentCount: commentCount,
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

func (ds *ThreadDatasource) FindById(params ThreadDatasourceFindByIDParams) (*model.Thread, error) {
	commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
	})
	if err != nil {
		return nil, err
	}

	orderFunc := ent.Desc
	if params.SortOrder == "asc" {
		orderFunc = ent.Asc
	}

	commentIDs, err := ds.client.ThreadComment.Query().
		Where(threadcomment.HasThreadWith(thread.IDEQ(params.ThreadID))).
		Order(orderFunc(threadcomment.FieldCreatedAt)).
		IDs(params.Ctx)
	if err != nil {
		return nil, err
	}

	entThread, err := ds.client.Thread.Query().
		Where(thread.IDEQ(params.ThreadID)).
		WithTags().
		WithComments(func(q *ent.ThreadCommentQuery) {
			q.Order(orderFunc(threadcomment.FieldCreatedAt)).
				Limit(params.Limit).
				Offset(params.Offset).
				WithAuthor().
				WithAttachments(func(aq *ent.ThreadCommentAttachmentQuery) {
					aq.Order(ent.Asc(threadcommentattachment.FieldDisplayOrder))
				}).
				WithReplies(func(rq *ent.ThreadCommentQuery) {
					rq.Order(orderFunc(threadcomment.FieldCreatedAt))
				})
		}).
		WithBoard().
		Only(params.Ctx)
	if err != nil {
		return nil, err
	}

	modelThread := &model.Thread{
		EntThread:    entThread,
		CommentCount: commentCount,
		CommentIDs:   commentIDs,
	}

	return modelThread, nil
}

type ThreadDatasourceFindByTitleParams struct {
	Ctx   context.Context
	Title string
}

func (ds *ThreadDatasource) FindByTitle(params ThreadDatasourceFindByTitleParams) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.TitleEQ(params.Title)).
		WithTags().
		All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		commentCount, err := ds.getCommentCount(ThreadDatasourceGetCommentCountParams{
			Ctx:      params.Ctx,
			ThreadID: entThread.ID,
		})
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:    entThread,
			CommentCount: commentCount,
		})
	}

	return modelThreads, nil
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
