package request

type UserAdminFindAllRequest struct {
	Limit     int     `form:"limit"`
	Offset    int     `form:"offset"`
	SortOrder *string `form:"sortOrder"`
	SortKey   *string `form:"sortKey"`
	Keyword   *string `form:"keyword"`
	Role      *int    `form:"role"`
}

type UserAdminUpdateRequest struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Role   int    `json:"role,omitempty"`
	Status int    `json:"status,omitempty"`
}
