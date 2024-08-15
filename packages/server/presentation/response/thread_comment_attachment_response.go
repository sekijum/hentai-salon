package response

import (
	"server/domain/model"
	"strconv"
)

type ThreadCommentAttachmentResponse struct {
	URL               string  `json:"url"`
	DisplayOrder      int     `json:"displayOrder"`
	Type              string  `json:"type"`
	ThreadID          *int    `json:"threadId,omitempty"`
	CommentID         *string `json:"commentId"`
	CommentAuthorName *string `json:"commentAuthorName,omitempty"`
	CommentContent    *string `json:"commentContent,omitempty"`
	CreatedAt         *string `json:"createdAt,omitempty"`
}

type NewThreadCommentAttachmentResponseParams struct {
	ThreadCommentAttachment                      *model.ThreadCommentAttachment
	CommentAuthorName, CommentContent, CreatedAt *string
	ThreadID                                     *int
}

func NewThreadCommentAttachmentResponse(params NewThreadCommentAttachmentResponseParams) *ThreadCommentAttachmentResponse {
	commentID := strconv.FormatUint(params.ThreadCommentAttachment.EntAttachment.CommentID, 10)
	return &ThreadCommentAttachmentResponse{
		URL:               params.ThreadCommentAttachment.EntAttachment.URL,
		DisplayOrder:      params.ThreadCommentAttachment.EntAttachment.DisplayOrder,
		Type:              params.ThreadCommentAttachment.TypeToString(),
		CommentID:         &commentID,
		CommentAuthorName: params.CommentAuthorName,
		CommentContent:    params.CommentContent,
		CreatedAt:         params.CreatedAt,
	}
}
