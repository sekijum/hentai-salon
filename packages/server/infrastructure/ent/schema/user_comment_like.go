package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserCommentLike struct {
	ent.Schema
}

func (UserCommentLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "comment_id"),
	}
}

func (UserCommentLike) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("comment_id"),
		field.Time("liked_at").Default(time.Now),
	}
}

func (UserCommentLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.To("comment", ThreadComment.Type).Unique().Required().Field("comment_id"),
	}
}
