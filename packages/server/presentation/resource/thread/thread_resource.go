package resource

import (
	"time"

	"server/domain/model"
)

type ThreadResource struct {
	Id           int    `json:"id"`
	BoardId      int    `json:"boardId"`
	UserId       int    `json:"userId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func NewThreadResource(t *model.Thread) *ThreadResource {
	description := ""
	if t.Description != nil {
		description = *t.Description
	}

	thumbnailUrl := ""
	if t.ThumbnailUrl != nil {
		thumbnailUrl = *t.ThumbnailUrl
	}

	return &ThreadResource{
		Id:           t.Id,
		BoardId:      t.BoardId,
		UserId:       t.UserId,
		Title:        t.Title,
		Description:  description,
		ThumbnailUrl: thumbnailUrl,
		Status:       t.Status.String(),
		CreatedAt:    t.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    t.UpdatedAt.Format(time.RFC3339),
	}
}
