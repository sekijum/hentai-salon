package response_admin

import (
	"server/domain/model"
	"time"
)

type BoardResponse struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userId"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      int     `json:"status"`
	StatusLabel string  `json:"statusLabel"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type NewBoardResponseParams struct {
	Board *model.Board
}

func NewBoardResponse(params NewBoardResponseParams) *BoardResponse {
	return &BoardResponse{
		ID:          params.Board.EntBoard.ID,
		UserID:      params.Board.EntBoard.UserID,
		Title:       params.Board.EntBoard.Title,
		Description: params.Board.EntBoard.Description,
		Status:      params.Board.EntBoard.Status,
		StatusLabel: params.Board.StatusToLabel(),
		CreatedAt:   params.Board.EntBoard.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   params.Board.EntBoard.UpdatedAt.Format(time.RFC3339),
	}
}
