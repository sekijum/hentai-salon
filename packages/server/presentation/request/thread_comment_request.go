package request

type ThreadCommentFindAllRequest struct {
	ThreadId int `form:"threadId" binding:"required"`
	Page     int `form:"page" binding:"required,min=1"`
	Limit    int `form:"limit" binding:"required,min=1,max=100"`
}

type ThreadCommentCreateRequest struct {
	ThreadId        int     `json:"threadId" binding:"required"`
	ParentCommentId *int    `json:"parentCommentId" binding:"omitempty"`
	Content         string  `json:"content" binding:"required"`
}
