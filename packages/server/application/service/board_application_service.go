package service

import (
	"context"
	"errors"
	"server/domain/model"
	domainService "server/domain/service"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	request "server/presentation/request"
	resource "server/presentation/resource"
	"time"

	"github.com/gin-gonic/gin"
)

type BoardApplicationService struct {
	boardDatasource    *datasource.BoardDatasource
	boardDomainService *domainService.BoardDomainService
}

func NewBoardApplicationService(
	boardDatasource *datasource.BoardDatasource,
	boardDomainService *domainService.BoardDomainService,
) *BoardApplicationService {
	return &BoardApplicationService{
		boardDatasource:    boardDatasource,
		boardDomainService: boardDomainService,
	}
}

func (svc *BoardApplicationService) FindAll(ctx context.Context) ([]*resource.BoardResource, error) {
	boards, err := svc.boardDatasource.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var boardResources []*resource.BoardResource
	for _, board := range boards {
		boardResources = append(boardResources, resource.NewBoardResource(board))
	}

	return boardResources, nil
}

func (svc *BoardApplicationService) Create(
	ctx context.Context,
	ginCtx *gin.Context,
	body request.BoardCreateRequest,
) error {
	userId, exists := ginCtx.Get("user_id")
	if !exists {
		return errors.New("ユーザーIDがコンテキストに存在しません")
	}

	if duplicated, err := svc.boardDomainService.IsTitleDuplicated(ctx, body.Title); err != nil || duplicated {
		if err != nil {
			return err
		}
		return errors.New("板タイトルが重複しています")
	}

	board := &model.Board{
		EntBoard: &ent.Board{
			Title:       body.Title,
			Description: *body.Description,
			UserID:      userId.(int),
			Status:      int(model.BoardStatusPublic),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if body.ThumbnailUrl != nil {
		board.EntBoard.ThumbnailURL = *body.ThumbnailUrl
	}

	_, err := svc.boardDatasource.Create(ctx, board)
	if err != nil {
		return err
	}

	return nil
}
