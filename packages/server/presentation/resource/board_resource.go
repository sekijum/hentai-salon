package resource

import (
	"server/domain/model"
)

type BoardResource struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type NewBoardResourceParams struct {
	Board *model.Board
}

func NewBoardResource(params NewBoardResourceParams) *BoardResource {
	description := ""
	if params.Board.EntBoard.Description != nil {
		description = *params.Board.EntBoard.Description
	}
	thumbnailURL := ""
	if params.Board.EntBoard.ThumbnailURL != nil {
		thumbnailURL = *params.Board.EntBoard.ThumbnailURL
	}

	return &BoardResource{
		Id:           params.Board.EntBoard.ID,
		Title:        params.Board.EntBoard.Title,
		Description:  description,
		ThumbnailURL: thumbnailURL,
	}
}
