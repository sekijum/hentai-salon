package request

type UserFindByIdRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type UserSignupRequest struct {
	Name        string  `json:"name" binding:"required,max=20"`
	Email       string  `json:"email" binding:"required,email,max=254"`
	Password    string  `json:"password" binding:"required,min=6"`
	ProfileLink *string `json:"profileLink" binding:"omitempty,url"`
}

type UserSigninRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateRequest struct {
	Name        string  `json:"name"`
	Email       string  `json:"email" binding:"required,max=254"`
	ProfileLink *string `json:"profileLink"`
}

type UserUpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required,min=6"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}

type UserForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email,max=254"`
}

type UserResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserVerifyResetPasswordTokenRequest struct {
	Token string `json:"token" binding:"required"`
}
