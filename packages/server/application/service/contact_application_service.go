package service

import (
	"context"
	"server/domain/model"
	"server/infrastructure/datasource"
	"server/infrastructure/ent"
	"server/presentation/request"
	"server/presentation/response"
)

type ContactApplicationService struct {
	contactDatasource *datasource.ContactDatasource
}

func NewContactApplicationService(contactDatasource *datasource.ContactDatasource) *ContactApplicationService {
	return &ContactApplicationService{contactDatasource: contactDatasource}
}

type ContactApplicationServiceCreateParams struct {
	Ctx      context.Context
	ClientIP string
	Body     request.ContactCreateRequest
}

func (svc *ContactApplicationService) Create(params ContactApplicationServiceCreateParams) (*response.ContactResponse, error) {
	contact := model.NewContact(model.NewContactParams{
		EntContact: &ent.Contact{
			Email:     &params.Body.Email,
			Subject:   params.Body.Subject,
			Message:   params.Body.Message,
			IPAddress: params.ClientIP,
		},
		OptionList: []func(*model.Contact){
			model.WithContactStatus(model.ContactStatusOpen),
		},
	})

	contact, err := svc.contactDatasource.Create(datasource.ContactDatasourceCreateParams{
		Ctx:     params.Ctx,
		Contact: contact,
	})
	if err != nil {
		return nil, err
	}

	dto := response.NewContactResponse(contact)

	return dto, nil
}
