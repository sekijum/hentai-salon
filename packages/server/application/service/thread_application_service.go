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

type ThreadApplicationServiceFindAllListParams struct {
	Ctx context.Context
	Qs  request.ThreadFindAllRequest
}

func (svc *ThreadApplicationService) FindAllList(params ThreadApplicationServiceFindAllListParams) (map[string][]*resource.ThreadResource, error) {
	threadsByCriteria := make(map[string][]*model.Thread)

	for _, criteria := range params.Qs.QueryCriteria {
		switch criteria {
		case "popularity":
			threads, err := svc.threadDatasource.FindByPopularity(datasource.ThreadDatasourceFindByPopularityParams{
				Ctx:    params.Ctx,
				Limit:  params.Qs.Limit,
				Offset: params.Qs.Offset,
			})
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByPopular"] = threads
		case "newest":
			threads, err := svc.threadDatasource.FindByNewest(datasource.ThreadDatasourceFindByNewestParams{
				Ctx:    params.Ctx,
				Limit:  params.Qs.Limit,
				Offset: params.Qs.Offset,
			})
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByNewest"] = threads
		case "history":
			if len(params.Qs.ThreadIDs) == 0 {
				return nil, errors.New("history順のためにthreadIdsが必要です")
			}
			threads, err := svc.threadDatasource.FindByHistory(datasource.ThreadDatasourceFindByHistoryParams{
				Ctx:       params.Ctx,
				ThreadIDs: params.Qs.ThreadIDs,
				Limit:     params.Qs.Limit,
				Offset:    params.Qs.Offset,
			})
			if err != nil {
				return nil, err
			}
			threadMap := make(map[int]*model.Thread)
			for _, thread := range threads {
				threadMap[thread.EntThread.ID] = thread
			}
			sortedThreads := make([]*model.Thread, 0, len(params.Qs.ThreadIDs))
			for _, id := range params.Qs.ThreadIDs {
				if thread, exists := threadMap[id]; exists {
					sortedThreads = append(sortedThreads, thread)
				}
			}
			threadsByCriteria["threadsByHistory"] = sortedThreads
		case "keyword":
			if params.Qs.Keyword != "" {
				threads, err := svc.threadDatasource.FindByKeyword(datasource.ThreadDatasourceFindByKeywordParams{
					Ctx:     params.Ctx,
					Keyword: params.Qs.Keyword,
					Limit:   params.Qs.Limit,
					Offset:  params.Qs.Offset,
				})
				if err != nil {
					return nil, err
				}
				threadsByCriteria["threadsByKeyword"] = threads
			}
		case "related":
			if len(params.Qs.ThreadIDs) == 0 {
				return nil, errors.New("関連順のためにthreadIdsが必要です")
			}
			threads, err := svc.threadDatasource.FindByRelatedTag(datasource.ThreadDatasourceFindByRelatedTagParams{
				Ctx:       params.Ctx,
				ThreadIDs: params.Qs.ThreadIDs,
				Limit:     params.Qs.Limit,
				Offset:    params.Qs.Offset,
			})
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByRelated"] = threads
		case "board":
			if params.Qs.BoardID == 0 {
				return nil, errors.New("BoardIDが必要です")
			}
			threads, err := svc.threadDatasource.FindByBoardID(datasource.ThreadDatasourceFindByBoardIDParams{
				Ctx:     params.Ctx,
				BoardID: params.Qs.BoardID,
				Limit:   params.Qs.Limit,
				Offset:  params.Qs.Offset,
			})
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
			resource := resource.NewThreadResource(resource.NewThreadResourceParams{Thread: thread})
			threadResources[key] = append(threadResources[key], resource)
		}
	}

	return threadResources, nil
}

type ThreadApplicationServiceFindByIDParams struct {
	Ctx      context.Context
	ThreadID int
	Qs       request.ThreadFindByIdRequest
}

func (svc *ThreadApplicationService) FindByID(params ThreadApplicationServiceFindByIDParams) (*resource.ThreadResource, error) {
	thread, err := svc.threadDatasource.FindById(datasource.ThreadDatasourceFindByIDParams{
		Ctx:       params.Ctx,
		SortOrder: params.Qs.SortOrder,
		Limit:     params.Qs.Limit,
		Offset:    params.Qs.Offset,
		ThreadID:  params.ThreadID,
	})
	if err != nil {
		return nil, err
	}

	resource := resource.NewThreadResource(resource.NewThreadResourceParams{
		Thread: thread,
		Limit:  params.Qs.Limit,
		Offset: params.Qs.Offset,
	})
	return resource, nil
}

type ThreadApplicationServiceCreateParams struct {
	Ctx    context.Context
	GinCtx *gin.Context
	Body   request.ThreadCreateRequest
}

func (svc *ThreadApplicationService) Create(params ThreadApplicationServiceCreateParams) (*resource.ThreadResource, error) {
	userID, exists := params.GinCtx.Get("userID")
	if !exists {
		return nil, errors.New("ユーザーIDがコンテキストに存在しません")
	}

	if duplicated, err := svc.threadDomainService.IsTitleDuplicated(domainService.ThreadDomainServiceTitleDuplicatedParams{
		Ctx:   params.Ctx,
		Title: params.Body.Title,
	}); err != nil || duplicated {
		return nil, err
	}

	tx, err := svc.client.Tx(params.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	modelTags, err := svc.tagDatasource.CreateManyTx(datasource.TagDatasourceCreateManyTxParams{
		Ctx: params.Ctx,
	})
	if err != nil {
		return nil, err
	}

	var tagIDs []int
	for _, tag := range modelTags {
		tagIDs = append(tagIDs, tag.EntTag.ID)
	}

	thread := &model.Thread{
		EntThread: &ent.Thread{
			Title:     params.Body.Title,
			BoardID:   params.Body.BoardId,
			UserID:    userID.(int),
			IPAddress: params.GinCtx.ClientIP(),
			Status:    int(model.ThreadStatusOpen),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if params.Body.Description != nil {
		thread.EntThread.Description = params.Body.Description
	}
	if params.Body.ThumbnailURL != nil {
		thread.EntThread.ThumbnailURL = params.Body.ThumbnailURL
	}

	savedThread, err := svc.threadDatasource.CreateTx(datasource.ThreadDatasourceCreateTxParams{
		Ctx:    params.Ctx,
		Tx:     tx,
		Thread: thread,
		TagIDs: tagIDs,
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resource := resource.NewThreadResource(resource.NewThreadResourceParams{Thread: savedThread})

	return resource, nil
}
