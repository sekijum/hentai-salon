package request

type TagFindAllRequest struct {
	Keyword *string `form:"keyword" binding:"omitempty,max=50"`
}
