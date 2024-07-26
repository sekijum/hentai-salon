package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name").Unique().MaxLen(50),
		field.Time("created_at").Default(time.Now),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("threads", Thread.Type).Ref("tags").Through("thread_tags", ThreadTag.Type),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}
