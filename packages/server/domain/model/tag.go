package model

import (
	"server/infrastructure/ent"
)

type Tag struct {
	EntTag *ent.Tag
}

type NewTagParams struct {
	EntTag     *ent.Tag
	OptionList []func(*Tag)
}

func NewTag(params NewTagParams) *Tag {
	tag := &Tag{EntTag: params.EntTag}

	for _, option_i := range params.OptionList {
		option_i(tag)
	}

	return tag
}
