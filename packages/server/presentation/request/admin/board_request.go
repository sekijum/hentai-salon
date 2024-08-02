package request_admin

type BoardFindAllRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Order   *string `form:"order"`
	Sort    *string `form:"sort"`
	Keyword *string `form:"keyword"`
	Status  *int    `form:"status"`
}

type BoardUpdateRequest struct {
	Title        string  `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Status       int     `json:"status,omitempty"`
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}

type BoardUpdateStatusRequest struct {
	Status int `json:"status,omitempty"`
}
