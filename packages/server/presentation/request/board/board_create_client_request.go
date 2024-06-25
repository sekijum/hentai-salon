package request

type BoardCreateClientRequest struct {
	Title        string              `json:"title" binding:"required,max=50"`
	Description  *string             `json:"description" binding:"omitempty,max=255"`
	ThumbnailUrl *string             `json:"thumbnail_url" binding:"omitempty,max=255"`
	Thread       ThreadCreateRequest `json:"default_thread" binding:"required"`
}

type ThreadCreateRequest struct {
	IsNotifyOnComment bool `json:"is_notify_on_comment"`
}
