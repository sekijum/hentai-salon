package resource

import (
	"server/domain/model"
	"time"
)

type BoardAdminResource struct {
	ID           int    `json:"id"`
	UserID       int    `json:"userId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Status       int    `json:"status"`
	StatusLabel  string `json:"statusLabel"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type NewBoardAdminResourceParams struct {
	Board *model.Board
}

func NewBoardAdminResource(params NewBoardAdminResourceParams) *BoardAdminResource {
	var thumbnailURL string
	if params.Board.EntBoard.ThumbnailURL != nil {
		thumbnailURL = *params.Board.EntBoard.ThumbnailURL
	}
	var description string
	if params.Board.EntBoard.Description != nil {
		description = *params.Board.EntBoard.Description
	}

	return &BoardAdminResource{
		ID:           params.Board.EntBoard.ID,
		UserID:       params.Board.EntBoard.UserID,
		Title:        params.Board.EntBoard.Title,
		Description:  description,
		ThumbnailURL: thumbnailURL,
		Status:       params.Board.EntBoard.Status,
		StatusLabel:  params.Board.StatusToLabel(),
		CreatedAt:    params.Board.EntBoard.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    params.Board.EntBoard.UpdatedAt.Format(time.RFC3339),
	}
}
