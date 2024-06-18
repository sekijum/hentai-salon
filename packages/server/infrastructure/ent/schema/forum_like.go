package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// ForumLike holds the schema definition for the ForumLike entity.
type ForumLike struct {
    ent.Schema
}

// Fields of the ForumLike.
func (ForumLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("forum_id"),
        field.Int("user_id"),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the ForumLike.
func (ForumLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("forum", Forum.Type).Ref("forum_likes").Unique().Field("forum_id").Required(),
        edge.From("user", User.Type).Ref("forum_likes").Unique().Field("user_id").Required(),
    }
}
