package resource

import (
	"server/domain/model"
	"time"
)

type ThreadAdminResource struct {
	ID           int                                        `json:"id"`
	BoardID      int                                        `json:"boardId"`
	UserID       int                                        `json:"userId"`
	Title        string                                     `json:"title"`
	Description  string                                     `json:"description,omitempty"`
	ThumbnailURL string                                     `json:"thumbnailUrl,omitempty"`
	IPAddress    string                                     `json:"ipAddress"`
	Status       int                                        `json:"status"`
	StatusLabel  string                                     `json:"statusLabel"`
	CreatedAt    string                                     `json:"createdAt"`
	UpdatedAt    string                                     `json:"updatedAt"`
	Board        *ThreadBoardAdminResource                  `json:"board"`
	Comments     *ListResource[*ThreadCommentAdminResource] `json:"comments"`
}

type ThreadCommentAdminResource struct {
	ID        int    `json:"id"`
	ThreadID  int    `json:"threadId"`
	UserID    int    `json:"userId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ThreadBoardAdminResource struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type NewThreadAdminResourceParams struct {
	Thread *model.Thread
	Limit  int
	Offset int
}

func NewThreadAdminResource(params NewThreadAdminResourceParams) *ThreadAdminResource {
	var threadCommentAdminResource []*ThreadCommentAdminResource
	for _, comment := range params.Thread.EntThread.Edges.Comments {
		userID := 0
		if comment.UserID != nil {
			userID = *comment.UserID
		}

		commentResource := &ThreadCommentAdminResource{
			ID:        comment.ID,
			ThreadID:  comment.ThreadID,
			UserID:    userID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format(time.RFC3339),
			UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
		}
		threadCommentAdminResource = append(threadCommentAdminResource, commentResource)
	}

	comments := &ListResource[*ThreadCommentAdminResource]{
		TotalCount: len(params.Thread.EntThread.Edges.Comments),
		Limit:      params.Limit,
		Offset:     params.Offset,
		Data:       threadCommentAdminResource,
	}

	var threadBoardAdminResource *ThreadBoardAdminResource
	if params.Thread.EntThread.Edges.Board != nil {
		threadBoardAdminResource = &ThreadBoardAdminResource{
			ID:    params.Thread.EntThread.BoardID,
			Title: params.Thread.EntThread.Edges.Board.Title,
		}
	}

	description := ""
	if params.Thread.EntThread.Description != nil {
		description = *params.Thread.EntThread.Description
	}

	thumbnailURL := ""
	if params.Thread.EntThread.ThumbnailURL != nil {
		thumbnailURL = *params.Thread.EntThread.ThumbnailURL
	}

	return &ThreadAdminResource{
		ID:           params.Thread.EntThread.ID,
		BoardID:      params.Thread.EntThread.BoardID,
		UserID:       params.Thread.EntThread.UserID,
		Title:        params.Thread.EntThread.Title,
		Description:  description,
		ThumbnailURL: thumbnailURL,
		IPAddress:    params.Thread.EntThread.IPAddress,
		Status:       params.Thread.EntThread.Status,
		StatusLabel:  params.Thread.StatusToLabel(),
		CreatedAt:    params.Thread.EntThread.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    params.Thread.EntThread.UpdatedAt.Format(time.RFC3339),
		Comments:     comments,
		Board:        threadBoardAdminResource,
	}
}
