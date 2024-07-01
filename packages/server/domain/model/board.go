package model

import (
	"errors"
	"net/url"
	"time"
)

type Board struct {
	Id           int
	UserId       int
	Title        string
	Description  *string
	ThumbnailUrl *string
	Status       BoardStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User    *User
	Threads []*Thread
}

func (b *Board) Validate() error {
	if b.UserId == 0 {
		return errors.New("ユーザーIDは必須です")
	}
	if b.Title == "" {
		return errors.New("タイトルは必須です")
	}
	if len(b.Title) > 50 {
		return errors.New("タイトルは50文字以内で入力してください")
	}
	if b.Description != nil && len(*b.Description) > 255 {
		return errors.New("説明は255文字以内で入力してください")
	}
	if b.ThumbnailUrl != nil {
		if _, err := url.ParseRequestURI(*b.ThumbnailUrl); err != nil {
			return errors.New("サムネイルURLは有効なURLである必要があります")
		}
	}
	if err := b.Status.Validate(); err != nil {
		return err
	}
	return nil
}

type BoardStatus int

const (
	BoardStatusPublic BoardStatus = iota
	BoardStatusPrivate
	BoardStatusPending
	BoardStatusArchived
)

func (s BoardStatus) String() string {
	switch s {
	case BoardStatusPublic:
		return "Public"
	case BoardStatusPrivate:
		return "Private"
	case BoardStatusPending:
		return "Pending"
	case BoardStatusArchived:
		return "Archived"
	default:
		return "Unknown"
	}
}

func (s BoardStatus) Validate() error {
	switch s {
	case BoardStatusPublic, BoardStatusPrivate, BoardStatusPending, BoardStatusArchived:
		return nil
	default:
		return errors.New("無効な板ステータスです")
	}
}

func (b BoardStatus) ToInt() int {
	BoardStatusToInt := map[BoardStatus]int{
		BoardStatusPublic:   0,
		BoardStatusPrivate:  1,
		BoardStatusPending:  2,
		BoardStatusArchived: 3,
	}
	return BoardStatusToInt[b]
}

func (s BoardStatus) Label() string {
	switch s {
	case BoardStatusPublic:
		return "公開"
	case BoardStatusPrivate:
		return "非公開"
	case BoardStatusPending:
		return "保留"
	case BoardStatusArchived:
		return "アーカイブ"
	default:
		return "不明なステータス"
	}
}
