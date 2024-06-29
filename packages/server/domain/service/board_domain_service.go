package service

import (
	"context"
	"errors"
	"server/infrastructure/datasource"
)

type BoardDomainService struct {
	boardDatasource *datasource.BoardDatasource
}

func NewBoardDomainService(boardDatasource *datasource.BoardDatasource) *BoardDomainService {
	return &BoardDomainService{boardDatasource: boardDatasource}
}

func (ds *BoardDomainService) IsTitleDuplicated(ctx context.Context, title string) (bool, error) {
	boards, err := ds.boardDatasource.FindByTitle(ctx, title)
	if err != nil {
		return false, err
	}
	if len(boards) > 0 {
		return true, errors.New("板タイトルが重複しています")
	}
	return false, nil
}
