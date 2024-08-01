package request

type ThreadAdminFindAllRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Sort    *string `form:"sort"`
	Order   *string `form:"order"`
	Keyword *string `form:"keyword"`
}

type ThreadAdminFindByIDRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Sort    *string `form:"sort"`
	Order   *string `form:"order"`
	Keyword *string `form:"keyword"`
}

type ThreadAdminUpdateRequest struct {
	Title        string  `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	Status       int     `json:"status,omitempty"`
}
