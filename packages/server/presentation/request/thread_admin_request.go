package request

type ThreadAdminFindAllRequest struct {
	Limit     int     `form:"limit"`
	Offset    int     `form:"offset"`
	SortKey   *string `form:"sortKey"`
	SortOrder *string `form:"sortOrder"`
	Keyword   *string `form:"keyword"`
	Status    *int    `form:"status"`
}

type ThreadAdminFindByIDRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type ThreadAdminUpdateRequest struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Status       *int    `json:"status,omitempty"`
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}
