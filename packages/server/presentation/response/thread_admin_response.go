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

type NewThreadAdminResponseParams struct {
	Thread                        *model.Thread
	Limit, Offset, CommentCount   int
	IncludeComments, IncludeBoard bool
}

func NewThreadAdminResponse(params NewThreadAdminResponseParams) *ThreadAdminResponse {
	var threadCommentResponseList []*ThreadCommentAdminResponse
	var commentCollection *Collection[*ThreadCommentAdminResponse]
	if params.IncludeComments {
		for _, comment_i := range params.Thread.EntThread.Edges.Comments {
			commentResponse := NewThreadCommentAdminResponse(NewThreadCommentAdminResponseParams{
				ThreadComment: model.NewThreadComment(model.NewThreadCommentParams{
					EntThreadComment: comment_i,
				}),
			})
			threadCommentResponseList = append(threadCommentResponseList, commentResponse)
		}

		commentCollection = NewCollection(NewCollectionParams[*ThreadCommentAdminResponse]{
			Data:       threadCommentResponseList,
			TotalCount: params.CommentCount,
			Limit:      params.Limit,
			Offset:     params.Offset,
		})
	}

	var threadBoardResponse *BoardAdminResponse
	if params.IncludeBoard {
		if params.Thread.EntThread.Edges.Board != nil {
			threadBoardResponse = NewBoardAdminResponse(NewBoardAdminResponseParams{
				Board: model.NewBoard(model.NewBoardParams{
					EntBoard: &ent.Board{
						ID:    params.Thread.EntThread.BoardID,
						Title: params.Thread.EntThread.Edges.Board.Title,
					},
				})})
		}
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
