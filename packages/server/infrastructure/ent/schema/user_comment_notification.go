package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// UserCommentNotification holds the schema definition for the UserCommentNotification entity.
type UserCommentNotification struct {
    ent.Schema
}

// Fields of the UserCommentNotification.
func (UserCommentNotification) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("user_id"),
        field.Int("comment_id"),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the UserCommentNotification.
func (UserCommentNotification) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("user_comment_notifications").Unique().Field("user_id").Required(),
        edge.From("comment", Comment.Type).Ref("user_comment_notifications").Unique().Field("comment_id").Required(),
    }
}
