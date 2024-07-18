package request

type ThreadCommentFindAllByUserIDRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type ThreadCommentFindByIDRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type ThreadCommentAttachmentRequest struct {
	URL          string `json:"url" binding:"required"`
	DisplayOrder int    `json:"displayOrder" binding:"omitempty"`
	Type         string `json:"type" binding:"required"`
}

type ThreadCommentCreateRequest struct {
	Content     string                           `json:"content" binding:"required"`
	GuestName   *string                          `json:"guestName" binding:"omitempty"`
	Attachments []ThreadCommentAttachmentRequest `json:"attachments" binding:"omitempty,dive"`
}
