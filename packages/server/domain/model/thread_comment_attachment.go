package model

import (
	"errors"
	"server/infrastructure/ent"
)

type ThreadCommentAttachment struct {
	EntAttachment *ent.ThreadCommentAttachment
}

type NewThreadCommentAttachmentParams struct {
	EntAttachment *ent.ThreadCommentAttachment
	OptionList    []func(*ThreadCommentAttachment)
}

func NewThreadCommentAttachment(params NewThreadCommentAttachmentParams) *ThreadCommentAttachment {
	attachment := &ThreadCommentAttachment{EntAttachment: params.EntAttachment}

	for _, option_i := range params.OptionList {
		option_i(attachment)
	}

	return attachment
}

type AttachmentType int

const (
	AttachmentTypeImage AttachmentType = iota
	AttachmentTypeVideo
)

func WithAttachmentTypeFromString(s string) func(*ThreadCommentAttachment) {
	return func(a *ThreadCommentAttachment) {
		attachmentType, err := AttachmentTypeFromString(s)
		if err != nil {
			a.EntAttachment.Type = -1 // Invalid type
		} else {
			a.EntAttachment.Type = int(attachmentType)
		}
	}
}

func (m ThreadCommentAttachment) TypeToString() string {
	switch AttachmentType(m.EntAttachment.Type) {
	case AttachmentTypeImage:
		return "Image"
	case AttachmentTypeVideo:
		return "Video"
	default:
		return "Unknown"
	}
}

func AttachmentTypeFromString(s string) (AttachmentType, error) {
	switch s {
	case "Image":
		return AttachmentTypeImage, nil
	case "Video":
		return AttachmentTypeVideo, nil
	default:
		return -1, errors.New("invalid attachment type")
	}
}
