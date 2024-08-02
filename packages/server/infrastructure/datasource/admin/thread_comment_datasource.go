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

type ThreadCommentDatasourceDeleteParams struct {
	Ctx           context.Context
	CommentID     int
	ThreadComment model.ThreadComment
}

func (ds *ThreadCommentDatasource) Delete(params ThreadCommentDatasourceDeleteParams) error {
	err := ds.client.ThreadComment.DeleteOneID(params.CommentID).Exec(params.Ctx)
	if err != nil {
		return err
	}
	return nil
}
