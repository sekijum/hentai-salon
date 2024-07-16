package resource

import (
	"server/domain/model"
	"time"
)

type ThreadResource struct {
	ID           int                                  `json:"id"`
	Board        *BoardResource                       `json:"board"`
	Title        string                               `json:"title"`
	Description  *string                              `json:"description"`
	ThumbnailURL *string                              `json:"thumbnailUrl"`
	TagNameList  []string                             `json:"tagNameList"`
	CreatedAt    string                               `json:"createdAt"`
	CommentCount int                                  `json:"commentCount"`
	Comments     ListResource[*ThreadCommentResource] `json:"comments"`
	Attachments  []*ThreadCommentAttachmentResource   `json:"attachments"`
}

type NewThreadResourceParams struct {
	Thread        *model.Thread
	Limit, Offset int
}

func NewThreadResource(params NewThreadResourceParams) *ThreadResource {
	var tagNameList []string
	for _, tag := range params.Thread.EntThread.Edges.Tags {
		tagNameList = append(tagNameList, tag.Name)
	}

	var boardResource *BoardResource
	if params.Thread.EntThread.Edges.Board != nil {
		boardResource = &BoardResource{
			Id:    params.Thread.EntThread.Edges.Board.ID,
			Title: params.Thread.EntThread.Edges.Board.Title,
		}
	}

	var comments []*ThreadCommentResource
	var attachments []*ThreadCommentAttachmentResource
	for i, comment := range params.Thread.EntThread.Edges.Comments {
		commentResource := NewThreadCommentResource(NewThreadCommentResourceParams{
			ThreadComment: &model.ThreadComment{EntThreadComment: comment},
			CommentIDs:    params.Thread.CommentIDs,
			Offset:        params.Offset,
			IDx:           &i,
		})
		comments = append(comments, commentResource)

		for _, attachment := range comment.Edges.Attachments {
			threadCommentAttachment := &model.ThreadCommentAttachment{EntAttachment: attachment}
			attachments = append(attachments, &ThreadCommentAttachmentResource{
				Url:          threadCommentAttachment.EntAttachment.URL,
				DisplayOrder: threadCommentAttachment.EntAttachment.DisplayOrder,
				Type:         threadCommentAttachment.TypeToString(),
				CommentID:    comment.ID,
				IDx:          commentResource.IDx,
			})
		}
	}

	commentsList := ListResource[*ThreadCommentResource]{
		TotalCount: params.Thread.CommentCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
		Data:       comments,
	}

	var description *string
	if params.Thread.EntThread.Description != nil {
		description = params.Thread.EntThread.Description
	}
	var thumbnailURL *string
	if params.Thread.EntThread.ThumbnailURL != nil {
		thumbnailURL = params.Thread.EntThread.ThumbnailURL
	}

	return &ThreadResource{
		ID:           params.Thread.EntThread.ID,
		Board:        boardResource,
		Title:        params.Thread.EntThread.Title,
		Description:  description,
		ThumbnailURL: thumbnailURL,
		TagNameList:  tagNameList,
		CreatedAt:    params.Thread.EntThread.CreatedAt.Format(time.RFC3339),
		CommentCount: params.Thread.CommentCount,
		Comments:     commentsList,
		Attachments:  attachments,
	}
}
