package resource

import (
	"server/domain/model"
	"time"
)

type UserResource struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	ProfileLink *string `json:"profileLink,omitempty"`
	Role        string  `json:"role"`
	RoleLabel   string  `json:"roleLabel"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type NewUserResourceParams struct {
	User          *model.User
	Limit, Offset int
}

func NewUserResource(params NewUserResourceParams) *UserResource {
	return &UserResource{
		ID:          params.User.EntUser.ID,
		Name:        params.User.EntUser.Name,
		Email:       params.User.EntUser.Email,
		ProfileLink: params.User.EntUser.ProfileLink,
		Role:        params.User.RoleToString(),
		RoleLabel:   params.User.RoleToLabel(),
		CreatedAt:   params.User.EntUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   params.User.EntUser.UpdatedAt.Format(time.RFC3339),
	}
}
