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
	boards, err := ds.client.Board.Query().WithThreads().All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelBoards []*model.Board
	for _, entBoard := range boards {
		modelBoards = append(modelBoards, &model.Board{EntBoard: entBoard})
	}

	return modelBoards, nil
}

type BoardDatasourceFindByTitleParams struct {
	Ctx   context.Context
	Title string
}

func (ds *BoardDatasource) FindByTitle(params BoardDatasourceFindByTitleParams) ([]*model.Board, error) {
	boards, err := ds.client.Board.Query().Where(board.TitleEQ(params.Title)).WithThreads().All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelBoards []*model.Board
	for _, entBoard := range boards {
		modelBoards = append(modelBoards, &model.Board{EntBoard: entBoard})
	}

	return modelBoards, nil
}

type BoardDatasourceCreateParams struct {
	Ctx   context.Context
	Board *model.Board
}

func (ds *BoardDatasource) Create(params BoardDatasourceCreateParams) (*model.Board, error) {

	boardBuilder := ds.client.Board.Create().
		SetUserID(params.Board.EntBoard.UserID).
		SetTitle(params.Board.EntBoard.Title).
		SetStatus(params.Board.EntBoard.Status)

	if params.Board.EntBoard.Description != nil {
		boardBuilder.SetDescription(*params.Board.EntBoard.Description)
	}
	if params.Board.EntBoard.ThumbnailURL != nil {
		boardBuilder.SetThumbnailURL(*params.Board.EntBoard.ThumbnailURL)
	}

	savedBoard, err := boardBuilder.Save(params.Ctx)

	if err != nil {
		return nil, err
	}

	return &model.Board{EntBoard: savedBoard}, nil
}
