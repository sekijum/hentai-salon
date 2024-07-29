package response

import (
	"server/domain/model"
	"time"
)

type UserAdminResponse struct {
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

type NewUserAdminResponseParams struct {
	User *model.User
}

func NewUserAdminResponse(params NewUserAdminResponseParams) *UserAdminResponse {
	return &UserAdminResponse{
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
