package service

import (
	"context"
	"errors"
	"server/infrastructure/datasource"
)

type BoardDomainService struct {
	boardClientDatasource *datasource.BoardClientDatasource
}

func NewBoardDomainService(boardClientDatasource *datasource.BoardClientDatasource) *BoardDomainService {
	return &BoardDomainService{boardClientDatasource: boardClientDatasource}
}

func (ds *BoardDomainService) IsTitleDuplicated(ctx context.Context, title string) (bool, error) {
	boards, err := ds.boardClientDatasource.GetByTitle(ctx, title)
	if err != nil {
		return false, err
	}
	if len(boards) > 0 {
		return true, errors.New("タイトルが重複しています")
	}
	return false, nil
}
