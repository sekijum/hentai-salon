package request_admin

type ThreadFindAllRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Sort    *string `form:"sort"`
	Order   *string `form:"order"`
	Keyword *string `form:"keyword"`
}

type ThreadFindByIDRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Sort    *string `form:"sort"`
	Order   *string `form:"order"`
	Keyword *string `form:"keyword"`
}

type ThreadUpdateRequest struct {
	Title        string  `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
	Status       int     `json:"status,omitempty"`
}
