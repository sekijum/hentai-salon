package request

type UserSignupRequest struct {
	Name        string  `json:"name" binding:"required,max=20"`
	Email       string  `json:"email" binding:"required,email,max=254"`
	Password    string  `json:"password" binding:"required,min=6"`
	DisplayName *string `json:"displayName" binding:"omitempty,max=20"`
	AvatarUrl   *string `json:"avatarUrl" binding:"omitempty,url"`
}

type UserSigninRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
