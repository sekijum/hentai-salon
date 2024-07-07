package model

import (
	"server/infrastructure/ent"
)

type Board struct {
	EntBoard *ent.Board
}

type BoardStatus int

const (
	BoardStatusPublic BoardStatus = iota
	BoardStatusPrivate
	BoardStatusPending
	BoardStatusArchived
)

func (m Board) StatusToString() string {
	switch m.EntBoard.Status {
	case int(BoardStatusPublic):
		return "Public"
	case int(BoardStatusPrivate):
		return "Private"
	case int(BoardStatusPending):
		return "Pending"
	case int(BoardStatusArchived):
		return "Archived"
	default:
		return "Unknown"
	}
}

func (m Board) StatusToLabel() string {
	switch m.EntBoard.Status {
	case int(BoardStatusPublic):
		return "公開"
	case int(BoardStatusPrivate):
		return "非公開"
	case int(BoardStatusPending):
		return "保留"
	case int(BoardStatusArchived):
		return "アーカイブ"
	default:
		return "不明なステータス"
	}
}
