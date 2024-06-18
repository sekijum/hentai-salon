package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
    ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("topic_id"),
        field.Int("user_id"),
        field.Int("parent_id").Optional(),
        field.Text("content"),
        field.Enum("status").Values("Visible", "Hidden", "Deleted", "Disapproved").Default("Visible"),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("topic", Topic.Type).Ref("comments").Unique().Field("topic_id").Required(),
        edge.From("user", User.Type).Ref("comments").Unique().Field("user_id").Required(),
        edge.From("parent", Comment.Type).Ref("replies").Unique().Field("parent_id"),
        edge.To("replies", Comment.Type),
        edge.To("comment_likes", CommentLike.Type),
        edge.To("comment_attachments", CommentAttachment.Type),
        edge.To("user_comment_notifications", UserCommentNotification.Type),
    }
}
