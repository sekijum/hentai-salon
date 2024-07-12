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
	threadsByCriteria := make(map[string][]*model.Thread)

	for _, criteria := range qs.QueryCriteria {
		switch criteria {
		case "popularity":
			threads, err := svc.threadDatasource.FindByPopularity(ctx, qs.Limit, qs.Offset)
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByPopular"] = threads
		case "newest":
			threads, err := svc.threadDatasource.FindByNewest(ctx, qs.Limit, qs.Offset)
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByNewest"] = threads
		case "history":
			if len(qs.ThreadIds) == 0 {
				return nil, errors.New("history順のためにthreadIdsが必要です")
			}
			threads, err := svc.threadDatasource.FindByHistory(ctx, qs.ThreadIds)
			if err != nil {
				return nil, err
			}
			threadMap := make(map[int]*model.Thread)
			for _, thread := range threads {
				threadMap[thread.EntThread.ID] = thread
			}
			sortedThreads := make([]*model.Thread, 0, len(qs.ThreadIds))
			for i := len(qs.ThreadIds) - 1; i >= 0; i-- {
				id := qs.ThreadIds[i]
				if thread, exists := threadMap[id]; exists {
					sortedThreads = append(sortedThreads, thread)
				}
			}
			threadsByCriteria["threadsByHistory"] = sortedThreads
		case "keyword":
			if qs.Keyword != "" {
				threads, err := svc.threadDatasource.FindByKeyword(ctx, qs.Keyword, qs.Limit, qs.Offset)
				if err != nil {
					return nil, err
				}
				threadsByCriteria["threadsByKeyword"] = threads
			}
		case "related":
			if len(qs.ThreadIds) == 0 {
				return nil, errors.New("関連順のためにthreadIdsが必要です")
			}
			threads, err := svc.threadDatasource.FindByRelatedTags(ctx, qs.ThreadIds, qs.Limit, qs.Offset)
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByRelated"] = threads
		case "board":
			if qs.BoardId == 0 {
				return nil, errors.New("BoardIDが必要です")
			}
			threads, err := svc.threadDatasource.FindByBoardId(ctx, qs.BoardId, qs.Limit, qs.Offset)
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByBoard"] = threads
		}
	}

	threadResources := make(map[string][]*resource.ThreadResource)
	threadResources["threadsByPopular"] = []*resource.ThreadResource{}
	threadResources["threadsByNewest"] = []*resource.ThreadResource{}
	threadResources["threadsByHistory"] = []*resource.ThreadResource{}
	threadResources["threadsByKeyword"] = []*resource.ThreadResource{}
	threadResources["threadsByRelated"] = []*resource.ThreadResource{}
	threadResources["threadsByBoard"] = []*resource.ThreadResource{}

	for key, threads := range threadsByCriteria {
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
