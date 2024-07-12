package datasource

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/threadcomment"
)

var ErrNotFound = errors.New("not found")

type ThreadCommentDatasource struct {
	client *ent.Client
}

func NewThreadCommentDatasource(client *ent.Client) *ThreadCommentDatasource {
	return &ThreadCommentDatasource{client: client}
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

func (ds *ThreadCommentDatasource) FindById(ctx context.Context, id int) (*model.ThreadComment, error) {
	comment, err := ds.client.ThreadComment.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &model.ThreadComment{
		EntThreadComment: comment,
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
