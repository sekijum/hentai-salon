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
