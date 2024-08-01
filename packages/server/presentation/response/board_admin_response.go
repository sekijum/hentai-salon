package response

import (
	"server/domain/model"
	"time"
)

type BoardAdminResponse struct {
	ID           int     `json:"id"`
	UserID       int     `json:"userId"`
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	ThumbnailURL *string `json:"thumbnailUrl"`
	Status       int     `json:"status"`
	StatusLabel  string  `json:"statusLabel"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

type NewBoardAdminResponseParams struct {
	Board *model.Board
}

func NewBoardAdminResponse(params NewBoardAdminResponseParams) *BoardAdminResponse {
	return &BoardAdminResponse{
		ID:           params.Board.EntBoard.ID,
		UserID:       params.Board.EntBoard.UserID,
		Title:        params.Board.EntBoard.Title,
		Description:  params.Board.EntBoard.Description,
		ThumbnailURL: params.Board.EntBoard.ThumbnailURL,
		Status:       params.Board.EntBoard.Status,
		StatusLabel:  params.Board.StatusToLabel(),
		CreatedAt:    params.Board.EntBoard.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    params.Board.EntBoard.UpdatedAt.Format(time.RFC3339),
	}
}
