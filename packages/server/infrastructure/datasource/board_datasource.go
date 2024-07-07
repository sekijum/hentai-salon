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

func (ds *BoardDatasource) Create(ctx context.Context, m *model.Board) (*model.Board, error) {
	tx, err := ds.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	boardBuilder := tx.Board.Create().
		SetUserID(m.EntBoard.UserID).
		SetTitle(m.EntBoard.Title).
		SetStatus(m.EntBoard.Status)
	if m.EntBoard.Description != "" {
		boardBuilder.SetDescription(m.EntBoard.Description)
	}
	if m.EntBoard.ThumbnailURL != "" {
		boardBuilder.SetThumbnailURL(m.EntBoard.ThumbnailURL)
	}

	savedBoard, err := boardBuilder.Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.Board{
		EntBoard: savedBoard,
	}, nil
}

func (ds *BoardDatasource) FindAll(ctx context.Context) ([]*model.Board, error) {
	boards, err := ds.client.Board.Query().WithThreads().All(ctx)
	if err != nil {
		return nil, err
	}

	var modelBoards []*model.Board
	for _, entBoard := range boards {
		modelBoards = append(modelBoards, &model.Board{
			EntBoard: entBoard,
		})
	}

	return modelBoards, nil
}

func (ds *BoardDatasource) FindByTitle(
	ctx context.Context,
	title string,
) ([]*model.Board, error) {
	boards, err := ds.client.Board.Query().Where(board.TitleEQ(title)).WithThreads().All(ctx)
	if err != nil {
		return nil, err
	}

	var modelBoards []*model.Board
	for _, entBoard := range boards {
		modelBoards = append(modelBoards, &model.Board{
			EntBoard: entBoard,
		})
	}

	return modelBoards, nil
}
