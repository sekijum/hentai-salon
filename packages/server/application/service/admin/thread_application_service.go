package service_admin

import (
	"context"
	"server/domain/model"
	datasource_admin "server/infrastructure/datasource/admin"
	"server/infrastructure/ent"
	request_admin "server/presentation/request/admin"
	response_admin "server/presentation/response/admin"
)

type ThreadApplicationService struct {
	threadDatasource *datasource_admin.ThreadDatasource
}

func NewThreadApplicationService(threadDatasource *datasource_admin.ThreadDatasource) *ThreadApplicationService {
	return &ThreadApplicationService{threadDatasource: threadDatasource}
}

type ThreadApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request_admin.ThreadFindAllRequest
}

func (svc *ThreadApplicationService) FindAll(params ThreadApplicationServiceFindAllParams) (*response_admin.Collection[*response_admin.ThreadResponse], error) {
	threadList, err := svc.threadDatasource.FindAll(datasource_admin.ThreadDatasourceFindAllParams{
		Ctx:     params.Ctx,
		Limit:   params.Qs.Limit,
		Offset:  params.Qs.Offset,
		Sort:    params.Qs.Sort,
		Order:   params.Qs.Order,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	threadCount, err := svc.threadDatasource.GetThreadCount(datasource_admin.ThreadDatasourceGetThreadCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	var threadResponseList []*response_admin.ThreadResponse
	for _, thread_i := range threadList {
		threadResponse := response_admin.NewThreadResponse(response_admin.NewThreadResponseParams{
			Thread:       thread_i,
			IncludeBoard: true,
		})
		threadResponseList = append(threadResponseList, threadResponse)
	}

	dto := response_admin.NewCollection(response_admin.NewCollectionParams[*response_admin.ThreadResponse]{
		Data:       threadResponseList,
		TotalCount: threadCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type ThreadApplicationServiceFindByIDParams struct {
	Ctx      context.Context
	ThreadID int
	Qs       request_admin.ThreadFindByIDRequest
}

func (svc *ThreadApplicationService) FindByID(params ThreadApplicationServiceFindByIDParams) (*response_admin.ThreadResponse, error) {
	thread, err := svc.threadDatasource.FindByID(datasource_admin.ThreadDatasourceFindByIDParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
		Limit:    params.Qs.Limit,
		Offset:   params.Qs.Offset,
		Order:    params.Qs.Order,
		Sort:     params.Qs.Sort,
		Keyword:  params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	threadCommentCount, err := svc.threadDatasource.GetThreadCommentCount(datasource_admin.ThreadDatasourceGetThreadCommentCountParams{
		Ctx:      params.Ctx,
		ThreadID: params.ThreadID,
		Keyword:  params.Qs.Keyword,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewThreadResponse(response_admin.NewThreadResponseParams{
		Thread:          thread,
		CommentCount:    threadCommentCount,
		IncludeComments: true,
	})

	return dto, nil
}

type ThreadApplicationServiceUpdateParams struct {
	Ctx      context.Context
	ThreadID int
	Body     request_admin.ThreadUpdateRequest
}

func (svc *ThreadApplicationService) Update(params ThreadApplicationServiceUpdateParams) (*response_admin.ThreadResponse, error) {
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

	thread, err := svc.threadDatasource.Update(datasource_admin.ThreadDatasourceUpdateParams{
		Ctx:    params.Ctx,
		Thread: *thread,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewThreadResponse(response_admin.NewThreadResponseParams{
		Thread: thread,
	})

	return dto, nil
}

type ThreadApplicationServiceUpdateStatusParams struct {
	Ctx      context.Context
	ThreadID int
	Body     request_admin.ThreadUpdateStatusRequest
}

func (svc *ThreadApplicationService) UpdateStatus(params ThreadApplicationServiceUpdateStatusParams) error {
	thread := model.NewThread(model.NewThreadParams{
		EntThread: &ent.Thread{
			ID: params.ThreadID,
		},
		OptionList: []func(*model.Thread){
			model.WithThreadStatus(model.ThreadStatus(params.Body.Status)),
		},
	})

	_, err := svc.threadDatasource.UpdateStatus(datasource_admin.ThreadDatasourceUpdateStatusParams{
		Ctx:    params.Ctx,
		Thread: *thread,
	})
	if err != nil {
		return err
	}

	return nil
}

type ThreadApplicationServiceDeleteParams struct {
	Ctx      context.Context
	ThreadID int
}

func (svc *ThreadApplicationService) Delete(params ThreadApplicationServiceDeleteParams) error {
	err := svc.threadDatasource.Delete(datasource_admin.ThreadDatasourceDeleteParams{
		Ctx:      params.Ctx,
		ThreadId: params.ThreadID,
	})
	if err != nil {
		return err
	}

	return nil
}
