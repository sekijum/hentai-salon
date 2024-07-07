package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name").MaxLen(20),
		field.String("email").Unique().MaxLen(254),
		field.String("password"),
		field.String("avatar_url").Optional().Nillable(),
		field.Int("status").Default(0).Comment("0: Active, 1: Withdrawn, 2: Suspended, 2: Inactive"),
		field.Int("role").Default(0).Comment("0: Member, 1: Admin"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("boards", Board.Type),
		edge.To("threads", Thread.Type),
		edge.To("comments", ThreadComment.Type),
		edge.To("liked_threads", Thread.Type).Through("user_thread_like", UserThreadLike.Type),
		edge.To("liked_comments", ThreadComment.Type).Through("user_comment_like", UserCommentLike.Type),
		edge.To("subscribed_threads", Thread.Type).Through("user_thread_subscription", UserThreadSubscription.Type),
		edge.To("subscribed_comments", ThreadComment.Type).Through("user_comment_subscription", UserCommentSubscription.Type),
	}
}
