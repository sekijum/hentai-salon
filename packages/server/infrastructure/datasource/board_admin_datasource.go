package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/board"
	"time"
)

type BoardAdminDatasource struct {
	client *ent.Client
}

func NewBoardAdminDatasource(client *ent.Client) *BoardAdminDatasource {
	return &BoardAdminDatasource{client: client}
}

type BoardAdminGetBoardCountParams struct {
	Ctx     context.Context
	Keyword *string
	Status  *int
}

func (ds *BoardAdminDatasource) GetBoardCount(params BoardAdminGetBoardCountParams) (int, error) {
	query := ds.client.Board.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(board.Or(
			board.TitleContains(*params.Keyword),
			board.DescriptionContains(*params.Keyword),
		))
	}

	if params.Status != nil && *params.Status != 0 {
		query = query.Where(board.StatusEQ(*params.Status))
	}

	boardCount, err := query.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return boardCount, nil
}

type BoardAdminFindByIDParams struct {
	Ctx     context.Context
	BoardID int
}

func (ds *BoardAdminDatasource) FindByID(params BoardAdminFindByIDParams) (*model.Board, error) {
	entBoard, err := ds.client.Board.Get(params.Ctx, params.BoardID)
	if err != nil {
		return nil, err
	}
	return &model.Board{EntBoard: entBoard}, nil
}

type BoardAdminFindAllParams struct {
	Ctx       context.Context
	Limit     int
	Offset    int
	SortKey   *string
	SortOrder *string
	Keyword   *string
	Status    *int
}

func (ds *BoardAdminDatasource) FindAll(params BoardAdminFindAllParams) ([]*model.Board, error) {
	query := ds.client.Board.Query()

	sortKey := board.FieldID
	if params.SortKey != nil && *params.SortKey != "" {
		sortKey = *params.SortKey
	}

	if params.SortOrder != nil && *params.SortOrder == "asc" {
		query = query.Order(ent.Asc(sortKey))
	} else {
		query = query.Order(ent.Desc(sortKey))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(board.Or(
			board.TitleContains(*params.Keyword),
			board.DescriptionContains(*params.Keyword),
		))
	}

	if params.Status != nil && *params.Status != 0 {
		query = query.Where(board.StatusEQ(*params.Status))
	}

	query = query.Limit(params.Limit)
	query = query.Offset(params.Offset)

	entBoards, err := query.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelBoards []*model.Board
	for _, entBoard := range entBoards {
		modelBoards = append(modelBoards, &model.Board{
			EntBoard: entBoard,
		})
	}

	return modelBoards, nil
}

type BoardAdminUpdateParams struct {
	Ctx   context.Context
	Board *model.Board
}

func (ds *BoardAdminDatasource) Update(params BoardAdminUpdateParams) (*model.Board, error) {
	board, err := ds.client.Board.Get(params.Ctx, params.Board.EntBoard.ID)
	if err != nil {
		return nil, err
	}

	update := board.Update()

	if params.Board.EntBoard.Title != "" {
		update = update.SetTitle(params.Board.EntBoard.Title)
	}
	if params.Board.EntBoard.Description != nil {
		update = update.SetDescription(*params.Board.EntBoard.Description)
	}
	if params.Board.EntBoard.Status != 0 {
		update = update.SetStatus(params.Board.EntBoard.Status)
	}
	if params.Board.EntBoard.ThumbnailURL != nil {
		update = update.SetThumbnailURL(*params.Board.EntBoard.ThumbnailURL)
	}
	update.SetUpdatedAt(time.Now())

	board, err = update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	return &model.Board{EntBoard: board}, nil
}
