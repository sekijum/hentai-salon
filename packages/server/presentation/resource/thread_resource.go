package resource

import (
	"server/domain/model"
	"time"
)

type ThreadBoardResource struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ThreadCommentAttachmentResourceForComment struct {
	Url          string `json:"url"`
	DisplayOrder int    `json:"displayOrder"`
	Type         string `json:"type"`
}

type ThreadCommentAttachmentResourceForThread struct {
	Url          string `json:"url"`
	DisplayOrder int    `json:"displayOrder"`
	Type         string `json:"type"`
	CommentID    int    `json:"commentId"`
	Idx          int    `json:"idx"`
}

type ThreadCommentUserResource struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ProfileLink *string `json:"profileLink"`
}

type ThreadCommentResource struct {
	Id               int                                          `json:"id"`
	Idx              int                                          `json:"idx,omitempty"`
	User             *ThreadCommentUserResource                   `json:"user,omitempty"`
	GuestName        *string                                      `json:"guestName,omitempty"`
	Content          string                                       `json:"content"`
	ParentCommentIdx int                                          `json:"parentCommentIdx,omitempty"`
	ParentCommentID  int                                          `json:"parentCommentId"`
	CreatedAt        string                                       `json:"createdAt"`
	UpdatedAt        string                                       `json:"updatedAt"`
	Attachments      []*ThreadCommentAttachmentResourceForComment `json:"attachments"`
	TotalReplies     int                                          `json:"totalReplies"`
}

type ThreadResource struct {
	Id           int                                         `json:"id"`
	Board        *ThreadBoardResource                        `json:"board"`
	Title        string                                      `json:"title"`
	Description  string                                      `json:"description"`
	ThumbnailUrl string                                      `json:"thumbnailUrl"`
	Tags         []string                                    `json:"tags"`
	CreatedAt    string                                      `json:"createdAt"`
	CommentCount int                                         `json:"commentCount"`
	Comments     ListResource[*ThreadCommentResource]        `json:"comments"`
	Attachments  []*ThreadCommentAttachmentResourceForThread `json:"attachments"`
}

func NewThreadBoardResource(b *model.Board) *ThreadBoardResource {
	return &ThreadBoardResource{
		Id:    b.EntBoard.ID,
		Title: b.EntBoard.Title,
	}
}

func NewThreadAttachmentResourceForComment(a *model.ThreadCommentAttachment) *ThreadCommentAttachmentResourceForComment {
	return &ThreadCommentAttachmentResourceForComment{
		Url:          a.EntAttachment.URL,
		DisplayOrder: a.EntAttachment.DisplayOrder,
		Type:         a.TypeToString(),
	}
}

func NewThreadAttachmentResourceForThread(a *model.ThreadCommentAttachment, commentId, idx int) *ThreadCommentAttachmentResourceForThread {
	return &ThreadCommentAttachmentResourceForThread{
		Url:          a.EntAttachment.URL,
		DisplayOrder: a.EntAttachment.DisplayOrder,
		Type:         a.TypeToString(),
		CommentID:    commentId,
		Idx:          idx,
	}
}

func NewThreadCommentUserResource(c *model.ThreadComment) *ThreadCommentUserResource {
	return &ThreadCommentUserResource{
		Id:          c.EntThreadComment.Edges.Author.ID,
		Name:        c.EntThreadComment.Edges.Author.Name,
		ProfileLink: c.EntThreadComment.Edges.Author.ProfileLink,
	}
}

func NewThreadCommentResource(c *model.ThreadComment, commentIDs []int, offset int) *ThreadCommentResource {
	var user *ThreadCommentUserResource
	if c.EntThreadComment.Edges.Author != nil {
		user = NewThreadCommentUserResource(c)
	}

	var guestName *string
	if c.EntThreadComment.GuestName != nil {
		guestName = c.EntThreadComment.GuestName
	}

	var parentCommentID int
	var parentCommentIdx int
	if c.EntThreadComment.ParentCommentID != nil {
		parentCommentID = *c.EntThreadComment.ParentCommentID
		if len(commentIDs) > 0 {
			parentCommentIdx = model.FindCommentIndexByID(commentIDs, parentCommentID) + 1
		} else {
			parentCommentIdx = 0
		}
	}

	idx := offset + 1

	var attachments []*ThreadCommentAttachmentResourceForComment
	for _, attachment := range c.EntThreadComment.Edges.Attachments {
		attachments = append(attachments, NewThreadAttachmentResourceForComment(&model.ThreadCommentAttachment{
			EntAttachment: attachment,
		}))
	}

	return &ThreadCommentResource{
		Id:               c.EntThreadComment.ID,
		Idx:              idx,
		User:             user,
		GuestName:        guestName,
		Content:          c.EntThreadComment.Content,
		ParentCommentID:  parentCommentID,
		ParentCommentIdx: parentCommentIdx,
		CreatedAt:        c.EntThreadComment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        c.EntThreadComment.UpdatedAt.Format(time.RFC3339),
		Attachments:      attachments,
		TotalReplies:     len(c.EntThreadComment.Edges.Replies),
	}
}

func NewThreadResource(t *model.Thread, limit, offset int) *ThreadResource {
	description := ""
	if t.EntThread.Description != "" {
		description = t.EntThread.Description
	}

	thumbnailUrl := t.EntThread.ThumbnailURL

	var tagNames []string
	for _, tag := range t.EntThread.Edges.Tags {
		tagNames = append(tagNames, tag.Name)
	}

	var boardResource *ThreadBoardResource
	if t.EntThread.Edges.Board != nil {
		boardResource = NewThreadBoardResource(&model.Board{EntBoard: t.EntThread.Edges.Board})
	}

	// コメントリストを作成
	var comments []*ThreadCommentResource
	var attachments []*ThreadCommentAttachmentResourceForThread
	for i, comment := range t.EntThread.Edges.Comments {
		commentResource := NewThreadCommentResource(&model.ThreadComment{
			EntThreadComment: comment,
		}, t.CommentIDs, offset+i)
		comments = append(comments, commentResource)

		for _, attachment := range comment.Edges.Attachments {
			attachments = append(attachments, NewThreadAttachmentResourceForThread(&model.ThreadCommentAttachment{
				EntAttachment: attachment,
			}, comment.ID, commentResource.Idx)) // commentIdとidxを渡す
		}
	}

	commentsList := ListResource[*ThreadCommentResource]{
		TotalCount: t.CommentCount,
		Limit:      limit,
		Offset:     offset,
		Data:       comments,
	}

	return &ThreadResource{
		Id:           t.EntThread.ID,
		Board:        boardResource,
		Title:        t.EntThread.Title,
		Description:  description,
		ThumbnailUrl: thumbnailUrl,
		Tags:         tagNames,
		CreatedAt:    t.EntThread.CreatedAt.Format(time.RFC3339),
		CommentCount: t.CommentCount,
		Comments:     commentsList,
		Attachments:  attachments,
	}
}
