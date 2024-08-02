package response_admin

import (
	"server/domain/model"
	"time"
)

type TagResponse struct {
	ID          int    `json:"id"`
	ThreadCount int    `json:"threadCount"`
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
}

type NewTagResponseParams struct {
	Tag *model.Tag
}

func NewTagResponse(params NewTagResponseParams) *TagResponse {

	return &TagResponse{
		ID:          params.Tag.EntTag.ID,
		Name:        params.Tag.EntTag.Name,
		ThreadCount: len(params.Tag.EntTag.Edges.Threads),
		CreatedAt:   params.Tag.EntTag.CreatedAt.Format(time.RFC3339),
	}
}
