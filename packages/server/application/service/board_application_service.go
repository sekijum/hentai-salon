package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
)

type BoardApplicationService struct {
	boardDatasource *datasource.BoardDatasource
}

func NewBoardApplicationService(
	boardDatasource *datasource.BoardDatasource,
) *BoardApplicationService {
	return &BoardApplicationService{
		boardDatasource: boardDatasource,
	}
}

func (svc *BoardApplicationService) FindAll(ctx context.Context) ([]*model.Board, error) {
	return svc.boardDatasource.FindAll(ctx)
}
