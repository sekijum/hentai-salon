package request_admin

type ContactFindAllRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Order   *string `form:"order"`
	Sort    *string `form:"sort"`
	Keyword *string `form:"keyword"`
	Status  *int    `form:"status"`
}

type ContactUpdateStatusRequest struct {
	Status int `json:"status,omitempty"`
}
