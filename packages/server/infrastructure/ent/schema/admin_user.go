package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type AdminUser struct {
	ent.Schema
}

func (AdminUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name").Unique().StorageKey("name").MaxLen(20),
		field.String("email").Unique().MaxLen(254),
		field.String("password"),
		field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).StorageKey("updated_at"),
	}
}
