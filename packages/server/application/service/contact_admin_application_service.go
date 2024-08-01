package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/response"
)

type ContactAdminApplicationService struct {
	ContactAdminDatasource *datasource.ContactAdminDatasource
}

func NewContactAdminApplicationService(ContactAdminDatasource *datasource.ContactAdminDatasource) *ContactAdminApplicationService {
	return &ContactAdminApplicationService{ContactAdminDatasource: ContactAdminDatasource}
}

type ContactAdminApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	ContactID int
}

func (svc *ContactAdminApplicationService) FindByID(params ContactAdminApplicationServiceFindByIDParams) (*response.ContactAdminResponse, error) {
	contact, err := svc.ContactAdminDatasource.FindByID(datasource.ContactAdminDatasourceFindByIDParams{
		Ctx:       params.Ctx,
		ContactID: params.ContactID,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewContactAdminResponse(response.NewContactAdminResponseParams{
		Contact: contact,
	})

	return dto, nil
}

type ContactAdminApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request.ContactAdminFindAllRequest
}

func (svc *ContactAdminApplicationService) FindAll(params ContactAdminApplicationServiceFindAllParams) (*response.Collection[*response.ContactAdminResponse], error) {
	ContactList, err := svc.ContactAdminDatasource.FindAll(datasource.ContactAdminDatasourceFindAllParams{
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

	ContactCount, err := svc.ContactAdminDatasource.GetContactCount(datasource.ContactAdminDatasourceGetContactCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var ContactAdminResponseList []*response.ContactAdminResponse
	for _, Contact_i := range ContactList {
		ContactAdminResponseList = append(ContactAdminResponseList, response.NewContactAdminResponse(response.NewContactAdminResponseParams{
			Contact: Contact_i,
		}))
	}

	dto := response.NewCollection(response.NewCollectionParams[*response.ContactAdminResponse]{
		Data:       ContactAdminResponseList,
		TotalCount: ContactCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type ContactAdminApplicationServiceUpdateParams struct {
	Ctx       context.Context
	ContactID int
	Body      request.ContactAdminUpdateRequest
}

func (svc *ContactAdminApplicationService) Update(params ContactAdminApplicationServiceUpdateParams) (*response.ContactAdminResponse, error) {
	Contact := model.NewContact(model.NewContactParams{
		EntContact: &ent.Contact{
			ID: params.ContactID,
		},
		OptionList: []func(*model.Contact){
			model.WithContactStatus(model.ContactStatus(params.Body.Status)),
		},
	})

	Contact, err := svc.ContactAdminDatasource.Update(datasource.ContactAdminDatasourceUpdateParams{
		Ctx:     params.Ctx,
		Contact: Contact,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewContactAdminResponse(response.NewContactAdminResponseParams{
		Contact: Contact,
	})

	return dto, nil
}
