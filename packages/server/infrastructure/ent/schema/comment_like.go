package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// CommentLike holds the schema definition for the CommentLike entity.
type CommentLike struct {
    ent.Schema
}

// Fields of the CommentLike.
func (CommentLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("comment_id"),
        field.Int("user_id"),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the CommentLike.
func (CommentLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("comment", Comment.Type).Ref("comment_likes").Unique().Field("comment_id").Required(),
        edge.From("user", User.Type).Ref("comment_likes").Unique().Field("user_id").Required(),
    }
}
