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
		field.ID("userId", "commentId"),
	}
}

func (UserCommentSubscription) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("commentId").StorageKey("comment_id"),
        field.Bool("isNotified").Default(true).StorageKey("is_notified").Comment("メール通知を受け取るかどうかのフラグ"),
        field.Bool("isChecked").Default(false).StorageKey("is_checked").Comment("通知画面で確認したかどうかのフラグ"),
        field.Time("subscribedAt").Default(time.Now).StorageKey("subscribed_at"),
    }
}

func (UserCommentSubscription) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("comment", ThreadComment.Type).Unique().Required().Field("commentId"),
    }
}
