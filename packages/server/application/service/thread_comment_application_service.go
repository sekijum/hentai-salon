package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	request "server/presentation/request"
	"time"

	"github.com/gin-gonic/gin"
)

type ThreadCommentApplicationService struct {
	threadCommentDatasource *datasource.ThreadCommentDatasource
}

func NewThreadCommentApplicationService(threadCommentDatasource *datasource.ThreadCommentDatasource) *ThreadCommentApplicationService {
	return &ThreadCommentApplicationService{threadCommentDatasource: threadCommentDatasource}
}

func (svc *ThreadCommentApplicationService) FindAll(
	ctx context.Context,
	qs request.ThreadCommentFindAllRequest,
) ([]*model.ThreadComment, error) {
	return svc.threadCommentDatasource.FindAll(ctx, qs.ThreadId, qs.Page, qs.Limit)
}

func (svc *ThreadCommentApplicationService) FindById(
	ctx context.Context,
	id int,
) (*model.ThreadComment, error) {
	return svc.threadCommentDatasource.FindById(ctx, id)
}

func (svc *ThreadCommentApplicationService) Create(
	ctx context.Context,
	ginCtx *gin.Context,
	body request.ThreadCommentCreateRequest,
) error {
	userId, exists := ginCtx.Get("user_id")

	comment := &model.ThreadComment{
		EntThreadComment: &ent.ThreadComment{
			ThreadID:  body.ThreadId,
			Content:   body.Content,
			IPAddress: ginCtx.ClientIP(),
			Status:    int(model.ThreadCommentStatusVisible),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if exists {
		userIdInt := userId.(int)
		comment.EntThreadComment.UserID = &userIdInt
	}

	if body.ParentCommentId != nil {
		comment.EntThreadComment.ParentCommentID = body.ParentCommentId
	}

	_, err := svc.threadCommentDatasource.Create(ctx, comment)
	if err != nil {
		return err
	}

	return nil
}
