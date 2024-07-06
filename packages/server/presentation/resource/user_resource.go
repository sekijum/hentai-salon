package resource

import (
	"server/domain/model"
	"time"
)

type UserResource struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserResource(user *model.User) *UserResource {
	var avatarUrl string
	if user.AvatarUrl != nil {
		avatarUrl = *user.AvatarUrl
	}

	return &UserResource{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		AvatarUrl: avatarUrl,
		Role:      user.Role.String(),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

func NewUserResourceList(users []*model.User) []*UserResource {
	var userResources []*UserResource
	for _, user := range users {
		userResources = append(userResources, NewUserResource(user))
	}
	return userResources
}
