package model

import (
	"errors"
	"time"
)

type ThreadCommentStatus int

const (
	ThreadCommentStatusVisible ThreadCommentStatus = iota
	ThreadCommentStatusDeleted
)

func (s ThreadCommentStatus) String() string {
	switch s {
	case ThreadCommentStatusVisible:
		return "Visible"
	case ThreadCommentStatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

func (s ThreadCommentStatus) Validate() error {
	switch s {
	case ThreadCommentStatusVisible, ThreadCommentStatusDeleted:
		return nil
	default:
		return errors.New("無効なコメントステータスです")
	}
}

func (s ThreadCommentStatus) ToInt() int {
	ThreadCommentStatusToInt := map[ThreadCommentStatus]int{
		ThreadCommentStatusVisible: 0,
		ThreadCommentStatusDeleted: 1,
	}
	return ThreadCommentStatusToInt[s]
}

func (s ThreadCommentStatus) Label() string {
	switch s {
	case ThreadCommentStatusVisible:
		return "Visible"
	case ThreadCommentStatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

type ThreadComment struct {
	Id             int
	ThreadId       int
	ParentCommentId *int
	UserId         *int
	GuestName      *string
	Content        string
	IpAddress      string
	Status         ThreadCommentStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time

	ParentComment *ThreadComment
	Replies      []*ThreadComment
	Attachments  []*ThreadCommentAttachment
}

func (c *ThreadComment) Validate() error {
	if c.ThreadId == 0 {
		return errors.New("スレッドIDは必須です")
	}
	if c.Content == "" {
		return errors.New("コンテンツは必須です")
	}
	if len(c.Content) > 1000 {
		return errors.New("コンテンツは1000文字以内で入力してください")
	}
	if err := c.Status.Validate(); err != nil {
		return err
	}
	if len(c.IpAddress) > 64 {
		return errors.New("IPアドレスは64文字以内で入力してください")
	}
	return nil
}
