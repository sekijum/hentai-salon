package service

import (
	"context"
	"server/infrastructure/datasource"
	resource "server/presentation/resource/board"
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

func (svc *BoardApplicationService) FindAll(ctx context.Context) ([]*resource.BoardResource, error) {
	boards, err := svc.boardDatasource.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var boardResources []*resource.BoardResource
	for _, board := range boards {
		boardResources = append(boardResources, resource.NewBoardResource(board))
	}

	return boardResources, nil
}
