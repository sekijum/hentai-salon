package service

import (
	"context"
	"errors"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/response"
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

func (svc *BoardApplicationService) FindAll(params BoardApplicationServiceFindAllParams) ([]*response.BoardResponse, error) {
	boardList, err := svc.boardDatasource.FindAll(datasource.BoardDatasourceFindAllParams{
		Ctx: params.Ctx,
	})
	if err != nil {
		return nil, err
	}

	var dto []*response.BoardResponse
	for _, board_i := range boardList {
		dto = append(dto, response.NewBoardResponse(response.NewBoardResponseParams{
			Board: board_i,
		}))
	}

	return dto, nil
}

type BoardApplicationServiceCreateParams struct {
	Ctx    context.Context
	UserID int
	Body   request.BoardCreateRequest
}

func (svc *BoardApplicationService) Create(params BoardApplicationServiceCreateParams) (*response.BoardResponse, error) {
	boardList, err := svc.boardDatasource.FindByTitle(datasource.BoardDatasourceFindByTitleParams{
		Ctx:   params.Ctx,
		Title: params.Body.Title,
	})
	if err != nil {
		return nil, err
	}
	if len(boardList) > 0 {
		return nil, errors.New("板タイトルが重複しています")
	}

	board := model.NewBoard(model.NewBoardParams{
		EntBoard: &ent.Board{
			Title:        params.Body.Title,
			UserID:       params.UserID,
			Description:  params.Body.Description,
			ThumbnailURL: params.Body.ThumbnailURL,
		},
		OptionList: []func(*model.Board){
			model.WithBoardStatus(model.BoardStatusPublic),
		},
	})

	board, err = svc.boardDatasource.Create(datasource.BoardDatasourceCreateParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewBoardResponse(response.NewBoardResponseParams{
		Board: board,
	})

	return dto, nil
}
