package model

import (
	"errors"
	"time"
)

type AttachmentType int

const (
	AttachmentTypeImage AttachmentType = iota
	AttachmentTypeVideo
)

func (a AttachmentType) String() string {
	switch a {
	case AttachmentTypeImage:
		return "Image"
	case AttachmentTypeVideo:
		return "Video"
	default:
		return "Unknown"
	}
}

func (a AttachmentType) Validate() error {
	switch a {
	case AttachmentTypeImage, AttachmentTypeVideo:
		return nil
	default:
		return errors.New("無効な添付ファイルのタイプです")
	}
}

func (a AttachmentType) ToInt() int {
	AttachmentTypeToInt := map[AttachmentType]int{
		AttachmentTypeImage: 0,
		AttachmentTypeVideo: 1,
	}
	return AttachmentTypeToInt[a]
}

type ThreadCommentAttachment struct {
	Id          int
	CommentId   int
	Url         string
	DisplayOrder int
	Type        AttachmentType
	CreatedAt   time.Time
}

func (a *ThreadCommentAttachment) Validate() error {
	if a.CommentId == 0 {
		return errors.New("コメントIDは必須です")
	}
	if a.Url == "" {
		return errors.New("URLは必須です")
	}
	if err := a.Type.Validate(); err != nil {
		return err
	}
	return nil
}
