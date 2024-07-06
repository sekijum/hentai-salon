package request

type ThreadFindAllRequest struct {
	Limit  int `form:"limit" binding:"min=1,max=100"`
	Offset int `form:"offset" binding:"min=0"`
}

type ThreadCreateRequest struct {
	BoardId             int   `json:"boardId" binding:"required"`
	Title             string   `json:"title" binding:"required,max=50"`
	Description       *string  `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl      *string  `json:"thumbnailUrl" binding:"omitempty,max=255"`
	TagNames              []string `json:"tagNames" binding:"omitempty,dive,max=20"`
}
