package resource

import (
	"server/domain/model"
	"time"
)

type ContactResource struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	IPAddress string    `json:"ip_address"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewContactResource(contact *model.Contact) *ContactResource {
	return &ContactResource{
		ID:        contact.EntContact.ID,
		Email:     contact.EntContact.Email,
		Subject:   contact.EntContact.Subject,
		Message:   contact.EntContact.Message,
		IPAddress: contact.EntContact.IPAddress,
		Status:    contact.StatusToString(),
		CreatedAt: contact.EntContact.CreatedAt,
		UpdatedAt: contact.EntContact.UpdatedAt,
	}
}
