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

type BoardAdminDatasourceGetBoardCountParams struct {
	Ctx     context.Context
	Keyword *string
	Status  *int
}

func (ds *BoardAdminDatasource) GetBoardCount(params BoardAdminDatasourceGetBoardCountParams) (int, error) {
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

type BoardAdminDatasourceFindByIDParams struct {
	Ctx     context.Context
	BoardID int
}

func (ds *BoardAdminDatasource) FindByID(params BoardAdminDatasourceFindByIDParams) (*model.Board, error) {
	entBoard, err := ds.client.Board.Get(params.Ctx, params.BoardID)
	if err != nil {
		return nil, err
	}

	board := model.NewBoard(model.NewBoardParams{EntBoard: entBoard})

	return board, nil
}

type BoardAdminDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
	Status  *int
}

func (ds *BoardAdminDatasource) FindAll(params BoardAdminDatasourceFindAllParams) ([]*model.Board, error) {
	query := ds.client.Board.Query()

	Sort := board.FieldID
	if params.Sort != nil && *params.Sort != "" {
		Sort = *params.Sort
	}

	if params.Order != nil && *params.Order == "asc" {
		query = query.Order(ent.Asc(Sort))
	} else {
		query = query.Order(ent.Desc(Sort))
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
		modelBoards = append(modelBoards, model.NewBoard(model.NewBoardParams{EntBoard: entBoard}))
	}

	return modelBoards, nil
}

type BoardAdminDatasourceUpdateParams struct {
	Ctx   context.Context
	Board *model.Board
}

func (ds *BoardAdminDatasource) Update(params BoardAdminDatasourceUpdateParams) (*model.Board, error) {
	update := ds.client.Board.UpdateOneID(params.Board.EntBoard.ID)

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
	update = update.SetUpdatedAt(time.Now())

	entBoard, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	board := model.NewBoard(model.NewBoardParams{
		EntBoard: entBoard,
	})

	return board, nil
}
