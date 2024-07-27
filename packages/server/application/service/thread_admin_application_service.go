package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/resource"
)

type ThreadAdminApplicationService struct {
	threadAdminDatasource *datasource.ThreadAdminDatasource
}

func NewThreadAdminApplicationService(threadAdminDatasource *datasource.ThreadAdminDatasource) *ThreadAdminApplicationService {
	return &ThreadAdminApplicationService{threadAdminDatasource: threadAdminDatasource}
}

type ThreadAdminApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.ThreadAdminFindAllRequest
}

func (svc *ThreadAdminApplicationService) FindAll(params ThreadAdminApplicationServiceFindAllParams) (*resource.Collection[*resource.ThreadAdminResource], error) {
	threadList, err := svc.threadAdminDatasource.FindAll(datasource.ThreadAdminDatasourceFindAllParams{
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

	threadCount, err := svc.threadAdminDatasource.GetThreadCount(datasource.ThreadAdminDatasourceGetThreadCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var threadResourceList []*resource.ThreadAdminResource
	for _, thread_i := range threadList {
		threadResource := resource.NewThreadAdminResource(resource.NewThreadAdminResourceParams{Thread: thread_i})
		threadResourceList = append(threadResourceList, threadResource)
	}

	dto := resource.NewCollection(resource.NewCollectionParams[*resource.ThreadAdminResource]{
		Data:       threadResourceList,
		TotalCount: threadCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
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

	dto := resource.NewThreadAdminResource(resource.NewThreadAdminResourceParams{Thread: thread})

	return dto, nil
}

type ThreadAdminApplicationServiceUpdateParams struct {
	Ctx      context.Context
	ThreadID int
	Body     request.ThreadAdminUpdateRequest
}

func (svc *ThreadAdminApplicationService) Update(params ThreadAdminApplicationServiceUpdateParams) (*resource.ThreadAdminResource, error) {
	thread := model.NewThread(model.NewThreadParams{
		EntThread: &ent.Thread{
			ID:           params.ThreadID,
			Title:        params.Body.Title,
			Description:  params.Body.Description,
			ThumbnailURL: params.Body.ThumbnailURL,
		},
		OptionList: []func(*model.Thread){
			model.WithThreadStatus(model.ThreadStatus(params.Body.Status)),
		},
	})

	thread, err := svc.threadAdminDatasource.Update(datasource.ThreadAdminDatasourceUpdateParams{
		Ctx:    params.Ctx,
		Thread: *thread,
	})
	if err != nil {
		return nil, err
	}

	dto := resource.NewThreadAdminResource(resource.NewThreadAdminResourceParams{
		Thread: thread,
	})

	return dto, nil
}
