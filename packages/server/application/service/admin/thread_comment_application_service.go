package service_admin

import (
	"context"
	datasource_admin "server/infrastructure/datasource/admin"
)

type ThreadCommentApplicationService struct {
	threadCommentDatasource *datasource_admin.ThreadCommentDatasource
}

func NewThreadCommentApplicationService(threadCommentDatasource *datasource_admin.ThreadCommentDatasource) *ThreadCommentApplicationService {
	return &ThreadCommentApplicationService{threadCommentDatasource: threadCommentDatasource}
}

type ThreadCommentApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	CommentID uint64
}

func (svc *ThreadCommentApplicationService) Delete(params ThreadCommentApplicationServiceFindByIDParams) error {
	err := svc.threadCommentDatasource.Delete(datasource_admin.ThreadCommentDatasourceDeleteParams{
		Ctx:       params.Ctx,
		CommentID: params.CommentID,
	})
	if err != nil {
		return err
	}

	return nil
}
