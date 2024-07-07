package model

import (
	"math"
	"server/infrastructure/ent"
	"time"
)

type Thread struct {
	EntThread     *ent.Thread
	Popularity    int
	TotalComments int
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

func (m *Thread) CalculatePopularity() {
	var totalPopularity float64
	currentTime := time.Now()

	for _, comment := range m.EntThread.Edges.Comments {
		timeDiff := currentTime.Sub(comment.CreatedAt).Hours() + 1
		totalPopularity += 1 / timeDiff
	}

	if m.EntThread.Title != "" {
		totalPopularity += 10
	}
	if m.EntThread.Description != "" {
		totalPopularity += 5
	}

	scaledPopularity := int(math.Min(totalPopularity, 100))
	m.Popularity = scaledPopularity
}
