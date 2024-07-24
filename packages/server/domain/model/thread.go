package model

import (
	"server/infrastructure/ent"
)

type Thread struct {
	EntThread *ent.Thread
}

type ThreadStatus int

const (
	ThreadStatusOpen ThreadStatus = iota
	ThreadStatusPending
	ThreadStatusArchived
)

func (m *Thread) StatusToString() string {
	switch ThreadStatus(m.EntThread.Status) {
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

func (m *Thread) StatusToLabel() string {
	switch ThreadStatus(m.EntThread.Status) {
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
