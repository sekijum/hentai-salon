package model

import (
	"server/infrastructure/ent"
)

type Contact struct {
	EntContact *ent.Contact
}

type NewContactParams struct {
	EntContact *ent.Contact
	OptionList []func(*Contact)
}

func NewContact(params NewContactParams) *Contact {
	contact := &Contact{EntContact: params.EntContact}

	for _, option := range params.OptionList {
		option(contact)
	}

	return contact
}

func WithContactStatus(status ContactStatus) func(*Contact) {
	return func(c *Contact) {
		c.EntContact.Status = int(status)
	}
}

type ContactStatus int

const (
	ContactStatusOpen ContactStatus = iota
	ContactStatusPending
	ContactStatusClosed
)

func (c *Contact) StatusToString() string {
	switch ContactStatus(c.EntContact.Status) {
	case ContactStatusOpen:
		return "Open"
	case ContactStatusPending:
		return "Pending"
	case ContactStatusClosed:
		return "Closed"
	default:
		return "Unknown"
	}
}

func (c *Contact) StatusToLabel() string {
	switch ContactStatus(c.EntContact.Status) {
	case ContactStatusOpen:
		return "未対応"
	case ContactStatusPending:
		return "対応中"
	case ContactStatusClosed:
		return "完了"
	default:
		return "不明なステータス"
	}
}
