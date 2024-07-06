package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ThreadTag struct {
	ent.Schema
}

func (ThreadTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
			field.ID("thread_id", "tag_id"),
	}
}

func (ThreadTag) Fields() []ent.Field {
	return []ent.Field{
			field.Int("thread_id"),
			field.Int("tag_id"),
	}
}

func (ThreadTag) Edges() []ent.Edge {
	return []ent.Edge{
			edge.To("thread", Thread.Type).Unique().Required().Field("thread_id"),
			edge.To("tag", Tag.Type).Unique().Required().Field("tag_id"),
	}
}
