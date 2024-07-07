package model

import (
	"server/infrastructure/ent"
)

type ThreadComment struct {
	EntThreadComment *ent.ThreadComment
}

type ThreadCommentStatus int

const (
	ThreadCommentStatusVisible ThreadCommentStatus = iota
	ThreadCommentStatusDeleted
)

func (m *ThreadComment) StatusToString() string {
	switch ThreadCommentStatus(m.EntThreadComment.Status) {
	case ThreadCommentStatusVisible:
		return "Visible"
	case ThreadCommentStatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

func (m *ThreadComment) ThreadCommentToLabel() string {
	switch ThreadCommentStatus(m.EntThreadComment.Status) {
	case ThreadCommentStatusVisible:
		return "可視"
	case ThreadCommentStatusDeleted:
		return "削除"
	default:
		return "不明なステータス"
	}
}

func (m *ThreadComment) TotalReplies() int {
	return len(m.EntThreadComment.Edges.Replies)
}
