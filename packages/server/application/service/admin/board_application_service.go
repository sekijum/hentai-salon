package service_admin

import (
	"context"
	"server/domain/model"
	datasource_admin "server/infrastructure/datasource/admin"
	"server/infrastructure/ent"
	request_admin "server/presentation/request/admin"
	response_admin "server/presentation/response/admin"
)

type BoardApplicationService struct {
	boardDatasource *datasource_admin.BoardDatasource
}

func NewBoardApplicationService(boardDatasource *datasource_admin.BoardDatasource) *BoardApplicationService {
	return &BoardApplicationService{boardDatasource: boardDatasource}
}

type BoardApplicationServiceFindByIDParams struct {
	Ctx     context.Context
	BoardID int
}

func (svc *BoardApplicationService) FindByID(params BoardApplicationServiceFindByIDParams) (*response_admin.BoardResponse, error) {
	board, err := svc.boardDatasource.FindByID(datasource_admin.BoardDatasourceFindByIDParams{
		Ctx:     params.Ctx,
		BoardID: params.BoardID,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewBoardResponse(response_admin.NewBoardResponseParams{
		Board: board,
	})

	return dto, nil
}

type BoardApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request_admin.BoardFindAllRequest
}

func (svc *BoardApplicationService) FindAll(params BoardApplicationServiceFindAllParams) (*response_admin.Collection[*response_admin.BoardResponse], error) {
	boardList, err := svc.boardDatasource.FindAll(datasource_admin.BoardDatasourceFindAllParams{
		Ctx:     params.Ctx,
		Limit:   params.Qs.Limit,
		Offset:  params.Qs.Offset,
		Sort:    params.Qs.Sort,
		Order:   params.Qs.Order,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	boardCount, err := svc.boardDatasource.GetBoardCount(datasource_admin.BoardDatasourceGetBoardCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var boardResponseList []*response_admin.BoardResponse
	for _, board_i := range boardList {
		boardResponseList = append(boardResponseList, response_admin.NewBoardResponse(response_admin.NewBoardResponseParams{
			Board: board_i,
		}))
	}

	dto := response_admin.NewCollection(response_admin.NewCollectionParams[*response_admin.BoardResponse]{
		Data:       boardResponseList,
		TotalCount: boardCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type BoardApplicationServiceUpdateParams struct {
	Ctx     context.Context
	BoardID int
	Body    request_admin.BoardUpdateRequest
}

func (svc *BoardApplicationService) Update(params BoardApplicationServiceUpdateParams) (*response_admin.BoardResponse, error) {
	board := model.NewBoard(model.NewBoardParams{
		EntBoard: &ent.Board{
			ID:          params.BoardID,
			Title:       params.Body.Title,
			Description: params.Body.Description,
		},
		OptionList: []func(*model.Board){
			model.WithBoardStatus(model.BoardStatus(params.Body.Status)),
		},
	})

	board, err := svc.boardDatasource.Update(datasource_admin.BoardDatasourceUpdateParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewBoardResponse(response_admin.NewBoardResponseParams{
		Board: board,
	})

	return dto, nil
}

type BoardApplicationServiceUpdateStatusParams struct {
	Ctx     context.Context
	BoardID int
	Body    request_admin.BoardUpdateStatusRequest
}

func (svc *BoardApplicationService) UpdateStatus(params BoardApplicationServiceUpdateStatusParams) error {
	board := model.NewBoard(model.NewBoardParams{
		EntBoard: &ent.Board{
			ID: params.BoardID,
		},
		OptionList: []func(*model.Board){
			model.WithBoardStatus(model.BoardStatus(params.Body.Status)),
		},
	})

	_, err := svc.boardDatasource.UpdateStatus(datasource_admin.BoardDatasourceUpdateStatusParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return err
	}

	return nil
}
