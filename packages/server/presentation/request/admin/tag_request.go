package request_admin

type TagFindAllRequest struct {
	Limit   int     `form:"limit"`
	Offset  int     `form:"offset"`
	Sort    *string `form:"sort"`
	Order   *string `form:"order"`
	Keyword *string `form:"keyword"`
}
