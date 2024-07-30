package response

import (
	"server/domain/model"
	"time"
)

type ContactAdminResponse struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
	Status      int    `json:"status"`
	StatusLabel string `json:"statusLabel"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type NewContactAdminResponseParams struct {
	Contact *model.Contact
}

func NewContactAdminResponse(params NewContactAdminResponseParams) *ContactAdminResponse {
	return &ContactAdminResponse{
		ID:          params.Contact.EntContact.ID,
		Email:       params.Contact.EntContact.Email,
		Subject:     params.Contact.EntContact.Subject,
		Message:     params.Contact.EntContact.Message,
		Status:      params.Contact.EntContact.Status,
		StatusLabel: params.Contact.StatusToLabel(),
		CreatedAt:   params.Contact.EntContact.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   params.Contact.EntContact.UpdatedAt.Format(time.RFC3339),
	}
}
