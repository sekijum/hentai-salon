package datasource

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/contact"
	"time"
)

type ContactAdminDatasource struct {
	client *ent.Client
}

func NewContactAdminDatasource(client *ent.Client) *ContactAdminDatasource {
	return &ContactAdminDatasource{client: client}
}

type ContactAdminDatasourceGetContactCountParams struct {
	Ctx     context.Context
	Keyword *string
	Status  *int
}

func (ds *ContactAdminDatasource) GetContactCount(params ContactAdminDatasourceGetContactCountParams) (int, error) {
	query := ds.client.Contact.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(contact.Or(
			contact.SubjectContains(*params.Keyword),
			contact.MessageContains(*params.Keyword),
		))
	}

	if params.Status != nil && *params.Status != 0 {
		query = query.Where(contact.StatusEQ(*params.Status))
	}

	ContactCount, err := query.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return ContactCount, nil
}

type ContactAdminDatasourceFindByIDParams struct {
	Ctx       context.Context
	ContactID int
}

func (ds *ContactAdminDatasource) FindByID(params ContactAdminDatasourceFindByIDParams) (*model.Contact, error) {
	entContact, err := ds.client.Contact.Get(params.Ctx, params.ContactID)
	if err != nil {
		return nil, err
	}

	Contact := model.NewContact(model.NewContactParams{EntContact: entContact})

	return Contact, nil
}

type ContactAdminDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
	Status  *int
}

func (ds *ContactAdminDatasource) FindAll(params ContactAdminDatasourceFindAllParams) ([]*model.Contact, error) {
	query := ds.client.Contact.Query()

	sort := contact.FieldID
	if params.Sort != nil && *params.Sort != "" {
		sort = *params.Sort
	}

	if params.Order != nil && *params.Order == "asc" {
		query = query.Order(ent.Asc(sort))
	} else {
		query = query.Order(ent.Desc(sort))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		query = query.Where(contact.Or(
			contact.SubjectContains(*params.Keyword),
			contact.MessageContains(*params.Keyword),
		))
	}

	if params.Status != nil && *params.Status != 0 {
		query = query.Where(contact.StatusEQ(*params.Status))
	}

	query = query.Limit(params.Limit)
	query = query.Offset(params.Offset)

	entContacts, err := query.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelContacts []*model.Contact
	for _, entContact := range entContacts {
		modelContacts = append(modelContacts, model.NewContact(model.NewContactParams{EntContact: entContact}))
	}

	return modelContacts, nil
}

type ContactAdminDatasourceUpdateParams struct {
	Ctx     context.Context
	Contact *model.Contact
}

func (ds *ContactAdminDatasource) Update(params ContactAdminDatasourceUpdateParams) (*model.Contact, error) {
	update := ds.client.Contact.UpdateOneID(params.Contact.EntContact.ID)

	if params.Contact.EntContact.Status != 0 {
		update = update.SetStatus(params.Contact.EntContact.Status)
	}
	update = update.SetUpdatedAt(time.Now())

	entContact, err := update.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	Contact := model.NewContact(model.NewContactParams{
		EntContact: entContact,
	})

	return Contact, nil
}
