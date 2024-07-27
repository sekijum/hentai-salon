package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/board"
)

type BoardDatasource struct {
	client *ent.Client
}

func NewBoardDatasource(client *ent.Client) *BoardDatasource {
	return &BoardDatasource{client: client}
}

type BoardDatasourceFindAllParams struct {
	Ctx context.Context
}

func (ds *BoardDatasource) FindAll(params BoardDatasourceFindAllParams) ([]*model.Board, error) {
	entBoards, err := ds.client.Board.Query().WithThreads().All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var boards []*model.Board
	for _, entBoard_i := range entBoards {
		boards = append(boards, model.NewBoard(model.NewBoardParams{EntBoard: entBoard_i}))
	}

	return boards, nil
}

type BoardDatasourceFindByTitleParams struct {
	Ctx   context.Context
	Title string
}

func (ds *BoardDatasource) FindByTitle(params BoardDatasourceFindByTitleParams) ([]*model.Board, error) {
	entBoards, err := ds.client.Board.Query().Where(board.TitleEQ(params.Title)).WithThreads().All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var boards []*model.Board
	for _, entBoard_i := range entBoards {
		boards = append(boards, model.NewBoard(model.NewBoardParams{EntBoard: entBoard_i}))
	}

	return boards, nil
}

type BoardDatasourceCreateParams struct {
	Ctx   context.Context
	Board *model.Board
}

func (ds *BoardDatasource) Create(params BoardDatasourceCreateParams) (*model.Board, error) {
	q := ds.client.Board.Create().
		SetUserID(params.Board.EntBoard.UserID).
		SetTitle(params.Board.EntBoard.Title).
		SetStatus(params.Board.EntBoard.Status)

	if params.Board.EntBoard.Description != nil {
		q.SetDescription(*params.Board.EntBoard.Description)
	}
	if params.Board.EntBoard.ThumbnailURL != nil {
		q.SetThumbnailURL(*params.Board.EntBoard.ThumbnailURL)
	}

	entBoard, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	contact := model.NewBoard(model.NewBoardParams{EntBoard: entBoard})

	return contact, nil
}
