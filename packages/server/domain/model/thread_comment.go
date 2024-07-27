package model

import (
	"server/infrastructure/ent"
)

type ThreadComment struct {
	EntThreadComment *ent.ThreadComment
}

type NewThreadCommentParams struct {
	EntThreadComment *ent.ThreadComment
	OptionList       []func(*ThreadComment)
}

func NewThreadComment(params NewThreadCommentParams) *ThreadComment {
	threadComment := &ThreadComment{EntThreadComment: params.EntThreadComment}

	for _, option_i := range params.OptionList {
		option_i(threadComment)
	}

	return threadComment
}

type ThreadCommentStatus int

const (
	ThreadCommentStatusVisible ThreadCommentStatus = iota
	ThreadCommentStatusDeleted
)

func WithThreadCommentStatus(status ThreadCommentStatus) func(*ThreadComment) {
	return func(tc *ThreadComment) {
		tc.EntThreadComment.Status = int(status)
	}
}

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
