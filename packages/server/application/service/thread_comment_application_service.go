package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/resource"
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

func (svc *ThreadCommentApplicationService) FindAllByUserID(params ThreadCommentApplicationServiceFindByUserIDParams) (*resource.Collection[*resource.ThreadCommentResource], error) {
	commentList, err := svc.threadCommentDatasource.FindAllByUserID(datasource.ThreadCommentDatasourceFindAllByUserIDParams{
		Ctx:    params.Ctx,
		UserID: params.UserID,
		Limit:  params.Qs.Limit,
		Offset: params.Qs.Offset,
	})
	if err != nil {
		return nil, err
	}

	commentCount, err := svc.threadCommentDatasource.GetCommentCount(datasource.ThreadDatasourceGetCommentCountParams{
		Ctx:    params.Ctx,
		UserID: &params.UserID,
	})
	if err != nil {
		return nil, err
	}

	var threadCommentResourceList []*resource.ThreadCommentResource
	for _, comment_i := range commentList {
		threadCommentResourceList = append(threadCommentResourceList, resource.NewThreadCommentResource(resource.NewThreadCommentResourceParams{
			ThreadComment: comment_i,
			ReplyCount:    len(comment_i.EntThreadComment.Edges.Replies),
		}))
	}

	dto := resource.NewCollection(resource.NewCollectionParams[*resource.ThreadCommentResource]{
		Data:       threadCommentResourceList,
		TotalCount: commentCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

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
		ReplyCount:    len(comment.EntThreadComment.Edges.Replies),
	})

	return dto, nil
}

type ThreadCommentApplicationServiceCreateParams struct {
	Ctx             context.Context
	UserID          *int
	ThreadID        int
	ClientIP        string
	ParentCommentID *int
	Body            request.ThreadCommentCreateRequest
}

func (svc *ThreadCommentApplicationService) Create(params ThreadCommentApplicationServiceCreateParams) (*resource.ThreadCommentResource, error) {

	comment := model.NewThreadComment(model.NewThreadCommentParams{
		EntThreadComment: &ent.ThreadComment{
			ThreadID:        params.ThreadID,
			Content:         params.Body.Content,
			IPAddress:       params.ClientIP,
			GuestName:       &params.Body.GuestName,
			UserID:          params.UserID,
			ParentCommentID: params.ParentCommentID,
		},
		OptionList: []func(*model.ThreadComment){
			model.WithThreadCommentStatus(model.ThreadCommentStatusVisible),
		},
	})

	if params.UserID != nil {
		comment.EntThreadComment.GuestName = nil
	}

	var attachments []*ent.ThreadCommentAttachment
	for _, attachment_i := range params.Body.Attachments {
		attachment := model.NewThreadCommentAttachment(model.NewThreadCommentAttachmentParams{
			EntAttachment: &ent.ThreadCommentAttachment{
				URL:          attachment_i.URL,
				DisplayOrder: attachment_i.DisplayOrder,
			},
			OptionList: []func(*model.ThreadCommentAttachment){
				model.WithAttachmentTypeFromString(attachment_i.Type),
			},
		},
		)
		attachments = append(attachments, attachment.EntAttachment)
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
