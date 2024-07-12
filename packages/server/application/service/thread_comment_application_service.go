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
	return svc.threadCommentDatasource.FindAll(ctx, qs.ThreadId)
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
	threadId int,
	parentCommentId *int,
	body request.ThreadCommentCreateRequest,
) error {
	userId, exists := ginCtx.Get("user_id")

	comment := &model.ThreadComment{
		EntThreadComment: &ent.ThreadComment{
			ThreadID:  threadId,
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

	if body.GuestName != nil {
		comment.EntThreadComment.GuestName = body.GuestName
	}

	if parentCommentId != nil {
		comment.EntThreadComment.ParentCommentID = parentCommentId
	}

	attachments := make([]*ent.ThreadCommentAttachment, len(body.Attachments))
	for i, a := range body.Attachments {
		attachmentType, err := model.AttachmentTypeFromString(a.Type)
		if err != nil {
			return err
		}

		attachments[i] = &ent.ThreadCommentAttachment{
			URL:          a.URL,
			DisplayOrder: a.DisplayOrder,
			Type:         int(attachmentType),
			CreatedAt:    time.Now(),
		}
	}

	comment.EntThreadComment.Edges.Attachments = attachments

	_, err := svc.threadCommentDatasource.Create(ctx, comment)
	if err != nil {
		return err
	}

	return nil
}
