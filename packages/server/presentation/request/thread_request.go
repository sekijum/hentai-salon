package request

type ThreadRequest struct {
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	Order  string `form:"order"`
}

type ThreadFindByIdRequest struct {
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	Order  string `form:"order"`
}

type ThreadFindAllRequest struct {
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	Filter    string `form:"filter"`
	ThreadIDs []int  `form:"threadIds[]" binding:"omitempty,dive,min=1"`
	Keyword   string `form:"keyword" binding:"omitempty,max=50"`
	BoardID   int    `form:"boardId" binding:"omitempty"`
}

type ThreadCreateRequest struct {
	BoardId      int      `json:"boardId" binding:"required"`
	Title        string   `json:"title" binding:"required,max=255"`
	Description  *string  `json:"description" binding:"omitempty"`
	ThumbnailURL *string  `json:"thumbnailUrl" binding:"omitempty,max=255"`
	TagNameList  []string `json:"tagNameList" binding:"omitempty,dive,max=50"`
}
