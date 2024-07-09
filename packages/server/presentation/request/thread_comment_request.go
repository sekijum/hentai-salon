package request

type ThreadCommentFindAllRequest struct {
	ThreadId int `form:"threadId" binding:"required"`
}

type ThreadCommentCreateRequest struct {
	Content   string  `json:"content" binding:"required"`
	GuestName *string `json:"guestName" binding:"omitempty"`
}
