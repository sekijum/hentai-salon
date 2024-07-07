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

func (ds *ThreadCommentDatasource) FindAll(ctx context.Context, threadId, page, limit int) ([]*model.ThreadComment, error) {
	offset := (page - 1) * limit
	comments, err := ds.client.ThreadComment.Query().
		Where(threadcomment.ThreadIDEQ(threadId)).
		Limit(limit).
		Offset(offset).
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
	commentBuilder := ds.client.ThreadComment.Create().
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
		return nil, err
	}

	return &model.ThreadComment{
		EntThreadComment: savedComment,
	}, nil
}
