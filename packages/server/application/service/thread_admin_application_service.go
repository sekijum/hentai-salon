package service

import (
	"context"
	"server/infrastructure/datasource"
	"server/presentation/request"
	"server/presentation/resource"
	"time"
)

type ThreadAdminApplicationService struct {
	threadAdminDatasource *datasource.ThreadAdminDatasource
}

func NewThreadAdminApplicationService(
	threadAdminDatasource *datasource.ThreadAdminDatasource,
) *ThreadAdminApplicationService {
	return &ThreadAdminApplicationService{
		threadAdminDatasource: threadAdminDatasource,
	}
}

type ThreadAdminApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.ThreadAdminFindAllRequest
}

func (svc *ThreadAdminApplicationService) FindAll(params ThreadAdminApplicationServiceFindAllParams) (*resource.ListResource[*resource.ThreadAdminResource], error) {
	threads, err := svc.threadAdminDatasource.FindAll(datasource.ThreadAdminDatasourceFindAllParams{
		Ctx:       params.Ctx,
		Limit:     params.Qs.Limit,
		Offset:    params.Qs.Offset,
		SortKey:   params.Qs.SortKey,
		SortOrder: params.Qs.SortOrder,
		Keyword:   params.Qs.Keyword,
		Status:    params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	totalCount, err := svc.threadAdminDatasource.GetThreadCount(datasource.ThreadAdminDatasourceGetThreadCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var threadResources []*resource.ThreadAdminResource
	for _, thread := range threads {
		threadResource := resource.NewThreadAdminResource(resource.NewThreadAdminResourceParams{
			Thread: thread,
			Limit:  params.Qs.Limit,
			Offset: params.Qs.Offset,
		})
		threadResources = append(threadResources, threadResource)
	}

	listResource := &resource.ListResource[*resource.ThreadAdminResource]{
		TotalCount: totalCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
		Data:       threadResources,
	}

	return listResource, nil
}

type ThreadAdminApplicationServiceFindByIDParams struct {
	Ctx      context.Context
	ThreadID int
	Qs       request.ThreadAdminFindByIDRequest
}

func (svc *ThreadAdminApplicationService) FindByID(params ThreadAdminApplicationServiceFindByIDParams) (*resource.ThreadAdminResource, error) {
	thread, err := svc.threadAdminDatasource.FindByID(datasource.ThreadAdminDatasourceFindByIDParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
		Limit:    params.Qs.Limit,
		Offset:   params.Qs.Offset,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewThreadAdminResource(resource.NewThreadAdminResourceParams{
		Thread: thread,
		Limit:  params.Qs.Limit,
		Offset: params.Qs.Offset,
	})

	return dto, nil
}

type ThreadAdminApplicationServiceUpdateParams struct {
	Ctx      context.Context
	ThreadID int
	Body     request.ThreadAdminUpdateRequest
}

func (svc *ThreadAdminApplicationService) Update(params ThreadAdminApplicationServiceUpdateParams) (*resource.ThreadAdminResource, error) {
	thread, err := svc.threadAdminDatasource.FindByID(datasource.ThreadAdminDatasourceFindByIDParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
		Limit:    0,
		Offset:   0,
	})
	if err != nil {
		return nil, err
	}

	if params.Body.Title != nil {
		thread.EntThread.Title = *params.Body.Title
	}
	if params.Body.Description != nil {
		thread.EntThread.Description = params.Body.Description
	}
	if params.Body.Status != nil {
		thread.EntThread.Status = *params.Body.Status
	}
	if params.Body.ThumbnailURL != nil {
		thread.EntThread.ThumbnailURL = params.Body.ThumbnailURL
	}
	thread.EntThread.UpdatedAt = time.Now()

	updatedThread, err := svc.threadAdminDatasource.Update(datasource.ThreadAdminDatasourceUpdateParams{
		Ctx:    params.Ctx,
		Thread: *thread,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewThreadAdminResource(resource.NewThreadAdminResourceParams{
		Thread: updatedThread,
		Limit:  0,
		Offset: 0,
	})

	return dto, nil
}
