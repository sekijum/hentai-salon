package resource

type Collection[T any] struct {
	Data       []T `json:"data"`
	TotalCount int `json:"totalCount"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
}

type NewCollectionParams[T any] struct {
	Data       []T
	TotalCount int
	Limit      int
	Offset     int
}

func NewCollection[T any](params NewCollectionParams[T]) *Collection[T] {
	return &Collection[T]{
		Data:       params.Data,
		TotalCount: params.TotalCount,
		Limit:      params.Limit,
		Offset:     params.Offset,
	}
}
