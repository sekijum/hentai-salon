package resource

type ListResource[T any] struct {
	TotalCount int `json:"totalCount"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	Data       []T `json:"data"`
}
