package model

import (
	"server/infrastructure/ent"
)

type ThreadCommentAttachment struct {
	EntAttachment *ent.ThreadCommentAttachment
}

type AttachmentType int

const (
	AttachmentTypeImage AttachmentType = iota
	AttachmentTypeVideo
)

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

func (m ThreadCommentAttachment) TypeToInt() int {
	return m.EntAttachment.Type
}
