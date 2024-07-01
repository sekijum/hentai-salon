package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/board"

	"github.com/mitchellh/mapstructure"
)

type BoardDatasource struct {
	client *ent.Client
}

func NewBoardDatasource(client *ent.Client) *BoardDatasource {
	return &BoardDatasource{client: client}
}

func (ds *BoardDatasource) Create(
	ctx context.Context,
	b *model.Board,
) (*model.Board, error) {
	tx, err := ds.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	boardBuilder := tx.Board.Create().
		SetUserId(b.UserId).
		SetTitle(b.Title).
		SetStatus(b.Status.ToInt())
	if b.Description != nil {
		boardBuilder.SetDescription(*b.Description)
	}
	if b.ThumbnailUrl != nil {
		boardBuilder.SetThumbnailUrl(*b.ThumbnailUrl)
	}

	savedBoard, err := boardBuilder.Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	modelBoard, err := entBoardToModelBoard(savedBoard)
	if err != nil {
		return nil, err
	}

	return modelBoard, nil
}

func (ds *BoardDatasource) FindAll(ctx context.Context) ([]*model.Board, error) {
	boards, err := ds.client.Board.Query().WithThreads().All(ctx)
	if err != nil {
		return nil, err
	}

	var modelBoards []*model.Board
	for _, entBoard := range boards {
		modelBoard, err := entBoardToModelBoard(entBoard)
		if err != nil {
			return nil, err
		}
		modelBoards = append(modelBoards, modelBoard)
	}

	return modelBoards, nil
}

func (ds *BoardDatasource) FindByTitle(
	ctx context.Context,
	title string,
) ([]*model.Board, error) {
	boards, err := ds.client.Board.Query().Where(board.TitleEQ(title)).All(ctx)
	if err != nil {
		return nil, err
	}

	var modelBoards []*model.Board
	for _, entBoard := range boards {
		modelBoard, err := entBoardToModelBoard(entBoard)
		if err != nil {
			return nil, err
		}
		modelBoards = append(modelBoards, modelBoard)
	}

	return modelBoards, nil
}

func entBoardToModelBoard(entBoard *ent.Board) (*model.Board, error) {
	var modelBoard model.Board
	err := mapstructure.Decode(entBoard, &modelBoard)
	if err != nil {
		return nil, err
	}

	var modelThreads []*model.Thread
	for _, entThread := range entBoard.Edges.Threads {
		var modelThread model.Thread
		err := mapstructure.Decode(entThread, &modelThread)
		if err != nil {
			return nil, err
		}
		modelThreads = append(modelThreads, &modelThread)
	}
	modelBoard.Threads = modelThreads

	return &modelBoard, nil
}
