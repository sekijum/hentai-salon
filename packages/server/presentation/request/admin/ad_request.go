package request_admin

type AdFindAllRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Order   *string `form:"order"`
	Sort    *string `form:"sort"`
	Keyword *string `form:"keyword"`
}

type AdCreateRequest struct {
	Content  string `json:"content,omitempty"`
	IsActive int    `json:"isActive,omitempty"`
}

type AdUpdateRequest struct {
	Content  string `json:"content,omitempty"`
	IsActive int    `json:"isActive,omitempty"`
}

type AdUpdateIsActiveRequest struct {
	IsActive int `json:"isActive,omitempty"`
}
