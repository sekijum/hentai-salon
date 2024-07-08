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

type ThreadApplicationService struct {
	client              *ent.Client
	threadDatasource    *datasource.ThreadDatasource
	tagDatasource       *datasource.TagDatasource
	threadDomainService *domainService.ThreadDomainService
}

func NewThreadApplicationService(
	client *ent.Client,
	threadDatasource *datasource.ThreadDatasource,
	tagDatasource *datasource.TagDatasource,
	threadDomainService *domainService.ThreadDomainService,
) *ThreadApplicationService {
	return &ThreadApplicationService{
		client:              client,
		threadDatasource:    threadDatasource,
		tagDatasource:       tagDatasource,
		threadDomainService: threadDomainService,
	}
}

func (svc *ThreadApplicationService) FindAll(ctx context.Context, qs request.ThreadFindAllRequest) (map[string][]*resource.ThreadResource, error) {
	threadsByOrders := make(map[string][]*model.Thread)

	for _, order := range qs.Orders {
		switch order {
		case "popularity":
			threads, err := svc.threadDatasource.FindByPopularity(ctx, qs.Limit, qs.Offset)
			if err != nil {
				return nil, err
			}
			threadsByOrders["threadsByPopular"] = threads
		case "newest":
			threads, err := svc.threadDatasource.FindByNewest(ctx, qs.Limit, qs.Offset)
			if err != nil {
				return nil, err
			}
			threadsByOrders["threadsByNewest"] = threads
		case "histories":
			if len(qs.ThreadIds) == 0 {
				return nil, errors.New("threadIds is required for histories order")
			}
			threads, err := svc.threadDatasource.FindByHistories(ctx, qs.ThreadIds)
			if err != nil {
				return nil, err
			}
			threadsByOrders["threadsByHistories"] = threads
		}
	}

	threadResources := make(map[string][]*resource.ThreadResource)
	threadResources["threadsByPopular"] = []*resource.ThreadResource{}
	threadResources["threadsByNewest"] = []*resource.ThreadResource{}
	threadResources["threadsByHistories"] = []*resource.ThreadResource{}

	for key, threads := range threadsByOrders {
		for _, thread := range threads {
			res := resource.NewThreadResource(thread, 0, 0)
			threadResources[key] = append(threadResources[key], res)
		}
	}

	return threadResources, nil
}

func (svc *ThreadApplicationService) FindById(ctx context.Context, id int, qs request.ThreadFindByIdRequest) (*resource.ThreadResource, error) {
	thread, err := svc.threadDatasource.FindById(ctx, id, qs.Limit, qs.Offset)
	if err != nil {
		return nil, err
	}
	return resource.NewThreadResource(thread, qs.Limit, qs.Offset), nil
}

func (svc *ThreadApplicationService) Create(ctx context.Context, ginCtx *gin.Context, body request.ThreadCreateRequest) (*resource.ThreadResource, error) {
	userId, exists := ginCtx.Get("user_id")
	if !exists {
		return nil, errors.New("ユーザーIDがコンテキストに存在しません")
	}

	if duplicated, err := svc.threadDomainService.IsTitleDuplicated(ctx, body.Title); err != nil || duplicated {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("スレタイが重複しています")
	}

	tx, err := svc.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	modelTags, err := svc.tagDatasource.CreateManyTx(ctx, tx, body.TagNames)
	if err != nil {
		return nil, err
	}

	tagIDs := make([]int, len(modelTags))
	for i, tag := range modelTags {
		tagIDs[i] = tag.EntTag.ID
	}

	description := ""
	if body.Description != nil {
		description = *body.Description
	}

	thumbnailUrl := ""
	if body.ThumbnailUrl != nil {
		thumbnailUrl = *body.ThumbnailUrl
	}

	thread := &model.Thread{
		EntThread: &ent.Thread{
			Title:        body.Title,
			BoardID:      body.BoardId,
			Description:  description,
			UserID:       userId.(int),
			ThumbnailURL: thumbnailUrl,
			IPAddress:    ginCtx.ClientIP(),
			Status:       int(model.ThreadStatusOpen),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	savedThread, err := svc.threadDatasource.CreateTx(ctx, tx, thread, tagIDs)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return resource.NewThreadResource(savedThread, 0, 0), nil
}
