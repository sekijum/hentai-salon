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

type ThreadCommentApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	CommentID int
	UserID    *int
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

	replyCount := len(comment.EntThreadComment.Edges.Replies)
	dto := resource.NewThreadCommentResource(resource.NewThreadCommentResourceParams{
		ThreadComment:        comment,
		Limit:                params.Qs.Limit,
		Offset:               params.Qs.Offset,
		UserID:               params.UserID,
		ReplyCount:           &replyCount,
		IncludeReplies:       true,
		IncludeParentComment: true,
		IncludeAttachments:   true,
		IncludeUser:          true,
		IncludeThread:        true,
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
			GuestName:       params.Body.GuestName,
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

type ThreadCommentApplicationServiceLikeParams struct {
	Ctx       context.Context
	UserID    int
	CommentID int
}

func (svc *ThreadCommentApplicationService) Like(params ThreadCommentApplicationServiceLikeParams) error {
	err := svc.threadCommentDatasource.Like(datasource.ThreadCommentDatasourceLikeParams{
		Ctx:       params.Ctx,
		UserID:    params.UserID,
		CommentID: params.CommentID,
	})
	if err != nil {
		return err
	}

	return nil
}

type ThreadCommentApplicationServiceUnLikeParams struct {
	Ctx       context.Context
	UserID    int
	CommentID int
}

func (svc *ThreadCommentApplicationService) Unlike(params ThreadCommentApplicationServiceUnLikeParams) error {
	_, err := svc.threadCommentDatasource.Unlike(datasource.ThreadCommentDatasourceUnLikeParams{
		Ctx:       params.Ctx,
		UserID:    params.UserID,
		CommentID: params.CommentID,
	})
	if err != nil {
		return err
	}

	return nil
}
