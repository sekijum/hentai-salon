package request

type ThreadCommentAttachmentFindAllRequest struct {
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	ThreadIDs []int  `form:"threadIds[]" binding:"omitempty,dive,min=1"`
	Filter    string `form:"filter"`
}
