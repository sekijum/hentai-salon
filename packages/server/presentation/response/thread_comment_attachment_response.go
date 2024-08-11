package response

import (
	"server/domain/model"
)

type ThreadCommentAttachmentResponse struct {
	URL          string `json:"url"`
	DisplayOrder int    `json:"displayOrder"`
	Type         string `json:"type"`
	CommentID    uint64 `json:"commentId"`
}

type NewThreadCommentAttachmentResponseParams struct {
	ThreadCommentAttachment *model.ThreadCommentAttachment
}

func NewThreadCommentAttachmentResponse(params NewThreadCommentAttachmentResponseParams) *ThreadCommentAttachmentResponse {
	return &ThreadCommentAttachmentResponse{
		URL:          params.ThreadCommentAttachment.EntAttachment.URL,
		DisplayOrder: params.ThreadCommentAttachment.EntAttachment.DisplayOrder,
		Type:         params.ThreadCommentAttachment.TypeToString(),
		CommentID:    params.ThreadCommentAttachment.EntAttachment.CommentID,
	}
}
