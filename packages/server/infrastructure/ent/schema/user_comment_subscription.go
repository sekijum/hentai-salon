package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema"
)

type UserCommentSubscription struct {
    ent.Schema
}

func (UserCommentSubscription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("userId", "commentId"),
	}
}

func (UserCommentSubscription) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("commentId").StorageKey("comment_id"),
        field.Bool("isNotified").Default(true).StorageKey("is_notified").Comment("通知を受け取るかどうかのカラムを追加"),
        field.Time("subscribedAt").Default(time.Now).StorageKey("subscribed_at"),
    }
}

func (UserCommentSubscription) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("comment", Comment.Type).Unique().Required().Field("commentId"),
    }
}
