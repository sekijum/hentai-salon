package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"time"
)

type ContactDatasource struct {
	client *ent.Client
}

func NewContactDatasource(client *ent.Client) *ContactDatasource {
	return &ContactDatasource{client: client}
}

type ContactDatasourceCreateParams struct {
	Ctx     context.Context
	Contact *model.Contact
}

func (ds *ContactDatasource) Create(params ContactDatasourceCreateParams) (*model.Contact, error) {
	q := ds.client.Contact.Create().
		SetEmail(params.Contact.EntContact.Email).
		SetSubject(params.Contact.EntContact.Subject).
		SetMessage(params.Contact.EntContact.Message).
		SetIPAddress(params.Contact.EntContact.IPAddress).
		SetStatus(params.Contact.EntContact.Status).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now())

	entContact, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	modelContact := model.NewContact(model.NewContactParams{EntContact: entContact})

	return modelContact, nil
}
