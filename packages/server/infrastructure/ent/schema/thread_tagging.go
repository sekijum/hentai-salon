package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
		"entgo.io/ent/schema"
)

type ThreadTagging struct {
	ent.Schema
}

func (ThreadTagging) Annotations() []schema.Annotation {
	return []schema.Annotation{
			field.ID("threadId", "tagId"),
	}
}

func (ThreadTagging) Fields() []ent.Field {
	return []ent.Field{
			field.Int("threadId").StorageKey("thread_id"),
			field.Int("tagId").StorageKey("tag_id"),
	}
}

func (ThreadTagging) Edges() []ent.Edge {
	return []ent.Edge{
			edge.To("thread", Thread.Type).Unique().Required().Field("threadId"),
			edge.To("tag", ThreadTag.Type).Unique().Required().Field("tagId"),
	}
}
