package resource

import (
	"time"

	"server/domain/model"
)

type ThreadCommentResource struct {
	Id              int    `json:"id"`
	ThreadId        int    `json:"threadId"`
	ParentCommentId *int   `json:"parentCommentId,omitempty"`
	UserId          int    `json:"userId"`
	Content         string `json:"content"`
	IpAddress       string `json:"ipAddress"`
	Status          string `json:"status"`
	CreatedAt       string `json:"createdAt"`
}

func NewThreadCommentResource(c *model.ThreadComment) *ThreadCommentResource {
	return &ThreadCommentResource{
		Id:              c.EntThreadComment.ID,
		ThreadId:        c.EntThreadComment.ThreadID,
		ParentCommentId: c.EntThreadComment.ParentCommentID,
		Content:         c.EntThreadComment.Content,
		IpAddress:       c.EntThreadComment.IPAddress,
		Status:          c.StatusToString(),
		CreatedAt:       c.EntThreadComment.CreatedAt.Format(time.RFC3339),
	}
}
