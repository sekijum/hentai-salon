package request

type BoardCreateClientRequest struct {
	Title         string        `json:"title" binding:"required,max=50"`
	Description   *string       `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl  *string       `json:"thumbnailUrl" binding:"omitempty,max=255"`
	DefaultThread DefaultThread `json:"defaultThread" binding:"required"`
}

type DefaultThread struct {
	Title             string  `json:"title" binding:"required,max=50"`
	Description       *string `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl      *string `json:"thumbnailUrl" binding:"omitempty,max=255"`
	IsNotifyOnComment bool    `json:"isNotifyOnComment"`
}
