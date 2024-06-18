package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// TopicLike holds the schema definition for the TopicLike entity.
type TopicLike struct {
    ent.Schema
}

// Fields of the TopicLike.
func (TopicLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("topic_id"),
        field.Int("user_id"),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the TopicLike.
func (TopicLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("topic", Topic.Type).Ref("topic_likes").Unique().Field("topic_id").Required(),
        edge.From("user", User.Type).Ref("topic_likes").Unique().Field("user_id").Required(),
    }
}
