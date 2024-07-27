package resource

import (
	"server/domain/model"
)

type ThreadCommentAttachmentResource struct {
	URL          string `json:"url"`
	DisplayOrder int    `json:"displayOrder"`
	Type         string `json:"type"`
	CommentID    int    `json:"commentId"`
}

type NewThreadCommentAttachmentResourceParams struct {
	ThreadCommentAttachment *model.ThreadCommentAttachment
}

func NewThreadCommentAttachmentResource(params NewThreadCommentAttachmentResourceParams) *ThreadCommentAttachmentResource {
	return &ThreadCommentAttachmentResource{
		URL:          params.ThreadCommentAttachment.EntAttachment.URL,
		DisplayOrder: params.ThreadCommentAttachment.EntAttachment.DisplayOrder,
		Type:         params.ThreadCommentAttachment.TypeToString(),
		CommentID:    params.ThreadCommentAttachment.EntAttachment.CommentID,
	}
}
