package resource

import (
	"server/domain/model"
)

type BoardResource struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	ThumbnailURL *string `json:"thumbnailUrl"`
}

type NewBoardResourceParams struct {
	Board *model.Board
}

func NewBoardResource(params NewBoardResourceParams) *BoardResource {
	return &BoardResource{
		Id:           params.Board.EntBoard.ID,
		Title:        params.Board.EntBoard.Title,
		Description:  params.Board.EntBoard.Description,
		ThumbnailURL: params.Board.EntBoard.ThumbnailURL,
	}
}
