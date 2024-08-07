package datasource_admin

import (
	"context"
	"server/domain/model"
	"server/infrastructure/ent"
	"server/infrastructure/ent/contact"
	"strconv"
	"time"
)

type ContactDatasource struct {
	client *ent.Client
}

func NewContactDatasource(client *ent.Client) *ContactDatasource {
	return &ContactDatasource{client: client}
}

type ContactDatasourceGetContactCountParams struct {
	Ctx     context.Context
	Keyword *string
	Status  *int
}

func (ds *ContactDatasource) GetContactCount(params ContactDatasourceGetContactCountParams) (int, error) {
	q := ds.client.Contact.Query()

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(contact.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(contact.IDEQ(id))
			}
		default:
			q = q.Where(contact.Or(
				contact.SubjectContainsFold(*params.Keyword),
				contact.MessageContainsFold(*params.Keyword),
			))
		}
	}

	ContactCount, err := q.Count(params.Ctx)
	if err != nil {
		return 0, err
	}
	return ContactCount, nil
}

type ContactDatasourceFindByIDParams struct {
	Ctx       context.Context
	ContactID int
}

func (ds *ContactDatasource) FindByID(params ContactDatasourceFindByIDParams) (*model.Contact, error) {
	entContact, err := ds.client.Contact.Get(params.Ctx, params.ContactID)
	if err != nil {
		return nil, err
	}

	Contact := model.NewContact(model.NewContactParams{EntContact: entContact})

	return Contact, nil
}

type ContactDatasourceFindAllParams struct {
	Ctx     context.Context
	Limit   int
	Offset  int
	Sort    *string
	Order   *string
	Keyword *string
}

func (ds *ContactDatasource) FindAll(params ContactDatasourceFindAllParams) ([]*model.Contact, error) {
	q := ds.client.Contact.Query()

	sort := contact.FieldID
	order := "desc"

	if params.Sort != nil && *params.Sort != "" {
		sort = *params.Sort
	}
	if params.Order != nil && *params.Order != "" {
		order = *params.Order
	}

	if order == "asc" {
		q = q.Order(ent.Asc(sort))
	} else {
		q = q.Order(ent.Desc(sort))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		switch {
		case len(*params.Keyword) > 7 && (*params.Keyword)[:7] == "status:":
			if status, err := strconv.Atoi((*params.Keyword)[7:]); err == nil {
				q = q.Where(contact.StatusEQ(status))
			}
		case len(*params.Keyword) > 3 && (*params.Keyword)[:3] == "id:":
			if id, err := strconv.Atoi((*params.Keyword)[3:]); err == nil {
				q = q.Where(contact.IDEQ(id))
			}
		default:
			q = q.Where(contact.Or(
				contact.SubjectContainsFold(*params.Keyword),
				contact.MessageContainsFold(*params.Keyword),
				contact.EmailContainsFold(*params.Keyword),
			))
		}
	}

	entContactList, err := q.All(params.Ctx)
	if err != nil {
		return nil, err
	}

	var modelContacts []*model.Contact
	for _, entContact := range entContactList {
		modelContacts = append(modelContacts, model.NewContact(model.NewContactParams{EntContact: entContact}))
	}

	return modelContacts, nil
}

type ContactDatasourceUpdateStatusParams struct {
	Ctx     context.Context
	Contact *model.Contact
}

func (ds *ContactDatasource) UpdateStatus(params ContactDatasourceUpdateStatusParams) (*model.Contact, error) {
	q := ds.client.Contact.UpdateOneID(params.Contact.EntContact.ID).
		SetStatus(params.Contact.EntContact.Status)

	q = q.SetUpdatedAt(time.Now())

	entContact, err := q.Save(params.Ctx)
	if err != nil {
		return nil, err
	}

	Contact := model.NewContact(model.NewContactParams{
		EntContact: entContact,
	})

	return Contact, nil
}
