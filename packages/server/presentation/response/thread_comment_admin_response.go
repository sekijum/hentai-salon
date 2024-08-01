package response

import (
	"server/domain/model"
	"time"
)

type ThreadCommentAdminResponse struct {
	ID          int    `json:"id"`
	ThreadID    int    `json:"threadId"`
	UserID      *int   `json:"userId"`
	Content     string `json:"content"`
	Status      int    `json:"status"`
	StatusLabel string `json:"statusLabel"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type NewThreadCommentAdminResponseParams struct {
	ThreadComment                 *model.ThreadComment
	Limit, Offset, CommentCount   int
	IncludeComments, IncludeBoard bool
}

func NewThreadCommentAdminResponse(params NewThreadCommentAdminResponseParams) *ThreadCommentAdminResponse {

	return &ThreadCommentAdminResponse{
		ID:          params.ThreadComment.EntThreadComment.ID,
		Content:     params.ThreadComment.EntThreadComment.Content,
		Status:      params.ThreadComment.EntThreadComment.Status,
		StatusLabel: params.ThreadComment.ThreadCommentToLabel(),
		CreatedAt:   params.ThreadComment.EntThreadComment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   params.ThreadComment.EntThreadComment.UpdatedAt.Format(time.RFC3339),
	}
}
