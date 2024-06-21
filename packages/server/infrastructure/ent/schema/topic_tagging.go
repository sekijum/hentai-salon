package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
		"entgo.io/ent/schema"
)

type TopicTagging struct {
	ent.Schema
}

func (TopicTagging) Annotations() []schema.Annotation {
	return []schema.Annotation{
			field.ID("topicId", "tagId"),
	}
}

func (TopicTagging) Fields() []ent.Field {
	return []ent.Field{
			field.Int("topicId").StorageKey("topic_id"),
			field.Int("tagId").StorageKey("tag_id"),
	}
}

func (TopicTagging) Edges() []ent.Edge {
	return []ent.Edge{
			edge.To("topic", Topic.Type).Unique().Required().Field("topicId"),
			edge.To("tag", TopicTag.Type).Unique().Required().Field("tagId"),
	}
}
