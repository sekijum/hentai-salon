package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/threadcommentattachment"
)

type ThreadCommentDatasource struct {
	client *ent.Client
}

func NewThreadCommentDatasource(client *ent.Client) *ThreadCommentDatasource {
	return &ThreadCommentDatasource{client: client}
}

type ThreadCommentDatasourceFindByIDParams struct {
	Ctx                      context.Context
	CommentID, Limit, Offset int
}

func (ds *ThreadCommentDatasource) FindByID(params ThreadCommentDatasourceFindByIDParams) (*model.ThreadComment, error) {
	comment, err := ds.client.ThreadComment.Query().
		Where(threadcomment.IDEQ(params.CommentID)).
		WithAttachments().
		WithThread().
		WithAuthor().
		WithParentComment(func(pq *ent.ThreadCommentQuery) {
			pq.WithAuthor()
		}).
		WithReplies(func(rq *ent.ThreadCommentQuery) {
			rq.Order(ent.Desc(threadcomment.FieldCreatedAt)).
				Limit(params.Limit).
				Offset(params.Offset).
				WithAuthor().
				WithAttachments(func(aq *ent.ThreadCommentAttachmentQuery) {
					aq.Order(ent.Asc(threadcommentattachment.FieldDisplayOrder))
				})
		}).
		Only(params.Ctx)

	if err != nil {
		return nil, err
	}

	replyCount, err := ds.client.ThreadComment.
		Query().
		Where(threadcomment.ParentCommentID(params.CommentID)).
		Count(params.Ctx)
	if err != nil {
		return nil, err
	}

	return &model.ThreadComment{EntThreadComment: comment, ReplyCount: replyCount}, nil
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

	commentBuilder := tx.ThreadComment.Create().
		SetThreadID(params.ThreadComment.EntThreadComment.ThreadID).
		SetContent(params.ThreadComment.EntThreadComment.Content).
		SetIPAddress(params.ThreadComment.EntThreadComment.IPAddress).
		SetStatus(params.ThreadComment.EntThreadComment.Status)

	if params.ThreadComment.EntThreadComment.UserID != nil {
		commentBuilder.SetUserID(*params.ThreadComment.EntThreadComment.UserID)
	}
	if params.ThreadComment.EntThreadComment.ParentCommentID != nil {
		commentBuilder.SetParentCommentID(*params.ThreadComment.EntThreadComment.ParentCommentID)
	}
	if params.ThreadComment.EntThreadComment.GuestName != nil {
		commentBuilder.SetGuestName(*params.ThreadComment.EntThreadComment.GuestName)
	}

	savedComment, err := commentBuilder.Save(params.Ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, attachment := range params.ThreadComment.EntThreadComment.Edges.Attachments {
		_, err := tx.ThreadCommentAttachment.Create().
			SetCommentID(savedComment.ID).
			SetURL(attachment.URL).
			SetDisplayOrder(attachment.DisplayOrder).
			SetType(attachment.Type).
			Save(params.Ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.ThreadComment{EntThreadComment: savedComment}, nil
}
