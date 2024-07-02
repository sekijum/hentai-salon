package request

import (
	pagination "server/domain/type"
)

type ThreadFindAllRequest struct {
	Pagination pagination.Pagination `form:"pagination"`
}
