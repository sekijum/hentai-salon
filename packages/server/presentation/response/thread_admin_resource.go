package response

import (
	"server/domain/model"
	"server/infrastructure/ent"
	"time"
)

type ThreadAdminResponse struct {
	ID           int                                      `json:"id"`
	BoardID      int                                      `json:"boardId"`
	UserID       int                                      `json:"userId"`
	Title        string                                   `json:"title"`
	Description  *string                                  `json:"description,omitempty"`
	ThumbnailURL *string                                  `json:"thumbnailUrl,omitempty"`
	IPAddress    string                                   `json:"ipAddress"`
	Status       int                                      `json:"status"`
	StatusLabel  string                                   `json:"statusLabel"`
	CreatedAt    string                                   `json:"createdAt"`
	UpdatedAt    string                                   `json:"updatedAt"`
	Board        *BoardAdminResponse                      `json:"board"`
	Comments     *Collection[*ThreadCommentAdminResponse] `json:"comments"`
}

type ThreadCommentAdminResponse struct {
	ID        int    `json:"id"`
	ThreadID  int    `json:"threadId"`
	UserID    *int   `json:"userId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type NewThreadAdminResponseParams struct {
	Thread *model.Thread
	Limit  *int
	Offset *int
}

func NewThreadAdminResponse(params NewThreadAdminResponseParams) *ThreadAdminResponse {
	limit := 0
	if params.Limit != nil {
		limit = *params.Limit
	}

	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	var threadCommentResponseList []*ThreadCommentAdminResponse
	for _, comment_i := range params.Thread.EntThread.Edges.Comments {
		commentResponse := &ThreadCommentAdminResponse{
			ID:        comment_i.ID,
			ThreadID:  comment_i.ThreadID,
			UserID:    comment_i.UserID,
			Content:   comment_i.Content,
			CreatedAt: comment_i.CreatedAt.Format(time.RFC3339),
			UpdatedAt: comment_i.UpdatedAt.Format(time.RFC3339),
		}
		threadCommentResponseList = append(threadCommentResponseList, commentResponse)
	}

	commentCollection := NewCollection(NewCollectionParams[*ThreadCommentAdminResponse]{
		Data:       threadCommentResponseList,
		TotalCount: len(params.Thread.EntThread.Edges.Comments),
		Limit:      limit,
		Offset:     offset,
	})

	var threadBoardResponse *BoardAdminResponse
	if params.Thread.EntThread.Edges.Board != nil {
		threadBoardResponse = NewBoardAdminResponse(NewBoardAdminResponseParams{
			Board: model.NewBoard(model.NewBoardParams{
				EntBoard: &ent.Board{
					ID:    params.Thread.EntThread.BoardID,
					Title: params.Thread.EntThread.Edges.Board.Title,
				},
			})})
	}

	return &ThreadAdminResponse{
		ID:           params.Thread.EntThread.ID,
		BoardID:      params.Thread.EntThread.BoardID,
		UserID:       params.Thread.EntThread.UserID,
		Title:        params.Thread.EntThread.Title,
		Description:  params.Thread.EntThread.Description,
		ThumbnailURL: params.Thread.EntThread.ThumbnailURL,
		IPAddress:    params.Thread.EntThread.IPAddress,
		Status:       params.Thread.EntThread.Status,
		StatusLabel:  params.Thread.StatusToLabel(),
		CreatedAt:    params.Thread.EntThread.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    params.Thread.EntThread.UpdatedAt.Format(time.RFC3339),
		Comments:     commentCollection,
		Board:        threadBoardResponse,
	}
}
