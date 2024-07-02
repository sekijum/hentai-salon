package resource

import (
	"time"

	"server/domain/model"
)

type ThreadCommentResource struct {
	Id             int    `json:"id"`
	ThreadId       int    `json:"threadId"`
	ParentCommentId *int   `json:"parentCommentId,omitempty"`
	UserId         int    `json:"userId"`
	Content        string `json:"content"`
	IpAddress      string `json:"ip_address"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func NewCommentResource(c *model.ThreadComment) *ThreadCommentResource {
	return &ThreadCommentResource{
		Id:             c.Id,
		ThreadId:       c.ThreadId,
		ParentCommentId: c.ParentCommentId,
		Content:        c.Content,
		IpAddress:      c.IpAddress,
		Status:         c.Status.String(),
		CreatedAt:      c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      c.UpdatedAt.Format(time.RFC3339),
	}
}
