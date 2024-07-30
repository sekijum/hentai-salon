package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/response"
)

type BoardAdminApplicationService struct {
	boardAdminDatasource *datasource.BoardAdminDatasource
}

func NewBoardAdminApplicationService(boardAdminDatasource *datasource.BoardAdminDatasource) *BoardAdminApplicationService {
	return &BoardAdminApplicationService{boardAdminDatasource: boardAdminDatasource}
}

type BoardAdminApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.BoardAdminFindAllRequest
}

func (svc *BoardAdminApplicationService) FindAll(params BoardAdminApplicationServiceFindAllParams) (*response.Collection[*response.BoardAdminResponse], error) {
	boardList, err := svc.boardAdminDatasource.FindAll(datasource.BoardAdminDatasourceFindAllParams{
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

	boardCount, err := svc.boardAdminDatasource.GetBoardCount(datasource.BoardAdminDatasourceGetBoardCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var boardAdminResponseList []*response.BoardAdminResponse
	for _, board_i := range boardList {
		boardAdminResponseList = append(boardAdminResponseList, response.NewBoardAdminResponse(response.NewBoardAdminResponseParams{
			Board: board_i,
		}))
	}

	dto := response.NewCollection(response.NewCollectionParams[*response.BoardAdminResponse]{
		Data:       boardAdminResponseList,
		TotalCount: boardCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type BoardAdminApplicationServiceUpdateParams struct {
	Ctx     context.Context
	BoardID int
	Body    request.BoardAdminUpdateRequest
}

func (svc *BoardAdminApplicationService) Update(params BoardAdminApplicationServiceUpdateParams) (*response.BoardAdminResponse, error) {
	board := model.NewBoard(model.NewBoardParams{
		EntBoard: &ent.Board{
			ID:           params.BoardID,
			Title:        params.Body.Title,
			Description:  params.Body.Description,
			ThumbnailURL: params.Body.ThumbnailURL,
		},
		OptionList: []func(*model.Board){
			model.WithBoardStatus(model.BoardStatus(params.Body.Status)),
		},
	})

	board, err := svc.boardAdminDatasource.Update(datasource.BoardAdminDatasourceUpdateParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewBoardAdminResponse(response.NewBoardAdminResponseParams{
		Board: board,
	})

	return dto, nil
}
