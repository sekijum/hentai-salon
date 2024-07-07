package resource

import (
	"server/domain/model"
	"time"
)

type UserResource struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewUserResource(u *model.User) *UserResource {
	var avatarUrl string
	if u.EntUser.AvatarURL != nil {
		avatarUrl = *u.EntUser.AvatarURL
	}

	return &UserResource{
		Id:        u.EntUser.ID,
		Name:      u.EntUser.Name,
		Email:     u.EntUser.Email,
		AvatarUrl: avatarUrl,
		Role:      u.RoleToString(),
		CreatedAt: u.EntUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.EntUser.UpdatedAt.Format(time.RFC3339),
	}
}
