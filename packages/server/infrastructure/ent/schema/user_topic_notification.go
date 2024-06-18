package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// UserTopicNotification holds the schema definition for the UserTopicNotification entity.
type UserTopicNotification struct {
    ent.Schema
}

// Fields of the UserTopicNotification.
func (UserTopicNotification) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("user_id"),
        field.Int("topic_id"),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the UserTopicNotification.
func (UserTopicNotification) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("user_topic_notifications").Unique().Field("user_id").Required(),
        edge.From("topic", Topic.Type).Ref("user_topic_notifications").Unique().Field("topic_id").Required(),
    }
}
