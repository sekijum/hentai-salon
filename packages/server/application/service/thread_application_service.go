package service

import (
	"context"
	"server/domain/model"
	domainService "server/domain/service"
	"server/infrastructure/datasource"

	"errors"
	request "server/presentation/request/thread"
	"time"

	"github.com/gin-gonic/gin"
)

type ThreadApplicationService struct {
	threadDatasource    *datasource.ThreadDatasource
	threadDomainService *domainService.ThreadDomainService
}

func NewThreadApplicationService(
	threadDatasource *datasource.ThreadDatasource,
	threadDomainService *domainService.ThreadDomainService,
) *ThreadApplicationService {
	return &ThreadApplicationService{
		threadDatasource:    threadDatasource,
		threadDomainService: threadDomainService,
	}
}

func (svc *ThreadApplicationService) FindAll(ctx context.Context) ([]*model.Thread, error) {
	return svc.threadDatasource.FindAll(ctx)
}

func (svc *ThreadApplicationService) Create(
	ctx context.Context,
	ginCtx *gin.Context,
	body request.ThreadCreateRequest,
) error {
	userId, exists := ginCtx.Get("user_id")
	if !exists {
		return errors.New("ユーザーIDがコンテキストに存在しません")
	}

	if duplicated, err := svc.threadDomainService.IsTitleDuplicated(ctx, body.Title); err != nil || duplicated {
		if err != nil {
			return err
		}
		return errors.New("スレタイが重複しています")
	}
	thread := &model.Thread{
		Title:             body.Title,
		Description:       body.Description,
		UserId:            userId.(int),
		ThumbnailUrl:      body.ThumbnailUrl,
		IpAddress:         ginCtx.ClientIP(),
		Status:            model.ThreadStatusOpen,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := thread.Validate(); err != nil {
		return err
	}

	_, err := svc.threadDatasource.Create(ctx, thread)
	if err != nil {
		return err
	}

	return nil
}
