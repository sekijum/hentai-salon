package request

type BoardCreateRequest struct {
	Title        string  `json:"title" binding:"required,max=50"`
	Description  *string `json:"description" binding:"omitempty,max=255"`
	ThumbnailURL *string `json:"thumbnailUrl" binding:"omitempty,max=255"`
}
