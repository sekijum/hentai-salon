package model

import (
	"errors"
	"server/infrastructure/ent"
)

type AttachmentType int

const (
	AttachmentTypeImage AttachmentType = iota
	AttachmentTypeVideo
)

type ThreadCommentAttachment struct {
	EntAttachment *ent.ThreadCommentAttachment
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

func (m ThreadCommentAttachment) TypeToInt() int {
	return int(m.EntAttachment.Type)
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
