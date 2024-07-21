package service

import (
	"context"
	"errors"
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

type ThreadApplicationServiceFindAllParams struct {
	Ctx    context.Context
	Qs     request.ThreadFindAllRequest
	UserID int
}

func (svc *ThreadApplicationService) FindAll(params ThreadApplicationServiceFindAllParams) ([]*resource.ThreadResource, error) {
	var threads []*model.Thread
	var err error

	criteria := params.Qs.QueryCriteria

	switch criteria {
	case "popularity":
		threads, err = svc.threadDatasource.FindByPopularity(datasource.ThreadDatasourceFindByPopularityParams{
			Ctx:    params.Ctx,
			Limit:  params.Qs.Limit,
			Offset: params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "related":
		if len(params.Qs.ThreadIDs) == 0 {
			return nil, errors.New("関連順のためにthreadIdsが必要です")
		}
		threads, err = svc.threadDatasource.FindByRelatedTag(datasource.ThreadDatasourceFindByRelatedTagParams{
			Ctx:       params.Ctx,
			ThreadIDs: params.Qs.ThreadIDs,
			Limit:     params.Qs.Limit,
			Offset:    params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "keyword":
		if params.Qs.Keyword == "" {
			return nil, errors.New("Keywordが必要です")
		}
		threads, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:     params.Ctx,
			Keyword: params.Qs.Keyword,
			Limit:   params.Qs.Limit,
			Offset:  params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "history":
		if len(params.Qs.ThreadIDs) == 0 {
			return nil, errors.New("History順のためにthreadIdsが必要です")
		}
		threads, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:       params.Ctx,
			ThreadIDs: params.Qs.ThreadIDs,
			Limit:     params.Qs.Limit,
			Offset:    params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "board":
		if params.Qs.BoardID == 0 {
			return nil, errors.New("BoardIDが必要です")
		}
		threads, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:     params.Ctx,
			BoardID: params.Qs.BoardID,
			Limit:   params.Qs.Limit,
			Offset:  params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "owner":
		if params.UserID == 0 {
			return nil, errors.New("UserIDが必要です")
		}
		threads, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:    params.Ctx,
			UserID: params.UserID,
			Limit:  params.Qs.Limit,
			Offset: params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	case "newest":
		threads, err = svc.threadDatasource.FindAll(datasource.ThreadDatasourceFindAllParams{
			Ctx:    params.Ctx,
			Limit:  params.Qs.Limit,
			Offset: params.Qs.Offset,
		})
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("無効なQueryCriteriaです")
	}

	var dto []*resource.ThreadResource
	for _, thread := range threads {
		resource := resource.NewThreadResource(resource.NewThreadResourceParams{
			Thread:             thread,
			ThreadCommentCount: thread.ThreadCommentCount,
		})
		dto = append(dto, resource)
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
