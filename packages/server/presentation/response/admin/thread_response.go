package response_admin

import (
	"server/domain/model"
	"server/infrastructure/ent"
	"time"
)

type ThreadResponse struct {
	ID           int                                 `json:"id"`
	BoardID      int                                 `json:"boardId"`
	UserID       int                                 `json:"userId"`
	Title        string                              `json:"title"`
	Description  *string                             `json:"description,omitempty"`
	ThumbnailURL *string                             `json:"thumbnailUrl,omitempty"`
	IPAddress    string                              `json:"ipAddress"`
	Status       int                                 `json:"status"`
	StatusLabel  string                              `json:"statusLabel"`
	CreatedAt    string                              `json:"createdAt"`
	UpdatedAt    string                              `json:"updatedAt"`
	Board        *BoardResponse                      `json:"board"`
	Comments     *Collection[*ThreadCommentResponse] `json:"comments"`
}

type NewThreadResponseParams struct {
	Thread                        *model.Thread
	Limit, Offset, CommentCount   int
	IncludeComments, IncludeBoard bool
}

func NewThreadResponse(params NewThreadResponseParams) *ThreadResponse {
	var threadCommentResponseList []*ThreadCommentResponse
	var commentCollection *Collection[*ThreadCommentResponse]
	if params.IncludeComments {
		for _, comment_i := range params.Thread.EntThread.Edges.Comments {
			commentResponse := NewThreadCommentResponse(NewThreadCommentResponseParams{
				ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
					EntThreadComment: comment_i,
				}),
			})
			threadCommentResponseList = append(threadCommentResponseList, commentResponse)
		}

		commentCollection = NewCollection(NewCollectionParams[*ThreadCommentResponse]{
			Data:       threadCommentResponseList,
			TotalCount: params.CommentCount,
			Limit:      params.Limit,
			Offset:     params.Offset,
		})
	}

	var threadBoardResponse *BoardResponse
	if params.IncludeBoard {
		if params.Thread.EntThread.Edges.Board != nil {
			threadBoardResponse = NewBoardResponse(NewBoardResponseParams{
				Board: model.NewBoard(model.NewBoardParams{
					EntBoard: &ent.Board{
						ID:    params.Thread.EntThread.BoardID,
						Title: params.Thread.EntThread.Edges.Board.Title,
					},
				})})
		}
	}

	return &ThreadResponse{
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
