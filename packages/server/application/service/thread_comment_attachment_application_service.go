package service

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/presentation/request"
	"server/presentation/response"
	"time"
)

type ThreadCommentAttachmentApplicationService struct {
	threadCommentAttachmentDatasource *datasource.ThreadCommentAttachmentDatasource
	tagDatasource                     *datasource.TagDatasource
}

func NewThreadCommentAttachmentApplicationService(
	threadCommentAttachmentDatasource *datasource.ThreadCommentAttachmentDatasource,
	tagDatasource *datasource.TagDatasource,
) *ThreadCommentAttachmentApplicationService {
	return &ThreadCommentAttachmentApplicationService{
		threadCommentAttachmentDatasource: threadCommentAttachmentDatasource,
		tagDatasource:                     tagDatasource,
	}
}

type ThreadCommentAttachmentApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.ThreadCommentAttachmentFindAllRequest
}

func (svc *ThreadCommentAttachmentApplicationService) FindAll(params ThreadCommentAttachmentApplicationServiceFindAllParams) ([]*response.ThreadCommentAttachmentResponse, error) {
	var attachmentList []*model.ThreadCommentAttachment

	switch params.Qs.Filter {
	case "related-by-history":
		if len(params.Qs.ThreadIDs) == 0 {
			return nil, nil
		}

		var tagIDs, err = svc.tagDatasource.FindAllIDs(datasource.TagDatasourceFindAllIDsParams{
			Ctx:       params.Ctx,
			ThreadIDs: params.Qs.ThreadIDs,
		})
		if err != nil {
			return nil, err
		}

		attachmentList, err = svc.threadCommentAttachmentDatasource.FindAllByHistory(datasource.ThreadCommentAttachmentDatasourceFindAllByHistoyParams{
			Ctx:    params.Ctx,
			TagIDs: tagIDs,
		})
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("無効なfilterです")
	}

	var dto []*response.ThreadCommentAttachmentResponse
	for _, attachment_i := range attachmentList {
		createdAt := attachment_i.EntAttachment.Edges.Comment.CreatedAt.Format(time.RFC3339)
		var commentAuthorName *string
		if attachment_i.EntAttachment.Edges.Comment.Edges.Author != nil {
			commentAuthorName = &attachment_i.EntAttachment.Edges.Comment.Edges.Author.Name
		} else if attachment_i.EntAttachment.Edges.Comment.GuestName != nil {
			commentAuthorName = attachment_i.EntAttachment.Edges.Comment.GuestName
		}

		response := response.NewThreadCommentAttachmentResponse(response.NewThreadCommentAttachmentResponseParams{
			ThreadCommentAttachment: attachment_i,
			ThreadID:                &attachment_i.EntAttachment.Edges.Comment.Edges.Thread.ID,
			CommentAuthorName:       commentAuthorName,
			CommentContent:          &attachment_i.EntAttachment.Edges.Comment.Content,
			CreatedAt:               &createdAt,
		})
		dto = append(dto, response)
	}

	return dto, nil
}
