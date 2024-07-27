package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/resource"
)

type BoardAdminApplicationService struct {
	boardDatasource *datasource.BoardAdminDatasource
}

func NewBoardAdminApplicationService(boardDatasource *datasource.BoardAdminDatasource) *BoardAdminApplicationService {
	return &BoardAdminApplicationService{boardDatasource: boardDatasource}
}

type BoardAdminApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.BoardAdminFindAllRequest
}

func (svc *BoardAdminApplicationService) FindAll(params BoardAdminApplicationServiceFindAllParams) (*resource.Collection[*resource.BoardAdminResource], error) {
	boardList, err := svc.boardDatasource.FindAll(datasource.BoardAdminFindAllParams{
		Ctx:       params.Ctx,
		Limit:     params.Qs.Limit,
		Offset:    params.Qs.Offset,
		SortKey:   params.Qs.SortKey,
		SortOrder: params.Qs.SortOrder,
		Keyword:   params.Qs.Keyword,
		Status:    params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	boardCount, err := svc.boardDatasource.GetBoardCount(datasource.BoardAdminGetBoardCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var boardAdminResourceList []*resource.BoardAdminResource
	for _, board_i := range boardList {
		boardAdminResourceList = append(boardAdminResourceList, resource.NewBoardAdminResource(resource.NewBoardAdminResourceParams{
			Board: board_i,
		}))
	}

	dto := resource.NewCollection(resource.NewCollectionParams[*resource.BoardAdminResource]{
		Data:       boardAdminResourceList,
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

func (svc *BoardAdminApplicationService) Update(params BoardAdminApplicationServiceUpdateParams) (*resource.BoardAdminResource, error) {
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

	board, err := svc.boardDatasource.Update(datasource.BoardAdminUpdateParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewBoardAdminResource(resource.NewBoardAdminResourceParams{
		Board: board,
	})

	return dto, nil
}
