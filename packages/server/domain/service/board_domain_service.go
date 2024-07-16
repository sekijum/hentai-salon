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

type BoardDomainServiceIsTitleDuplicatedParams struct {
	Ctx   context.Context
	Title string
}

func (ds *BoardDomainService) IsTitleDuplicated(params BoardDomainServiceIsTitleDuplicatedParams) (bool, error) {
	boards, err := ds.boardDatasource.FindByTitle(datasource.BoardDatasourceFindByTitleParams{
		Ctx:   params.Ctx,
		Title: params.Title,
	})
	if err != nil {
		return false, err
	}
	if len(boards) > 0 {
		return true, errors.New("板タイトルが重複しています")
	}
	return false, nil
}
