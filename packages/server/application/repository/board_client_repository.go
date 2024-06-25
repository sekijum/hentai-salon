package repository

import (
	"context"
	"server/domain/model"
)

type BoardClientRepository interface {
	Create(ctx context.Context, board *model.Board, thread *model.Thread) (*model.Board, error)
}
