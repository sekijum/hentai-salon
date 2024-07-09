package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
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

func (ds *ThreadDatasource) FindById(ctx context.Context, id int, limit, offset int) (*model.Thread, error) {
	totalComments, err := ds.client.ThreadComment.Query().
		Where(threadcomment.HasThreadWith(thread.IDEQ(id))).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	allCommentIDs, err := ds.client.ThreadComment.Query().
		Where(threadcomment.HasThreadWith(thread.IDEQ(id))).
		Order(ent.Desc(threadcomment.FieldCreatedAt)).
		IDs(ctx)
	if err != nil {
		return nil, err
	}

	entThread, err := ds.client.Thread.Query().
		Where(thread.IDEQ(id)).
		WithTags().
		WithComments(func(q *ent.ThreadCommentQuery) {
			q.Order(ent.Desc(threadcomment.FieldCreatedAt)).
				Limit(limit).
				Offset(offset).
				WithAuthor().
				WithAttachments(func(aq *ent.ThreadCommentAttachmentQuery) {
					aq.Order(ent.Asc(threadcommentattachment.FieldDisplayOrder))
				}).
				WithReplies(func(rq *ent.ThreadCommentQuery) {
					rq.Order(ent.Desc(threadcomment.FieldCreatedAt))
				})
		}).
		WithBoard().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	modelThread := &model.Thread{
		EntThread:     entThread,
		TotalComments: totalComments,
		CommentIDs:    allCommentIDs,
	}

	return modelThread, nil
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
