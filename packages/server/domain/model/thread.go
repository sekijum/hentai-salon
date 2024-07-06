package model

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"time"
)

type Thread struct {
	Id                int
	BoardId           int
	UserId            int
	Title             string
	Description       *string
	ThumbnailUrl      *string
	IpAddress         string
	Status            ThreadStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time

	User  *User
	Board *Board
	Tags              []*Tag
}

func (t *Thread) Validate() error {

	if t.UserId == 0 {
		return errors.New("ユーザーIDは必須です")
	}
	if t.Title == "" {
		return errors.New("スレッドのタイトルは必須です")
	}
	if len(t.Title) > 50 {
		return errors.New("スレッドのタイトルは50文字以内である必要があります")
	}
	if t.Description != nil && len(*t.Description) > 255 {
		return errors.New("説明は255文字以内で入力してください")
	}
	if t.ThumbnailUrl != nil {
		if _, err := url.ParseRequestURI(*t.ThumbnailUrl); err != nil {
			return errors.New("サムネイルURLは有効なURLである必要があります")
		}
	}
	if t.IpAddress == "" {
		return errors.New("スレッドのIPアドレスは必須です")
	}
	if net.ParseIP(t.IpAddress) == nil {
		return errors.New("スレッドのIPアドレスは有効な形式である必要があります")
	}
	if err := t.Status.Validate(); err != nil {
		return err
	}

	return nil
}

type ThreadStatus int

const (
	ThreadStatusOpen ThreadStatus = iota
	ThreadStatusPending
	ThreadStatusArchived
)

func (s ThreadStatus) String() string {
	switch s {
	case ThreadStatusOpen:
		return "Open"
	case ThreadStatusPending:
		return "Pending"
	case ThreadStatusArchived:
		return "Archived"
	default:
		return "Unknown"
	}
}

func (s ThreadStatus) Validate() error {
	switch s {
	case ThreadStatusOpen, ThreadStatusArchived, ThreadStatusPending:
		return nil
	default:
		return errors.New("無効なスレッドステータスです")
	}
}

func (s ThreadStatus) ToInt() int {
	boardStatusToInt := map[ThreadStatus]int{
		ThreadStatusOpen:     0,
		ThreadStatusPending:  1,
		ThreadStatusArchived: 2,
	}
	return boardStatusToInt[s]
}

func (s ThreadStatus) Label() string {
	switch s {
	case ThreadStatusOpen:
		return "公開"
	case ThreadStatusPending:
		return "保留"
	case ThreadStatusArchived:
		return "アーカイブ"
	default:
		return "不明なステータス"
	}
}

func (t *Thread) GenerateDefaultTitle() {
	t.Title = fmt.Sprintf("%d", time.Now().Unix())
}
