package model

import (
	"server/infrastructure/ent"
)

type Tag struct {
	EntTag *ent.Tag
}

func (t *Tag) Name() string {
	return t.EntTag.Name
}
