package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	request "server/presentation/request"

	"time"

	"github.com/gin-gonic/gin"
)

type ThreadCommentApplicationService struct {
	threadCommentDatasource    *datasource.ThreadCommentDatasource
}

func NewThreadCommentApplicationService(threadCommentDatasource *datasource.ThreadCommentDatasource) *ThreadCommentApplicationService {
	return &ThreadCommentApplicationService{threadCommentDatasource: threadCommentDatasource	}
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
		ThreadId:        body.ThreadId,
		ParentCommentId: body.ParentCommentId,
		Content:         body.Content,
		IpAddress:       ginCtx.ClientIP(),
		Status:          model.ThreadCommentStatusVisible,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if exists {
		userIdInt := userId.(int)
		comment.UserId = &userIdInt
	} 

	if err := comment.Validate(); err != nil {
		return err
	}

	_, err := svc.threadCommentDatasource.Create(ctx, comment)
	if err != nil {
		return err
	}

	return nil
}
