package request_admin

type UserFindAllRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Order   *string `form:"order"`
	Sort    *string `form:"sort"`
	Keyword *string `form:"keyword"`
	Role    *int    `form:"role"`
}

type UserUpdateRequest struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Role   int    `json:"role,omitempty"`
	Status int    `json:"status,omitempty"`
}

type UserUpdateStatusRequest struct {
	Status int `json:"status,omitempty"`
}
