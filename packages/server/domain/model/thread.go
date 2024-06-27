package model

import (
	"errors"
	"fmt"
	"time"
)

type Thread struct {
	Id                int
	BoardId           int
	UserId            int
	Title             string
	Description       *string
	ThumbnailUrl      *string
	IsAutoGenerated   bool
	IsNotifyOnComment bool
	IpAddress         string
	Status            ThreadStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Board             *Board
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
		return "公開中"
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
