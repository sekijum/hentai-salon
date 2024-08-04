package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserThreadLike struct {
	ent.Schema
}

func (UserThreadLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "thread_id"),
	}
}

func (UserThreadLike) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("thread_id"),
		field.Time("liked_at").Default(time.Now),
	}
}

func (UserThreadLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.To("thread", Thread.Type).Unique().Required().Field("thread_id"),
	}
}
