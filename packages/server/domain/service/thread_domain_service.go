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

func (ds *ThreadDomainService) IsTitleDuplicated(ctx context.Context, title string) (bool, error) {
	threads, err := ds.threadDatasource.FindByTitle(ctx, title)
	if err != nil {
		return false, err
	}
	if len(threads) > 0 {
		return true, errors.New("スレタイが重複しています")
	}
	return false, nil
}
