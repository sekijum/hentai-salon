package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
	"server/infrastructure/ent/usercommentlike"
	"time"
)

type ThreadCommentDatasource struct {
	client *ent.Client
}

func NewThreadCommentDatasource(client *ent.Client) *ThreadCommentDatasource {
	return &ThreadCommentDatasource{client: client}
}

type ThreadDatasourceGetCommentCountParams struct {
	Ctx      context.Context
	UserID   *int
	ThreadID *int
}

func (ds *ThreadCommentDatasource) GetCommentCount(params ThreadDatasourceGetCommentCountParams) (int, error) {
	q := ds.client.ThreadComment.Query()

	if params.ThreadID != nil {
		q = q.Where(threadcomment.ThreadID(*params.ThreadID))
	}

	if params.UserID != nil {
		q = q.Where(threadcomment.UserID(*params.UserID))
	}

	count, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type ThreadCommentDatasourceFindByIDParams struct {
	Ctx                      context.Context
	UserID                   *int
	CommentID, Limit, Offset int
}

func (ds *ThreadCommentDatasource) FindByID(params ThreadCommentDatasourceFindByIDParams) (*model.ThreadComment, error) {
	entComment, err := ds.client.ThreadComment.Query().
		Where(threadcomment.IDEQ(params.CommentID)).
		WithAttachments().
		WithThread().
		WithAuthor().
		WithParentComment(func(pq *ent.ThreadCommentQuery) {
			pq.WithAuthor()
		}).
		WithReplies(func(q *ent.ThreadCommentQuery) {
			q.Order(ent.Desc(threadcomment.FieldCreatedAt)).
				Limit(params.Limit).
				Offset(params.Offset).
				WithAuthor().
				WithReplies(func(q *ent.ThreadCommentQuery) {
					q.Select(threadcomment.FieldID)
				}).
				WithAttachments(func(aq *ent.ThreadCommentAttachmentQuery) {
					aq.Order(ent.Asc(threadcommentattachment.FieldDisplayOrder))
				}).
				WithLikedUsers()
		}).
		WithLikedUsers().
		Only(params.Ctx)

	if err != nil {
		return nil, err
	}

	comment := model.NewThreadComment(model.NewThreadCommentParams{EntThreadComment: entComment})

	return comment, nil
}

type ThreadCommentDatasourceCreateParams struct {
	Ctx           context.Context
	ThreadComment *model.ThreadComment
}

func (ds *ThreadCommentDatasource) Create(params ThreadCommentDatasourceCreateParams) (*model.ThreadComment, error) {
	tx, err := ds.client.Tx(params.Ctx)
	if err != nil {
		return nil, err
	}

	q := tx.ThreadComment.Create().
		SetThreadID(params.ThreadComment.EntThreadComment.ThreadID).
		SetContent(params.ThreadComment.EntThreadComment.Content).
		SetIPAddress(params.ThreadComment.EntThreadComment.IPAddress).
		SetStatus(params.ThreadComment.EntThreadComment.Status)

	if params.ThreadComment.EntThreadComment.UserID != nil {
		q.SetUserID(*params.ThreadComment.EntThreadComment.UserID)
	}
	if params.ThreadComment.EntThreadComment.ParentCommentID != nil {
		q.SetParentCommentID(*params.ThreadComment.EntThreadComment.ParentCommentID)
	}
	if params.ThreadComment.EntThreadComment.GuestName != nil {
		q.SetGuestName(*params.ThreadComment.EntThreadComment.GuestName)
	}

	entComment, err := q.Save(params.Ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, attachment_i := range params.ThreadComment.EntThreadComment.Edges.Attachments {
		_, err := tx.ThreadCommentAttachment.Create().
			SetCommentID(entComment.ID).
			SetURL(attachment_i.URL).
			SetDisplayOrder(attachment_i.DisplayOrder).
			SetType(attachment_i.Type).
			Save(params.Ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	comment := model.NewThreadComment(model.NewThreadCommentParams{EntThreadComment: entComment})

	return comment, nil
}

type ThreadCommentDatasourceLikeParams struct {
	Ctx       context.Context
	UserID    int
	CommentID int
}

func (ds *ThreadCommentDatasource) Like(params ThreadCommentDatasourceLikeParams) error {
	_, err := ds.client.UserCommentLike.Create().
		SetUserID(params.UserID).
		SetCommentID(params.CommentID).
		SetLikedAt(time.Now()).
		Save(params.Ctx)
	return err
}

type ThreadCommentDatasourceUnLikeParams struct {
	Ctx       context.Context
	UserID    int
	CommentID int
}

func (ds *ThreadCommentDatasource) Unlike(params ThreadCommentDatasourceUnLikeParams) (int, error) {
	return ds.client.UserCommentLike.Delete().
		Where(usercommentlike.UserIDEQ(params.UserID), usercommentlike.CommentIDEQ(params.CommentID)).
		Exec(params.Ctx)
}
