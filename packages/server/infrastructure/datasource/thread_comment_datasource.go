package datasource

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
)

var ErrNotFound = errors.New("not found")

type ThreadCommentDatasource struct {
	client *ent.Client
}

func NewThreadCommentDatasource(client *ent.Client) *ThreadCommentDatasource {
	return &ThreadCommentDatasource{client: client}
}

func (ds *ThreadCommentDatasource) getReplyCount(ctx context.Context, commentId int) (int, error) {
	replyCount, err := ds.client.ThreadComment.Query().
		Where(threadcomment.HasParentCommentWith(threadcomment.IDEQ(commentId))).
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return replyCount, nil
}

func (ds *ThreadCommentDatasource) FindAll(ctx context.Context, threadId int) ([]*model.ThreadComment, error) {
	comments, err := ds.client.ThreadComment.Query().
		Where(threadcomment.ThreadIDEQ(threadId)).
		WithReplies().
		WithAttachments().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelComments []*model.ThreadComment
	for _, entComment := range comments {
		modelComments = append(modelComments, &model.ThreadComment{
			EntThreadComment: entComment,
		})
	}

	return modelComments, nil
}

func (ds *ThreadCommentDatasource) FindById(ctx context.Context, threadId, commentId, limit, offset int) (*model.ThreadComment, error) {
	replyCount, err := ds.getReplyCount(ctx, commentId)
	if err != nil {
		return nil, err
	}

	allCommentIDs, err := ds.client.ThreadComment.Query().
		Where(threadcomment.HasThreadWith(thread.IDEQ(threadId))).
		Order(ent.Desc(threadcomment.FieldCreatedAt)).
		IDs(ctx)
	if err != nil {
		return nil, err
	}

	comment, err := ds.client.ThreadComment.Query().
		Where(threadcomment.IDEQ(commentId)).
		WithAttachments().
		WithThread().
		WithAuthor().
		WithParentComment(func(pq *ent.ThreadCommentQuery) {
			pq.WithAuthor().
				WithReplies()
		}).
		WithReplies(func(rq *ent.ThreadCommentQuery) {
			rq.Order(ent.Desc(threadcomment.FieldCreatedAt)).
				Limit(limit).
				Offset(offset).
				WithAuthor().
				WithAttachments(func(aq *ent.ThreadCommentAttachmentQuery) {
					aq.Order(ent.Asc(threadcommentattachment.FieldDisplayOrder))
				}).
				WithReplies()
		}).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &model.ThreadComment{
		EntThreadComment: comment,
		ReplyCount:       replyCount,
		RepliesIDs:       allCommentIDs,
	}, nil
}

func (ds *ThreadCommentDatasource) Create(ctx context.Context, m *model.ThreadComment) (*model.ThreadComment, error) {
	tx, err := ds.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	commentBuilder := tx.ThreadComment.Create().
		SetThreadID(m.EntThreadComment.ThreadID).
		SetContent(m.EntThreadComment.Content).
		SetIPAddress(m.EntThreadComment.IPAddress).
		SetStatus(m.EntThreadComment.Status)
	if m.EntThreadComment.UserID != nil {
		commentBuilder.SetUserID(*m.EntThreadComment.UserID)
	}
	if m.EntThreadComment.ParentCommentID != nil {
		commentBuilder.SetParentCommentID(*m.EntThreadComment.ParentCommentID)
	}
	if m.EntThreadComment.GuestName != nil {
		commentBuilder.SetGuestName(*m.EntThreadComment.GuestName)
	}

	savedComment, err := commentBuilder.Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, attachment := range m.EntThreadComment.Edges.Attachments {
		_, err := tx.ThreadCommentAttachment.Create().
			SetCommentID(savedComment.ID).
			SetURL(attachment.URL).
			SetDisplayOrder(attachment.DisplayOrder).
			SetType(attachment.Type).
			Save(ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.ThreadComment{
		EntThreadComment: savedComment,
	}, nil
}
