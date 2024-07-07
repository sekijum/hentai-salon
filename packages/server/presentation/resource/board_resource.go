package resource

import (
	"server/domain/model"
)

type BoardResource struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

func NewBoardResource(b *model.Board) *BoardResource {
	description := ""
	if b.EntBoard.Description != "" {
		description = b.EntBoard.Description
	}

	thumbnailUrl := b.EntBoard.ThumbnailURL

	return &BoardResource{
		Id:           b.EntBoard.ID,
		Title:        b.EntBoard.Title,
		Description:  description,
		ThumbnailUrl: thumbnailUrl,
	}
}
