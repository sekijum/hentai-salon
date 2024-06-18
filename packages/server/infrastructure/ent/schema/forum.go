package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// Forum holds the schema definition for the Forum entity.
type Forum struct {
    ent.Schema
}

// Fields of the Forum.
func (Forum) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("user_id"),
        field.String("name"),
        field.String("description").Optional(),
        field.Enum("status").Values("Public", "Private", "Archived", "Disapproved").Default("Public"),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the Forum.
func (Forum) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("forums").Unique().Field("user_id").Required(),
        edge.To("topics", Topic.Type),
        edge.To("forum_likes", ForumLike.Type),
    }
}
