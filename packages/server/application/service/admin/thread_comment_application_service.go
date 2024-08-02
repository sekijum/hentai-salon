package service_admin

import (
	"context"
	"server/domain/model"
	datasource_admin "server/infrastructure/datasource/admin"
	request_admin "server/presentation/request/admin"
)

type ThreadCommentApplicationService struct {
	threadCommentDatasource *datasource_admin.ThreadCommentDatasource
}

func NewThreadCommentApplicationService(threadCommentDatasource *datasource_admin.ThreadCommentDatasource) *ThreadCommentApplicationService {
	return &ThreadCommentApplicationService{threadCommentDatasource: threadCommentDatasource}
}

type ThreadCommentApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	CommentID int
	Body      request_admin.ThreadCommentUpdateRequest
}

func (svc *ThreadCommentApplicationService) Update(params ThreadCommentApplicationServiceFindByIDParams) error {
	comment := model.NewThreadComment(model.NewThreadCommentParams{
		OptionList: []func(*model.ThreadComment){
			model.WithThreadCommentStatus(model.ThreadCommentStatusVisible),
		},
	})

	err := svc.threadCommentDatasource.Update(datasource_admin.ThreadCommentDatasourceUpdateParams{
		Ctx:           params.Ctx,
		CommentID:     params.CommentID,
		ThreadComment: *comment,
	})
	if err != nil {
		return err
	}

	return nil
}
