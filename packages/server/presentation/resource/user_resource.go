package resource

import (
	"server/domain/model"
	"time"
)

type UserResource struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	ProfileLink *string `json:"profileLink,omitempty"`
	Role        string  `json:"role"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type NewUserResourceParams struct {
	User *model.User
}

func NewUserResource(params NewUserResourceParams) *UserResource {
	var avatarURL *string
	if params.User.EntUser.AvatarURL != nil {
		avatarURL = params.User.EntUser.AvatarURL
	}
	var profileLink *string
	if params.User.EntUser.ProfileLink != nil {
		profileLink = params.User.EntUser.ProfileLink
	}

	return &UserResource{
		ID:          params.User.EntUser.ID,
		Name:        params.User.EntUser.Name,
		Email:       params.User.EntUser.Email,
		AvatarURL:   avatarURL,
		ProfileLink: profileLink,
		Role:        params.User.RoleToString(),
		CreatedAt:   params.User.EntUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   params.User.EntUser.UpdatedAt.Format(time.RFC3339),
	}
}
