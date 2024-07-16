package service

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	request "server/presentation/request"
	resource "server/presentation/resource"
	"time"
)

type BoardApplicationService struct {
	boardDatasource *datasource.BoardDatasource
}

func NewBoardApplicationService(boardDatasource *datasource.BoardDatasource) *BoardApplicationService {
	return &BoardApplicationService{boardDatasource: boardDatasource}
}

type BoardApplicationServiceFindAllParams struct {
	Ctx context.Context
}

func (svc *BoardApplicationService) FindAll(params BoardApplicationServiceFindAllParams) ([]*resource.BoardResource, error) {
	boards, err := svc.boardDatasource.FindAll(datasource.BoardDatasourceFindAllParams{
		Ctx: params.Ctx,
	})
	if err != nil {
		return nil, err
	}

	var dto []*resource.BoardResource
	for _, board := range boards {
		dto = append(dto, resource.NewBoardResource(resource.NewBoardResourceParams{
			Board: board,
		}))
	}

	return dto, nil
}

type BoardApplicationServiceCreateParams struct {
	Ctx    context.Context
	UserID int
	Body   request.BoardCreateRequest
}

func (svc *BoardApplicationService) Create(params BoardApplicationServiceCreateParams) (*resource.BoardResource, error) {

	boards, err := svc.boardDatasource.FindByTitle(datasource.BoardDatasourceFindByTitleParams{
		Ctx:   params.Ctx,
		Title: params.Body.Title,
	})
	if err != nil {
		return nil, err
	}
	if len(boards) > 0 {
		return nil, errors.New("板タイトルが重複しています")
	}

	board := &model.Board{
		EntBoard: &ent.Board{
			Title:     params.Body.Title,
			UserID:    params.UserID,
			Status:    int(model.BoardStatusPublic),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if params.Body.Description != nil {
		board.EntBoard.Description = params.Body.Description
	}
	if params.Body.ThumbnailURL != nil {
		board.EntBoard.ThumbnailURL = params.Body.ThumbnailURL
	}

	board, err = svc.boardDatasource.Create(datasource.BoardDatasourceCreateParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewBoardResource(resource.NewBoardResourceParams{
		Board: board,
	})

	return dto, nil
}
