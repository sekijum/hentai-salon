package response

import (
	"server/domain/model"
	"time"
)

type ContactResponse struct {
	ID        int       `json:"id"`
	Email     *string   `json:"email,omitempty"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	IPAddress string    `json:"ip_address"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewContactResponse(contact *model.Contact) *ContactResponse {
	return &ContactResponse{
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
