package response_admin

import (
	"server/domain/model"
	"time"
)

type AdResponse struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	IsActive  int    `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type NewAdResponseParams struct {
	Ad *model.Ad
}

func NewAdResponse(params NewAdResponseParams) *AdResponse {
	return &AdResponse{
		ID:        params.Ad.EntAd.ID,
		IsActive:  params.Ad.EntAd.IsActive,
		Content:   params.Ad.EntAd.Content,
		CreatedAt: params.Ad.EntAd.CreatedAt.Format(time.RFC3339),
		UpdatedAt: params.Ad.EntAd.UpdatedAt.Format(time.RFC3339),
	}
}
