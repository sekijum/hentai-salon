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

func (ds *ThreadDatasource) getTotalComments(ctx context.Context, threadID int) (int, error) {
	totalComments, err := ds.client.ThreadComment.Query().
		Where(threadcomment.HasThreadWith(thread.IDEQ(threadID))).
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return totalComments, nil
}

func (ds *ThreadDatasource) FindByBoardId(ctx context.Context, boardId, limit, offset int) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.BoardIDEQ(boardId)).
		Limit(limit).
		Offset(offset).
		WithTags().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		totalComments, err := ds.getTotalComments(ctx, entThread.ID)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:     entThread,
			TotalComments: totalComments,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByRelatedTags(ctx context.Context, threadIds []int, limit, offset int) ([]*model.Thread, error) {
	tags, err := ds.client.Tag.Query().
		Where(tag.HasThreadsWith(thread.IDIn(threadIds...))).
		All(ctx)
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
				thread.Not(thread.IDIn(threadIds...)),
			),
		).
		Limit(limit).
		Offset(offset).
		WithTags().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		totalComments, err := ds.getTotalComments(ctx, entThread.ID)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:     entThread,
			TotalComments: totalComments,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByKeyword(ctx context.Context, keyword string, limit, offset int) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(
			thread.Or(
				thread.TitleContainsFold(keyword),
				thread.DescriptionContainsFold(keyword),
			),
		).
		Limit(limit).
		Offset(offset).
		WithTags().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		totalComments, err := ds.getTotalComments(ctx, entThread.ID)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:     entThread,
			TotalComments: totalComments,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByPopularity(ctx context.Context, limit, offset int) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		WithTags().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	threadWithComments := make([]struct {
		Thread        *ent.Thread
		TotalComments int
	}, len(threads))

	for i, entThread := range threads {
		totalComments, err := ds.getTotalComments(ctx, entThread.ID)
		if err != nil {
			return nil, err
		}
		threadWithComments[i] = struct {
			Thread        *ent.Thread
			TotalComments int
		}{
			Thread:        entThread,
			TotalComments: totalComments,
		}
	}

	sort.Slice(threadWithComments, func(i, j int) bool {
		return threadWithComments[i].TotalComments > threadWithComments[j].TotalComments
	})

	start := offset
	if start > len(threadWithComments) {
		start = len(threadWithComments)
	}
	end := offset + limit
	if end > len(threadWithComments) {
		end = len(threadWithComments)
	}
	threadWithComments = threadWithComments[start:end]

	var modelThreads []*model.Thread
	for _, twc := range threadWithComments {
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:     twc.Thread,
			TotalComments: twc.TotalComments,
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
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		totalComments, err := ds.getTotalComments(ctx, entThread.ID)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:     entThread,
			TotalComments: totalComments,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindByHistory(ctx context.Context, threadIds []int, limit, offset int) ([]*model.Thread, error) {
	threads, err := ds.client.Thread.Query().
		Where(thread.IDIn(threadIds...)).
		Limit(limit).
		Offset(offset).
		WithTags().
		WithBoard().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range threads {
		totalComments, err := ds.getTotalComments(ctx, entThread.ID)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:     entThread,
			TotalComments: totalComments,
		})
	}

	return modelThreads, nil
}

func (ds *ThreadDatasource) FindById(ctx context.Context, id int, limit, offset int) (*model.Thread, error) {
	totalComments, err := ds.getTotalComments(ctx, id)
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
		totalComments, err := ds.getTotalComments(ctx, entThread.ID)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &model.Thread{
			EntThread:     entThread,
			TotalComments: totalComments,
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
