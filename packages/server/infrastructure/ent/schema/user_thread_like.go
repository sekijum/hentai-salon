package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema"
)

type UserThreadLike struct {
    ent.Schema
}

func (UserThreadLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("userId", "threadId"),
	}
}

func (UserThreadLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("threadId").StorageKey("thread_id"),
        field.Time("likedAt").Default(time.Now).StorageKey("liked_at"),
    }
}

func (UserThreadLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("thread", Thread.Type).Unique().Required().Field("threadId"),
    }
}
