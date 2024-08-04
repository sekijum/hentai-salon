package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Contact struct {
	ent.Schema
}

func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("email").MaxLen(255),
		field.String("subject").MaxLen(255),
		field.Text("message"),
		field.String("ip_address").MaxLen(64),
		field.Int("status").Default(0).Comment("0: Open, 1: Pending, 2: Closed"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}
