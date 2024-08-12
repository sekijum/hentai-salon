package response

import (
	"server/domain/model"
)

type ThreadCommentAttachmentResponse struct {
	URL               string  `json:"url"`
	DisplayOrder      int     `json:"displayOrder"`
	Type              string  `json:"type"`
	CommentID         uint64  `json:"commentId"`
	CommentAuthorName *string `json:"commentAuthorName,omitempty"`
	CommentContent    *string `json:"commentContent,omitempty"`
	CreatedAt         *string `json:"createdAt,omitempty"`
}

type NewThreadCommentAttachmentResponseParams struct {
	ThreadCommentAttachment                      *model.ThreadCommentAttachment
	CommentAuthorName, CommentContent, CreatedAt *string
}

func NewThreadCommentAttachmentResponse(params NewThreadCommentAttachmentResponseParams) *ThreadCommentAttachmentResponse {
	return &ThreadCommentAttachmentResponse{
		URL:               params.ThreadCommentAttachment.EntAttachment.URL,
		DisplayOrder:      params.ThreadCommentAttachment.EntAttachment.DisplayOrder,
		Type:              params.ThreadCommentAttachment.TypeToString(),
		CommentID:         params.ThreadCommentAttachment.EntAttachment.CommentID,
		CommentAuthorName: params.CommentAuthorName,
		CommentContent:    params.CommentContent,
		CreatedAt:         params.CreatedAt,
	}
}
