package request


type ThreadFindAllRequest struct {
	Limit     int      `form:"limit"`
	Offset    int      `form:"offset"`
	Orders    []string `form:"orders[]"`
	ThreadIds []int    `form:"threadIds" binding:"omitempty,dive,min=1"`
}

type ThreadCreateRequest struct {
	BoardId      int      `json:"boardId" binding:"required"`
	Title        string   `json:"title" binding:"required,max=50"`
	Description  *string  `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl *string  `json:"thumbnailUrl" binding:"omitempty,max=255"`
	TagNames     []string `json:"tagNames" binding:"omitempty,dive,max=20"`
}
