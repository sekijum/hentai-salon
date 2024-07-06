package resource

import (
	"server/domain/model"
)

type BoardResource struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ThumbnailUrl string   `json:"thumbnailUrl"`
}

func NewBoardResource(b *model.Board) *BoardResource {
	description := ""
	if b.Description != nil {
		description = *b.Description
	}

	thumbnailUrl := ""
	if b.ThumbnailUrl != nil {
		thumbnailUrl = *b.ThumbnailUrl
	}

	return &BoardResource{
		Id:          b.Id,
		UserId:      b.UserId,
		Title:       b.Title,
		Description: description,
		ThumbnailUrl: thumbnailUrl,
	}
}
