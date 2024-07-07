package resource

import (
	"fmt"
	"server/domain/model"
	"time"
)

type ThreadBoardResource struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ThreadCommentAttachmentResource struct {
	Url          string `json:"url"`
	DisplayOrder int    `json:"displayOrder"`
	Type         string `json:"type"`
}

type ThreadCommentResource struct {
	Id              int                                `json:"id"`
	UserId          *int                               `json:"userId,omitempty"`
	GuestName       *string                            `json:"guestName,omitempty"`
	Content         string                             `json:"content"`
	ParentCommentID int                                `json:"parentCommentID"`
	CreatedAt       string                             `json:"createdAt"`
	UpdatedAt       string                             `json:"updatedAt"`
	Attachments     []*ThreadCommentAttachmentResource `json:"attachments"`
	TotalReplies    int                                `json:"totalReplies"`
}

type ThreadResource struct {
	Id           int                      `json:"id"`
	Board        *ThreadBoardResource     `json:"board"`
	Title        string                   `json:"title"`
	Description  string                   `json:"description"`
	ThumbnailUrl string                   `json:"thumbnailUrl"`
	Tags         []string                 `json:"tags"`
	CreatedAt    string                   `json:"createdAt"`
	CommentCount int                      `json:"commentCount"`
	Popularity   string                   `json:"popularity"`
	Comments     []*ThreadCommentResource `json:"comments"`
}

func NewThreadBoardResource(b *model.Board) *ThreadBoardResource {
	return &ThreadBoardResource{
		Id:    b.EntBoard.ID,
		Title: b.EntBoard.Title,
	}
}

func NewThreadAttachmentResource(a *model.ThreadCommentAttachment) *ThreadCommentAttachmentResource {
	return &ThreadCommentAttachmentResource{
		Url:          a.EntAttachment.URL,
		DisplayOrder: a.EntAttachment.DisplayOrder,
		Type:         a.TypeToString(),
	}
}

func NewThreadCommentResource(c *model.ThreadComment) *ThreadCommentResource {
	var userId *int
	if c.EntThreadComment.UserID != nil {
		userId = c.EntThreadComment.UserID
	}

	var guestName *string
	if c.EntThreadComment.GuestName != nil {
		guestName = c.EntThreadComment.GuestName
	}

	var parentCommentID int
	if c.EntThreadComment.ParentCommentID != nil {
		parentCommentID = *c.EntThreadComment.ParentCommentID
	}

	var attachments []*ThreadCommentAttachmentResource
	for _, attachment := range c.EntThreadComment.Edges.Attachments {
		attachments = append(attachments, NewThreadAttachmentResource(&model.ThreadCommentAttachment{
			EntAttachment: attachment,
		}))
	}

	return &ThreadCommentResource{
		Id:              c.EntThreadComment.ID,
		UserId:          userId,
		GuestName:       guestName,
		Content:         c.EntThreadComment.Content,
		ParentCommentID: parentCommentID,
		CreatedAt:       c.EntThreadComment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       c.EntThreadComment.UpdatedAt.Format(time.RFC3339),
		Attachments:     attachments,
		TotalReplies:    len(c.EntThreadComment.Edges.Replies),
	}
}

func NewThreadResource(t *model.Thread) *ThreadResource {
	description := ""
	if t.EntThread.Description != "" {
		description = t.EntThread.Description
	}

	thumbnailUrl := t.EntThread.ThumbnailURL

	var tagNames []string
	for _, tag := range t.EntThread.Edges.Tags {
		tagNames = append(tagNames, tag.Name)
	}

	commentCount := len(t.EntThread.Edges.Comments)

	popularity := fmt.Sprintf("%d%%", t.Popularity)

	var boardResource *ThreadBoardResource
	if t.EntThread.Edges.Board != nil {
		boardResource = NewThreadBoardResource(&model.Board{EntBoard: t.EntThread.Edges.Board})
	}

	var comments []*ThreadCommentResource
	for _, comment := range t.EntThread.Edges.Comments {
		comments = append(comments, NewThreadCommentResource(&model.ThreadComment{
			EntThreadComment: comment,
		}))
	}

	return &ThreadResource{
		Id:           t.EntThread.ID,
		Board:        boardResource,
		Title:        t.EntThread.Title,
		Description:  description,
		ThumbnailUrl: thumbnailUrl,
		Tags:         tagNames,
		CreatedAt:    t.EntThread.CreatedAt.Format(time.RFC3339),
		CommentCount: commentCount,
		Popularity:   popularity,
		Comments:     comments,
	}
}
