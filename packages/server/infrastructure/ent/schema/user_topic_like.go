package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema"
)

type UserTopicLike struct {
    ent.Schema
}

func (UserTopicLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("userId", "topicId"),
	}
}

func (UserTopicLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("topicId").StorageKey("topic_id"),
        field.Enum("type").Values("like", "dislike"),
        field.Time("likedAt").Default(time.Now).StorageKey("liked_at"),
    }
}

func (UserTopicLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("topic", Topic.Type).Unique().Required().Field("topicId"),
    }
}
