package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Board struct {
	ent.Schema
}

func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("user_id"),
		field.String("title").Unique().MaxLen(50).Comment("板名"),
		field.String("description").Optional().MaxLen(255),
		field.String("thumbnail_url").Optional(),
		field.Int("status").Default(0).Comment("0: Public, 1: Private, 3: Pending, 3: Archived"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Board) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("boards").Unique().Field("user_id").Required(),
		edge.To("threads", Thread.Type),
	}
}

func (Board) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique(),
	}
}
