package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/presentation/request"
)

type ThreadCommentAdminApplicationService struct {
	threadCommentAdminDatasource *datasource.ThreadCommentAdminDatasource
}

func NewThreadCommentAdminApplicationService(threadCommentAdminDatasource *datasource.ThreadCommentAdminDatasource) *ThreadCommentAdminApplicationService {
	return &ThreadCommentAdminApplicationService{threadCommentAdminDatasource: threadCommentAdminDatasource}
}

type ThreadCommentAdminApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	CommentID int
	Body      request.ThreadCommentAdminUpdateRequest
}

func (svc *ThreadCommentAdminApplicationService) Update(params ThreadCommentAdminApplicationServiceFindByIDParams) error {
	comment := model.NewThreadComment(model.NewThreadCommentParams{
		OptionList: []func(*model.ThreadComment){
			model.WithThreadCommentStatus(model.ThreadCommentStatusVisible),
		},
	})

	err := svc.threadCommentAdminDatasource.Update(datasource.ThreadCommentAdminDatasourceUpdateParams{
		Ctx:           params.Ctx,
		CommentID:     params.CommentID,
		ThreadComment: *comment,
	})
	if err != nil {
		return err
	}

	return nil
}
