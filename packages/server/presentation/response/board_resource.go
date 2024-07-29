package response

import (
	"server/domain/model"
)

type BoardResponse struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Description  *string `json:"description,omitempty"`
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}

type NewBoardResponseParams struct {
	Board *model.Board
}

func NewBoardResponse(params NewBoardResponseParams) *BoardResponse {
	return &BoardResponse{
		Id:           params.Board.EntBoard.ID,
		Title:        params.Board.EntBoard.Title,
		Description:  params.Board.EntBoard.Description,
		ThumbnailURL: params.Board.EntBoard.ThumbnailURL,
	}
}
