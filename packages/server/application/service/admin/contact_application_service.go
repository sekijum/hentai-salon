package service_admin

import (
	"context"
	"server/domain/model"
	datasource_admin "server/infrastructure/datasource/admin"
	"server/infrastructure/ent"
	request_admin "server/presentation/request/admin"
	response_admin "server/presentation/response/admin"
)

type ContactApplicationService struct {
	ContactDatasource *datasource_admin.ContactDatasource
}

func NewContactApplicationService(ContactDatasource *datasource_admin.ContactDatasource) *ContactApplicationService {
	return &ContactApplicationService{ContactDatasource: ContactDatasource}
}

type ContactApplicationServiceFindByIDParams struct {
	Ctx       context.Context
	ContactID int
}

func (svc *ContactApplicationService) FindByID(params ContactApplicationServiceFindByIDParams) (*response_admin.ContactResponse, error) {
	contact, err := svc.ContactDatasource.FindByID(datasource_admin.ContactDatasourceFindByIDParams{
		Ctx:       params.Ctx,
		ContactID: params.ContactID,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewContactResponse(response_admin.NewContactResponseParams{
		Contact: contact,
	})

	return dto, nil
}

type ContactApplicationServiceFindAllParams struct {
	Ctx context.Context
	Qs  request_admin.ContactFindAllRequest
}

func (svc *ContactApplicationService) FindAll(params ContactApplicationServiceFindAllParams) (*response_admin.Collection[*response_admin.ContactResponse], error) {
	ContactList, err := svc.ContactDatasource.FindAll(datasource_admin.ContactDatasourceFindAllParams{
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

	ContactCount, err := svc.ContactDatasource.GetContactCount(datasource_admin.ContactDatasourceGetContactCountParams{
		Ctx:     params.Ctx,
		Keyword: params.Qs.Keyword,
		Status:  params.Qs.Status,
	})
	if err != nil {
		return nil, err
	}

	var ContactResponseList []*response_admin.ContactResponse
	for _, Contact_i := range ContactList {
		ContactResponseList = append(ContactResponseList, response_admin.NewContactResponse(response_admin.NewContactResponseParams{
			Contact: Contact_i,
		}))
	}

	dto := response_admin.NewCollection(response_admin.NewCollectionParams[*response_admin.ContactResponse]{
		Data:       ContactResponseList,
		TotalCount: ContactCount,
		Limit:      params.Qs.Limit,
		Offset:     params.Qs.Offset,
	})

	return dto, nil
}

type ContactApplicationServiceUpdateParams struct {
	Ctx       context.Context
	ContactID int
	Body      request_admin.ContactUpdateRequest
}

func (svc *ContactApplicationService) Update(params ContactApplicationServiceUpdateParams) (*response_admin.ContactResponse, error) {
	Contact := model.NewContact(model.NewContactParams{
		EntContact: &ent.Contact{
			ID: params.ContactID,
		},
		OptionList: []func(*model.Contact){
			model.WithContactStatus(model.ContactStatus(params.Body.Status)),
		},
	})

	Contact, err := svc.ContactDatasource.Update(datasource_admin.ContactDatasourceUpdateParams{
		Ctx:     params.Ctx,
		Contact: Contact,
	})
	if err != nil {
		return nil, err
	}

	dto := response_admin.NewContactResponse(response_admin.NewContactResponseParams{
		Contact: Contact,
	})

	return dto, nil
}
