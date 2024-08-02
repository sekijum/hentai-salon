package datasource_admin

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
)

type ThreadCommentDatasource struct {
	client *ent.Client
}

func NewThreadCommentDatasource(client *ent.Client) *ThreadCommentDatasource {
	return &ThreadCommentDatasource{client: client}
}

type ThreadCommentDatasourceUpdateParams struct {
	Ctx           context.Context
	CommentID     int
	ThreadComment model.ThreadComment
}

func (ds *ThreadCommentDatasource) Update(params ThreadCommentDatasourceUpdateParams) error {
	update := ds.client.ThreadComment.UpdateOneID(params.CommentID)
	update.SetStatus(params.ThreadComment.EntThreadComment.Status)

	_, err := update.Save(params.Ctx)
	if err != nil {
		return err
	}

	return nil
}
