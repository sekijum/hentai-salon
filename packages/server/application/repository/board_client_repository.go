package repository

import (
	"context"
	"server/domain/model"
)

type BoardClientRepository interface {
	Create(ctx context.Context, board *model.Board) (*model.Board, error)
}
