package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
)

type ThreadCommentAdminDatasource struct {
	client *ent.Client
}

func NewThreadCommentAdminDatasource(client *ent.Client) *ThreadCommentAdminDatasource {
	return &ThreadCommentAdminDatasource{client: client}
}

type ThreadCommentAdminDatasourceUpdateParams struct {
	Ctx           context.Context
	CommentID     int
	ThreadComment model.ThreadComment
}

func (ds *ThreadCommentAdminDatasource) Update(params ThreadCommentAdminDatasourceUpdateParams) error {
	update := ds.client.ThreadComment.UpdateOneID(params.CommentID)
	update.SetStatus(params.ThreadComment.EntThreadComment.Status)

	_, err := update.Save(params.Ctx)
	if err != nil {
		return err
	}

	return nil
}
