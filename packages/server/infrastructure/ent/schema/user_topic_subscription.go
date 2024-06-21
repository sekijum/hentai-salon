package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema"
)

type UserTopicSubscription struct {
    ent.Schema
}

func (UserTopicSubscription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("userId", "topicId"),
	}
}

func (UserTopicSubscription) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("topicId").StorageKey("topic_id"),
        field.Bool("isNotified").Default(true).StorageKey("is_notified").Comment("通知を受け取るかどうかのカラムを追加"),
        field.Time("subscribedAt").Default(time.Now).StorageKey("subscribed_at"),
    }
}

func (UserTopicSubscription) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("topic", Topic.Type).Unique().Required().Field("topicId"),
    }
}
