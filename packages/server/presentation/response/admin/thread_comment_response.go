package response_admin

import (
	"server/domain/model"
	"strconv"
	"time"
)

type ThreadCommentResponse struct {
	ID        string `json:"id"`
	ThreadID  int    `json:"threadId"`
	UserID    *int   `json:"userId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type NewThreadCommentResponseParams struct {
	ThreadComment                 *model.ThreadComment
	Limit, Offset, CommentCount   int
	IncludeComments, IncludeBoard bool
}

func NewThreadCommentResponse(params NewThreadCommentResponseParams) *ThreadCommentResponse {

	return &ThreadCommentResponse{
		ID:        strconv.FormatUint(params.ThreadComment.EntThreadComment.ID, 10),
		Content:   params.ThreadComment.EntThreadComment.Content,
		CreatedAt: params.ThreadComment.EntThreadComment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: params.ThreadComment.EntThreadComment.UpdatedAt.Format(time.RFC3339),
	}
}
