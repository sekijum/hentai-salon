package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// CommentAttachment holds the schema definition for the CommentAttachment entity.
type CommentAttachment struct {
    ent.Schema
}

// Fields of the CommentAttachment.
func (CommentAttachment) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("comment_id"),
        field.String("path"),
        field.Enum("type").Values("image", "video"),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the CommentAttachment.
func (CommentAttachment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("comment", Comment.Type).Ref("comment_attachments").Field("comment_id").Unique().Required(),
    }
}
