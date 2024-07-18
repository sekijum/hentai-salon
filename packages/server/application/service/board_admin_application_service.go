package service

import (
	"context"
	"server/infrastructure/datasource"
	"server/presentation/request"
	"server/presentation/resource"
	"time"
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

func (svc *BoardAdminApplicationService) FindAll(params BoardAdminApplicationServiceFindAllParams) (*resource.ListResource[*resource.BoardAdminResource], error) {
	boards, err := svc.boardDatasource.FindAll(datasource.BoardAdminFindAllParams{
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

	totalCount, err := svc.boardDatasource.GetBoardCount(datasource.BoardAdminGetBoardCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var boardAdminResourceList []*resource.BoardAdminResource
	for _, board := range boards {
		boardAdminResourceList = append(boardAdminResourceList, resource.NewBoardAdminResource(resource.NewBoardAdminResourceParams{
			Board: board,
		}))
	}

	dto := &resource.ListResource[*resource.BoardAdminResource]{
		TotalCount: totalCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
		Data:       boardAdminResourceList,
	}

	return dto, nil
}

type BoardAdminApplicationServiceUpdateParams struct {
	Ctx     context.Context
	BoardID int
	Body    request.BoardAdminUpdateRequest
}

func (svc *BoardAdminApplicationService) Update(params BoardAdminApplicationServiceUpdateParams) (*resource.BoardAdminResource, error) {
	board, err := svc.boardDatasource.FindByID(datasource.BoardAdminFindByIDParams{
		Ctx:     params.Ctx,
		BoardID: params.BoardID,
	})
	if err != nil {
		return nil, err
	}

	if params.Body.Title != nil {
		board.EntBoard.Title = *params.Body.Title
	}
	if params.Body.Description != nil {
		board.EntBoard.Description = params.Body.Description
	}
	if params.Body.Status != nil {
		board.EntBoard.Status = *params.Body.Status
	}
	if params.Body.ThumbnailURL != nil {
		board.EntBoard.ThumbnailURL = params.Body.ThumbnailURL
	}
	board.EntBoard.UpdatedAt = time.Now()

	updatedBoard, err := svc.boardDatasource.Update(datasource.BoardAdminUpdateParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewBoardAdminResource(resource.NewBoardAdminResourceParams{
		Board: updatedBoard,
	})

	return dto, nil
}
