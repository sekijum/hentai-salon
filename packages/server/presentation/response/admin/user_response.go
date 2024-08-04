package response_admin

import (
	"server/domain/model"
	"time"
)

type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        int    `json:"role"`
	RoleLabel   string `json:"roleLabel"`
	Status      int    `json:"status"`
	StatusLabel string `json:"statusLabel"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type NewUserResponseParams struct {
	User *model.User
}

func NewUserResponse(params NewUserResponseParams) *UserResponse {
	return &UserResponse{
		ID:          params.User.EntUser.ID,
		Name:        params.User.EntUser.Name,
		Email:       params.User.EntUser.Email,
		Role:        params.User.EntUser.Role,
		RoleLabel:   params.User.RoleToLabel(),
		Status:      params.User.EntUser.Status,
		StatusLabel: params.User.StatusToLabel(),
		CreatedAt:   params.User.EntUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   params.User.EntUser.UpdatedAt.Format(time.RFC3339),
	}
}
