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

type BoardApplicationServiceFindAllParams struct {
	Ctx context.Context
}

func (svc *BoardApplicationService) FindAll(params BoardApplicationServiceFindAllParams) ([]*resource.BoardResource, error) {
	boards, err := svc.boardDatasource.FindAll(datasource.BoardDatasourceFindAllParams{
		Ctx: params.Ctx,
	})
	if err != nil {
		return nil, err
	}

	var listResource []*resource.BoardResource
	for _, board := range boards {
		listResource = append(listResource, resource.NewBoardResource(resource.NewBoardResourceParams{
			Board: board,
		}))
	}

	return listResource, nil
}

type BoardApplicationServiceCreateParams struct {
	Ctx    context.Context
	GinCtx *gin.Context
	Body   request.BoardCreateRequest
}

func (svc *BoardApplicationService) Create(params BoardApplicationServiceCreateParams) (*resource.BoardResource, error) {
	userID, exists := params.GinCtx.Get("userID")
	if !exists {
		return nil, errors.New("ユーザーIDがコンテキストに存在しません")
	}

	if duplicated, err := svc.boardDomainService.IsTitleDuplicated(domainService.BoardDomainServiceIsTitleDuplicatedParams{
		Ctx:   params.Ctx,
		Title: params.Body.Title,
	}); err != nil || duplicated {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("板タイトルが重複しています")
	}

	board := &model.Board{
		EntBoard: &ent.Board{
			Title:     params.Body.Title,
			UserID:    userID.(int),
			Status:    int(model.BoardStatusPublic),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if params.Body.Description != nil {
		board.EntBoard.Description = params.Body.Description
	}
	if params.Body.ThumbnailURL != nil {
		board.EntBoard.ThumbnailURL = params.Body.ThumbnailURL
	}

	board, err := svc.boardDatasource.Create(datasource.BoardDatasourceCreateParams{
		Ctx:   params.Ctx,
		Board: board,
	})
	if err != nil {
		return nil, err
	}

	resource := resource.NewBoardResource(resource.NewBoardResourceParams{
		Board: board,
	})

	return resource, nil
}
