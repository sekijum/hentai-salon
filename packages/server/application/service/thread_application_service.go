package service

import (
	"context"
	"errors"
	"fmt"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	request "server/presentation/request"
	resource "server/presentation/resource"
	"time"
)

type ThreadApplicationService struct {
	client           *ent.Client
	threadDatasource *datasource.ThreadDatasource
	tagDatasource    *datasource.TagDatasource
}

func NewThreadApplicationService(
	client *ent.Client,
	threadDatasource *datasource.ThreadDatasource,
	tagDatasource *datasource.TagDatasource,
) *ThreadApplicationService {
	return &ThreadApplicationService{
		client:           client,
		threadDatasource: threadDatasource,
		tagDatasource:    tagDatasource,
	}
}

type ThreadApplicationServiceFindAllListParams struct {
	Ctx    context.Context
	Qs     request.ThreadFindAllRequest
	UserID int
}

func (svc *ThreadApplicationService) FindAllList(params ThreadApplicationServiceFindAllListParams) (map[string][]*resource.ThreadResource, error) {
	threadsByCriteria := make(map[string][]*model.Thread)

	fmt.Printf("params.UserID", params.UserID)

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
		case "owner":
			if params.UserID == 0 {
				return nil, errors.New("UserIDが必要です")
			}
			threads, err := svc.threadDatasource.FindByUserID(datasource.ThreadDatasourceFindByUserIDParams{
				Ctx:    params.Ctx,
				UserID: params.UserID,
				Limit:  params.Qs.Limit,
				Offset: params.Qs.Offset,
			})
			if err != nil {
				return nil, err
			}
			threadsByCriteria["threadsByOwner"] = threads
		}
	}

	dto := make(map[string][]*resource.ThreadResource)
	dto["threadsByPopular"] = []*resource.ThreadResource{}
	dto["threadsByNewest"] = []*resource.ThreadResource{}
	dto["threadsByHistory"] = []*resource.ThreadResource{}
	dto["threadsByKeyword"] = []*resource.ThreadResource{}
	dto["threadsByRelated"] = []*resource.ThreadResource{}
	dto["threadsByBoard"] = []*resource.ThreadResource{}
	dto["threadsByOwner"] = []*resource.ThreadResource{}

	for key, threads := range threadsByCriteria {
		for _, thread := range threads {
			resource := resource.NewThreadResource(resource.NewThreadResourceParams{Thread: thread, ThreadCommentCount: thread.ThreadCommentCount})
			dto[key] = append(dto[key], resource)
		}
	}

	return dto, nil
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

	fmt.Print("ThreadCommentReplyCountMap", thread.ThreadCommentReplyCountMap)
	dto := resource.NewThreadResource(resource.NewThreadResourceParams{
		Thread:                     thread,
		Limit:                      params.Qs.Limit,
		Offset:                     params.Qs.Offset,
		ThreadCommentCount:         thread.ThreadCommentCount,
		ThreadCommentReplyCountMap: thread.ThreadCommentReplyCountMap,
	})
	return dto, nil
}

type ThreadApplicationServiceCreateParams struct {
	Ctx      context.Context
	UserID   int
	ClientIP string
	Body     request.ThreadCreateRequest
}

func (svc *ThreadApplicationService) Create(params ThreadApplicationServiceCreateParams) (*resource.ThreadResource, error) {
	threads, err := svc.threadDatasource.FindByTitle(datasource.ThreadDatasourceFindByTitleParams{
		Ctx:   params.Ctx,
		Title: params.Body.Title,
	})
	if err != nil {
		return nil, err
	}
	if len(threads) > 0 {
		return nil, errors.New("スレタイが重複しています")
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
			UserID:    params.UserID,
			IPAddress: params.ClientIP,
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

	dto := resource.NewThreadResource(resource.NewThreadResourceParams{Thread: savedThread})

	return dto, nil
}
