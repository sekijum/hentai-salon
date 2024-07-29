package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/response"
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

func (svc *ThreadAdminApplicationService) FindAll(params ThreadAdminApplicationServiceFindAllParams) (*response.Collection[*response.ThreadAdminResponse], error) {
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

	var threadResponseList []*response.ThreadAdminResponse
	for _, thread_i := range threadList {
		threadResponse := response.NewThreadAdminResponse(response.NewThreadAdminResponseParams{Thread: thread_i})
		threadResponseList = append(threadResponseList, threadResponse)
	}

	dto := response.NewCollection(response.NewCollectionParams[*response.ThreadAdminResponse]{
		Data:       threadResponseList,
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

func (svc *ThreadAdminApplicationService) FindByID(params ThreadAdminApplicationServiceFindByIDParams) (*response.ThreadAdminResponse, error) {
	thread, err := svc.threadAdminDatasource.FindByID(datasource.ThreadAdminDatasourceFindByIDParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
		Limit:    params.Qs.Limit,
		Offset:   params.Qs.Offset,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewThreadAdminResponse(response.NewThreadAdminResponseParams{Thread: thread})

	return dto, nil
}

type ThreadAdminApplicationServiceUpdateParams struct {
	Ctx      context.Context
	ThreadID int
	Body     request.ThreadAdminUpdateRequest
}

func (svc *ThreadAdminApplicationService) Update(params ThreadAdminApplicationServiceUpdateParams) (*response.ThreadAdminResponse, error) {
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

	dto := response.NewThreadAdminResponse(response.NewThreadAdminResponseParams{
		Thread: thread,
	})

	return dto, nil
}
