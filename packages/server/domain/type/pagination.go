package pagination

type Pagination struct {
	Limit  int `form:"limit" binding:"required,min=1,max=100"`
	Offset int `form:"offset" binding:"required,min=0"`
}
