package resource

import (
	"server/domain/model"
	"time"
)

type ThreadCommentAttachmentResource struct {
	Url          string `json:"url"`
	DisplayOrder int    `json:"displayOrder"`
	Type         string `json:"type"`
	CommentID    int    `json:"commentId"`
	IDx          int    `json:"idx"`
}

type ThreadCommentResource struct {
	ID               int                                  `json:"id"`
	IDx              int                                  `json:"idx,omitempty"`
	User             *UserResource                        `json:"user,omitempty"`
	GuestName        *string                              `json:"guestName,omitempty"`
	Content          string                               `json:"content"`
	ParentCommentID  int                                  `json:"parentCommentId"`
	ParentCommentIDx int                                  `json:"parentCommentIdx,omitempty"`
	CreatedAt        string                               `json:"createdAt"`
	UpdatedAt        string                               `json:"updatedAt"`
	Thread           *ThreadResource                      `json:"thread,omitempty"`
	ParentComment    *ThreadCommentResource               `json:"parentComment"`
	Attachments      []*ThreadCommentAttachmentResource   `json:"attachments"`
	ReplyCount       int                                  `json:"replyCount"`
	Replies          ListResource[*ThreadCommentResource] `json:"replies"`
}

type NewThreadCommentResourceParams struct {
	ThreadComment *model.ThreadComment
	CommentIDs    []int
	Limit, Offset int
	IDx           *int
}

func NewThreadCommentResource(params NewThreadCommentResourceParams) *ThreadCommentResource {
	var IDx int
	if params.IDx != nil {
		IDx = params.Offset + *params.IDx + 1
	}

	var user *UserResource
	if params.ThreadComment.EntThreadComment.Edges.Author != nil {
		user = &UserResource{
			ID:          params.ThreadComment.EntThreadComment.Edges.Author.ID,
			Name:        params.ThreadComment.EntThreadComment.Edges.Author.Name,
			ProfileLink: params.ThreadComment.EntThreadComment.Edges.Author.ProfileLink,
		}
	}

	var parentCommentID int
	var parentCommentIDx int
	if params.ThreadComment.EntThreadComment.ParentCommentID != nil {
		parentCommentID = *params.ThreadComment.EntThreadComment.ParentCommentID
		if len(params.CommentIDs) > 0 {
			parentCommentIDx = model.FindCommentIndexByID(params.CommentIDs, parentCommentID) + 1
		} else {
			parentCommentIDx = 0
		}
	}

	var parentComment *ThreadCommentResource
	if params.ThreadComment.EntThreadComment.Edges.ParentComment != nil {
		parentComment = NewThreadCommentResource(NewThreadCommentResourceParams{
			ThreadComment: &model.ThreadComment{EntThreadComment: params.ThreadComment.EntThreadComment.Edges.ParentComment},
		})
	}

	var attachments []*ThreadCommentAttachmentResource
	for _, attachment := range params.ThreadComment.EntThreadComment.Edges.Attachments {
		threadCommentAttachment := &model.ThreadCommentAttachment{EntAttachment: attachment}
		attachments = append(attachments, &ThreadCommentAttachmentResource{
			Url:          threadCommentAttachment.EntAttachment.URL,
			DisplayOrder: threadCommentAttachment.EntAttachment.DisplayOrder,
			Type:         threadCommentAttachment.TypeToString(),
		})
	}

	var replies []*ThreadCommentResource
	for i, reply := range params.ThreadComment.EntThreadComment.Edges.Replies {
		commentResource := NewThreadCommentResource(NewThreadCommentResourceParams{
			ThreadComment: &model.ThreadComment{EntThreadComment: reply},
			Offset:        params.Offset,
			IDx:           &i,
		})
		replies = append(replies, commentResource)
	}

	replyList := ListResource[*ThreadCommentResource]{
		TotalCount: params.ThreadComment.ReplyCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
		Data:       replies,
	}

	var thread *ThreadResource
	if params.ThreadComment.EntThreadComment.Edges.Thread != nil {
		var description *string
		if params.ThreadComment.EntThreadComment.Edges.Thread.Description != nil {
			description = params.ThreadComment.EntThreadComment.Edges.Thread.Description
		}
		thread = &ThreadResource{
			ID:          params.ThreadComment.EntThreadComment.Edges.Thread.ID,
			Title:       params.ThreadComment.EntThreadComment.Edges.Thread.Title,
			Description: description,
		}
	}

	var guestName *string
	if params.ThreadComment.EntThreadComment.GuestName != nil {
		guestName = params.ThreadComment.EntThreadComment.GuestName
	}

	return &ThreadCommentResource{
		ID:               params.ThreadComment.EntThreadComment.ID,
		IDx:              IDx,
		User:             user,
		GuestName:        guestName,
		Content:          params.ThreadComment.EntThreadComment.Content,
		ParentCommentID:  parentCommentID,
		ParentCommentIDx: parentCommentIDx,
		ParentComment:    parentComment,
		CreatedAt:        params.ThreadComment.EntThreadComment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        params.ThreadComment.EntThreadComment.UpdatedAt.Format(time.RFC3339),
		Thread:           thread,
		Attachments:      attachments,
		ReplyCount:       len(params.ThreadComment.EntThreadComment.Edges.Replies),
		Replies:          replyList,
	}
}
