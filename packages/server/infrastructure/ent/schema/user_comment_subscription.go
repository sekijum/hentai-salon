package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserCommentSubscription struct {
	ent.Schema
}

func (UserCommentSubscription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "comment_id"),
	}
}

func (UserCommentSubscription) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("comment_id"),
		field.Bool("is_notified").Default(true).Comment("メール通知を受け取るかどうかのフラグ"),
		field.Bool("is_checked").Default(false).Comment("通知画面で確認したかどうかのフラグ"),
		field.Time("subscribed_at").Default(time.Now),
	}
}

func (UserCommentSubscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.To("comment", ThreadComment.Type).Unique().Required().Field("comment_id"),
	}
}
