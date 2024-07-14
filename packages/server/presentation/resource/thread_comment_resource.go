package resource

import (
	"server/domain/model"
	"time"
)

type CommentThreadResource struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CommentResource struct {
	Id              int                                          `json:"id"`
	GuestName       *string                                      `json:"guestName,omitempty"`
	Content         string                                       `json:"content"`
	CreatedAt       string                                       `json:"createdAt"`
	UpdatedAt       string                                       `json:"updatedAt"`
	Attachments     []*ThreadCommentAttachmentResourceForComment `json:"attachments"`
	User            *ThreadCommentUserResource                   `json:"user,omitempty"`
	Thread          *CommentThreadResource                       `json:"thread,omitempty"`
	ParentCommentID *int                                         `json:"parentCommentId"`
	ParentComment   *ThreadCommentResource                       `json:"parentComment"`
	TotalReplies    int                                          `json:"totalReplies"`
	Replies         ListResource[*ThreadCommentResource]         `json:"replies,omitempty"`
}

func NewCommentThreadResource(t *model.Thread) *CommentThreadResource {
	return &CommentThreadResource{
		Id:          t.EntThread.ID,
		Title:       t.EntThread.Title,
		Description: t.EntThread.Description,
	}
}

func NewCommentResource(c *model.ThreadComment, limit, offset int) *CommentResource {
	var user *ThreadCommentUserResource
	if c.EntThreadComment.Edges.Author != nil {
		user = NewThreadCommentUserResource(c)
	}

	var guestName *string
	if c.EntThreadComment.GuestName != nil {
		guestName = c.EntThreadComment.GuestName
	}

	var parentComment *ThreadCommentResource
	if c.EntThreadComment.ParentCommentID != nil {
		parentComment = NewThreadCommentResource(&model.ThreadComment{
			EntThreadComment: c.EntThreadComment.Edges.ParentComment,
		}, nil, 0)
	}

	var attachments []*ThreadCommentAttachmentResourceForComment
	for _, attachment := range c.EntThreadComment.Edges.Attachments {
		attachments = append(attachments, NewThreadAttachmentResourceForComment(&model.ThreadCommentAttachment{
			EntAttachment: attachment,
		}))
	}

	var replies []*ThreadCommentResource
	for i, reply := range c.EntThreadComment.Edges.Replies {
		replies = append(replies, NewThreadCommentResource(&model.ThreadComment{
			EntThreadComment: reply,
		}, nil, offset+i))
	}

	replyList := ListResource[*ThreadCommentResource]{
		TotalCount: c.ReplyCount,
		Limit:      limit,
		Offset:     offset,
		Data:       replies,
	}

	return &CommentResource{
		Id:              c.EntThreadComment.ID,
		User:            user,
		Thread:          NewCommentThreadResource(&model.Thread{EntThread: c.EntThreadComment.Edges.Thread}),
		GuestName:       guestName,
		Content:         c.EntThreadComment.Content,
		CreatedAt:       c.EntThreadComment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       c.EntThreadComment.UpdatedAt.Format(time.RFC3339),
		Attachments:     attachments,
		ParentCommentID: c.EntThreadComment.ParentCommentID,
		ParentComment:   parentComment,
		TotalReplies:    c.ReplyCount,
		Replies:         replyList,
	}
}
