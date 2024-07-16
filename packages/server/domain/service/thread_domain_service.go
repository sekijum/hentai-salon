package service

import (
	"context"
	"errors"
	"server/infrastructure/datasource"
)

type ThreadDomainService struct {
	threadDatasource *datasource.ThreadDatasource
}

func NewThreadDomainService(threadDatasource *datasource.ThreadDatasource) *ThreadDomainService {
	return &ThreadDomainService{threadDatasource: threadDatasource}
}

type ThreadDomainServiceTitleDuplicatedParams struct {
	Ctx   context.Context
	Title string
}

func (ds *ThreadDomainService) IsTitleDuplicated(params ThreadDomainServiceTitleDuplicatedParams) (bool, error) {
	threads, err := ds.threadDatasource.FindByTitle(datasource.ThreadDatasourceFindByTitleParams{
		Ctx:   params.Ctx,
		Title: params.Title,
	})
	if err != nil {
		return false, err
	}
	if len(threads) > 0 {
		return true, errors.New("スレタイが重複しています")
	}
	return false, nil
}
