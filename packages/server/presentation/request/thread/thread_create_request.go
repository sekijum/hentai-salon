package request

type ThreadCreateRequest struct {
	BoardId             int   `json:"boardId" binding:"required"`
	Title             string   `json:"title" binding:"required,max=50"`
	Description       *string  `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl      *string  `json:"thumbnailUrl" binding:"omitempty,max=255"`
	TagNameList              []string `json:"tagNameList" binding:"omitempty,dive,max=20"`
}
