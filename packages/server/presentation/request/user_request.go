package request

type UserSignupRequest struct {
	Name        string  `json:"name" binding:"required,max=20"`
	Email       string  `json:"email" binding:"required,email,max=254"`
	Password    string  `json:"password" binding:"required,min=6"`
	AvatarUrl   *string `json:"avatarUrl" binding:"omitempty,url"`
	ProfileLink *string `json:"profileLink" binding:"omitempty,url"`
}

type UserSigninRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
