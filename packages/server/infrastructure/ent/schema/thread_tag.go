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
			field.ID("threadId", "tagId"),
	}
}

func (ThreadTag) Fields() []ent.Field {
	return []ent.Field{
			field.Int("threadId").StorageKey("thread_id"),
			field.Int("tagId").StorageKey("tag_id"),
	}
}

func (ThreadTag) Edges() []ent.Edge {
	return []ent.Edge{
			edge.To("thread", Thread.Type).Unique().Required().Field("threadId"),
			edge.To("tag", Tag.Type).Unique().Required().Field("tagId"),
	}
}
