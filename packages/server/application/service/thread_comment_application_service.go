package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	request "server/presentation/request"
	resource "server/presentation/resource"
	"time"

	"github.com/gin-gonic/gin"
)

type ThreadCommentApplicationService struct {
	threadCommentDatasource *datasource.ThreadCommentDatasource
}

func NewThreadCommentApplicationService(threadCommentDatasource *datasource.ThreadCommentDatasource) *ThreadCommentApplicationService {
	return &ThreadCommentApplicationService{threadCommentDatasource: threadCommentDatasource}
}

type ThreadCommentApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	CommentID int
	Qs        request.ThreadFindByIdRequest
}

func (svc *ThreadCommentApplicationService) FindByID(params ThreadCommentApplicationServiceFindByIDParams) (*resource.ThreadCommentResource, error) {
	comment, err := svc.threadCommentDatasource.FindByID(datasource.ThreadCommentDatasourceFindByIDParams{
		Ctx:       params.Ctx,
		CommentID: params.CommentID,
		Limit:     params.Qs.Limit,
		Offset:    params.Qs.Offset,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewThreadCommentResource(resource.NewThreadCommentResourceParams{
		ThreadComment: comment,
		Limit:         params.Qs.Limit,
		Offset:        params.Qs.Offset,
	})

	return dto, nil
}

type ThreadCommentApplicationServiceCreateParams struct {
	Ctx             context.Context
	GinCtx          *gin.Context
	ThreadID        int
	ParentCommentID *int
	Body            request.ThreadCommentCreateRequest
}

func (svc *ThreadCommentApplicationService) Create(params ThreadCommentApplicationServiceCreateParams) (*resource.ThreadCommentResource, error) {
	userID, exists := params.GinCtx.Get("userID")

	comment := &model.ThreadComment{
		EntThreadComment: &ent.ThreadComment{
			ThreadID:  params.ThreadID,
			Content:   params.Body.Content,
			IPAddress: params.GinCtx.ClientIP(),
			Status:    int(model.ThreadCommentStatusVisible),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if params.Body.GuestName != nil {
		comment.EntThreadComment.GuestName = params.Body.GuestName
	}

	if exists {
		userIdInt := userID.(int)
		comment.EntThreadComment.UserID = &userIdInt
		comment.EntThreadComment.GuestName = nil
	}

	if params.ParentCommentID != nil {
		comment.EntThreadComment.ParentCommentID = params.ParentCommentID
	}

	attachments := make([]*ent.ThreadCommentAttachment, len(params.Body.Attachments))
	for i, a := range params.Body.Attachments {
		attachmentType, err := model.AttachmentTypeFromString(a.Type)
		if err != nil {
			return nil, err
		}

		attachments[i] = &ent.ThreadCommentAttachment{
			URL:          a.URL,
			DisplayOrder: a.DisplayOrder,
			Type:         int(attachmentType),
			CreatedAt:    time.Now(),
		}
	}

	comment.EntThreadComment.Edges.Attachments = attachments

	_, err := svc.threadCommentDatasource.Create(datasource.ThreadCommentDatasourceCreateParams{
		Ctx:           params.Ctx,
		ThreadComment: comment,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewThreadCommentResource(resource.NewThreadCommentResourceParams{
		ThreadComment: comment,
	})

	return dto, nil
}
