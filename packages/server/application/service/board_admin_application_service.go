package service

import (
	"context"
	"errors"
	"server/domain/model"
	domainService "server/domain/service"
	"server/infrastructure/datasource"
	request "server/presentation/request/board"
	"time"

	"github.com/gin-gonic/gin"
)

type BoardAdminApplicationService struct {
	boardDatasource     *datasource.BoardDatasource
	boardDomainService  *domainService.BoardDomainService
	threadDomainService *domainService.ThreadDomainService
}

func NewBoardAdminApplicationService(
	boardDatasource *datasource.BoardDatasource,
	boardDomainService *domainService.BoardDomainService,
	threadDomainService *domainService.ThreadDomainService,
) *BoardAdminApplicationService {
	return &BoardAdminApplicationService{
		boardDatasource:     boardDatasource,
		boardDomainService:  boardDomainService,
		threadDomainService: threadDomainService,
	}
}

func (svc *BoardAdminApplicationService) Create(
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
		Title:        body.Title,
		Description:  body.Description,
		UserId:       userId.(int),
		ThumbnailUrl: body.ThumbnailUrl,
		Status:       model.BoardStatusPublic,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := board.Validate(); err != nil {
		return err
	}

	_, err := svc.boardDatasource.Create(ctx, board)
	if err != nil {
		return err
	}

	return nil
}
