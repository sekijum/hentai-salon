package request

type ThreadFindByIdRequest struct {
	Limit       int `form:"limit"`
	Offset      int `form:"offset"`
	SortByKey   int `form:"sort_by_key"`
	SortByOrder int `form:"sort_by_order"`
}

type ThreadFindAllRequest struct {
	Limit         int      `form:"limit"`
	Offset        int      `form:"offset"`
	QueryCriteria []string `form:"queryCriteria[]"`
	ThreadIds     []int    `form:"threadIds[]" binding:"omitempty,dive,min=1"`
	Keyword       string   `form:"keyword" binding:"omitempty,max=50"`
	BoardId       int      `form:"boardId" binding:"omitempty"`
}

type ThreadCreateRequest struct {
	BoardId      int      `json:"boardId" binding:"required"`
	Title        string   `json:"title" binding:"required,max=50"`
	Description  *string  `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl *string  `json:"thumbnailUrl" binding:"omitempty,max=255"`
	TagNames     []string `json:"tagNames" binding:"omitempty,dive,max=20"`
}
