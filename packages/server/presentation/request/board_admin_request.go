package request

type BoardAdminFindAllRequest struct {
	Limit     int     `form:"limit"`
	Offset    int     `form:"offset"`
	SortOrder *string `form:"sortOrder"`
	SortKey   *string `form:"sortKey"`
	Keyword   *string `form:"keyword"`
	Status    *int    `form:"status"`
}

type BoardAdminUpdateRequest struct {
	Title        string  `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Status       int     `json:"status,omitempty"`
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}
