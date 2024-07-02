package datasource

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/ent"

	"github.com/mitchellh/mapstructure"
)

var ErrNotFound = errors.New("not found")

type ThreadCommentDatasource struct {
	client *ent.Client
}

func NewThreadCommentDatasource(client *ent.Client) *ThreadCommentDatasource {
	return &ThreadCommentDatasource{client: client}
}

func (ds *ThreadCommentDatasource) FindAll(
	ctx context.Context,
	threadId, page, limit int,
) ([]*model.ThreadComment, error) {
	offset := (page - 1) * limit
	comments, err := ds.client.ThreadComment.Query().
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var modelComments []*model.ThreadComment
	for _, entComment := range comments {
		var modelComment model.ThreadComment
		err := mapstructure.Decode(entComment, &modelComment)
		if err != nil {
			return nil, err
		}
		modelComments = append(modelComments, &modelComment)
	}

	return modelComments, nil
}

func (ds *ThreadCommentDatasource) FindById(
	ctx context.Context,
	id int,
) (*model.ThreadComment, error) {
	comment, err := ds.client.ThreadComment.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var modelComment model.ThreadComment
	err = mapstructure.Decode(comment, &modelComment)
	if err != nil {
		return nil, err
	}

	return &modelComment, nil
}

func (ds *ThreadCommentDatasource) Create(ctx context.Context, c *model.ThreadComment) (*model.ThreadComment, error) {
	commentBuilder := ds.client.ThreadComment.Create().
		SetThreadID(c.ThreadId).
		SetContent(c.Content).
		SetIPAddress(c.IpAddress).
		SetStatus(c.Status.ToInt())
		if c.UserId != nil {
			commentBuilder.SetUserId(*c.UserId)
		}
		if c.ParentCommentId != nil {
			commentBuilder.SetParentCommentID(*c.ParentCommentId)
		}
		if c.GuestName != nil {
			commentBuilder.SetGuestName(*c.GuestName)
		}

	savedComment, err := commentBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	var modelComment model.ThreadComment
	err = mapstructure.Decode(savedComment, &modelComment)
	if err != nil {
		return nil, err
	}

	return &modelComment, nil
}
