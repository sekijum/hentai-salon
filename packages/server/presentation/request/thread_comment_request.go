package request

type ThreadCommentFindAllRequest struct {
	ThreadId int `form:"threadId" binding:"required"`
}

type ThreadCommentCreateRequest struct {
	ThreadId        int    `json:"threadId" binding:"required"`
	ParentCommentId *int   `json:"parentCommentId" binding:"omitempty"`
	Content         string `json:"content" binding:"required"`
}
