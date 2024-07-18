package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/resource"
	"time"
)

type ThreadCommentApplicationService struct {
	threadCommentDatasource *datasource.ThreadCommentDatasource
}

func NewThreadCommentApplicationService(threadCommentDatasource *datasource.ThreadCommentDatasource) *ThreadCommentApplicationService {
	return &ThreadCommentApplicationService{threadCommentDatasource: threadCommentDatasource}
}

type ThreadCommentApplicationServiceFindByUserIDParams struct {
	Ctx    context.Context
	UserID int
	Qs     request.ThreadCommentFindAllByUserIDRequest
}

func (svc *ThreadCommentApplicationService) FindAllByUserID(params ThreadCommentApplicationServiceFindByUserIDParams) (*resource.ListResource[*resource.ThreadCommentResource], error) {
	comments, commentCount, err := svc.threadCommentDatasource.FindAllByUserID(datasource.ThreadCommentDatasourceFindAllByUserIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
		Limit:  params.Qs.Limit,
		Offset: params.Qs.Offset,
	})
	if err != nil {
		return nil, err
	}

	var threadCommentResourceList []*resource.ThreadCommentResource
	for _, comment := range comments {
		threadCommentResourceList = append(threadCommentResourceList, resource.NewThreadCommentResource(resource.NewThreadCommentResourceParams{
			ThreadComment: comment,
		}))
	}

	dto := &resource.ListResource[*resource.ThreadCommentResource]{
		TotalCount: commentCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
		Data:       threadCommentResourceList,
	}

	return dto, nil
}

type ThreadCommentApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	CommentID int
	Qs        request.ThreadCommentFindByIDRequest
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
	UserID          int
	ClientIP        string
	ThreadID        int
	ParentCommentID *int
	Body            request.ThreadCommentCreateRequest
}

func (svc *ThreadCommentApplicationService) Create(params ThreadCommentApplicationServiceCreateParams) (*resource.ThreadCommentResource, error) {

	comment := &model.ThreadComment{
		EntThreadComment: &ent.ThreadComment{
			ThreadID:  params.ThreadID,
			Content:   params.Body.Content,
			IPAddress: params.ClientIP,
			Status:    int(model.ThreadCommentStatusVisible),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if params.Body.GuestName != nil {
		comment.EntThreadComment.GuestName = params.Body.GuestName
	}

	if params.UserID != 0 {
		comment.EntThreadComment.UserID = &params.UserID
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
