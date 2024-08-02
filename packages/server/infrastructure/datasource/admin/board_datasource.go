package datasource_admin

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/board"
	"strconv"
	"time"
)

type BoardDatasource struct {
	client *ent.Client
}

func NewBoardDatasource(client *ent.Client) *BoardDatasource {
	return &BoardDatasource{client: client}
}

type BoardDatasourceGetBoardCountParams struct {
	Ctx     context.Context
	Keyword *string
	Status  *int
}

func (ds *BoardDatasource) GetBoardCount(params BoardDatasourceGetBoardCountParams) (int, error) {
	query := ds.client.Board.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				query = query.Where(board.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				query = query.Where(board.IDEQ(id))
			}
		default:
			query = query.Where(board.Or(
				board.TitleContainsFold(*params.Keyword),
			))
		}
	}

	boardCount, err := query.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return boardCount, nil
}

type BoardDatasourceFindByIDParams struct {
	Ctx     context.Context
	BoardID int
}

func (ds *BoardDatasource) FindByID(params BoardDatasourceFindByIDParams) (*model.Board, error) {
	entBoard, err := ds.client.Board.Get(params.Ctx, params.BoardID)
	if err != nil {
		return nil, err
	}

	board := model.NewBoard(model.NewBoardParams{EntBoard: entBoard})

	return board, nil
}

type BoardDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
	Status  *int
}

func (ds *BoardDatasource) FindAll(params BoardDatasourceFindAllParams) ([]*model.Board, error) {
	query := ds.client.Board.Query()

	sort := board.FieldID
	order := "desc"

	if params.Sort != nil && *params.Sort != "" {
		sort = *params.Sort
	}
	if params.Order != nil && *params.Order != "" {
		order = *params.Order
	}

	if order == "asc" {
		query = query.Order(ent.Asc(sort))
	} else {
		query = query.Order(ent.Desc(sort))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				query = query.Where(board.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				query = query.Where(board.IDEQ(id))
			}
		default:
			query = query.Where(board.Or(
				board.TitleContainsFold(*params.Keyword),
			))
		}
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

type BoardDatasourceUpdateParams struct {
	Ctx   context.Context
	Board *model.Board
}

func (ds *BoardDatasource) Update(params BoardDatasourceUpdateParams) (*model.Board, error) {
	update := ds.client.Board.UpdateOneID(params.Board.EntBoard.ID)

	update = update.SetTitle(params.Board.EntBoard.Title).
		SetDescription(*params.Board.EntBoard.Description).
		SetThumbnailURL(*params.Board.EntBoard.ThumbnailURL).
		SetStatus(params.Board.EntBoard.Status).
		SetUpdatedAt(time.Now())

	entBoard, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	board := model.NewBoard(model.NewBoardParams{
		EntBoard: entBoard,
	})

	return board, nil
}

type BoardDatasourceUpdateStatusParams struct {
	Ctx   context.Context
	Board *model.Board
}

func (ds *BoardDatasource) UpdateStatus(params BoardDatasourceUpdateStatusParams) (*model.Board, error) {
	update := ds.client.Board.UpdateOneID(params.Board.EntBoard.ID)

	update = update.
		SetStatus(params.Board.EntBoard.Status).
		SetUpdatedAt(time.Now())

	entBoard, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	board := model.NewBoard(model.NewBoardParams{
		EntBoard: entBoard,
	})

	return board, nil
}
