package dto

type BoardCreateClientRequest struct {
	Title                    string `json:"title" binding:"required"`
	Description              string `json:"description"`
	UserId                   int    `json:"user_id" binding:"required"`
	DefaultThreadTitle       string `json:"default_thread_title"`
	DefaultThreadDescription string `json:"default_thread_description"`
}
