package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "time"
)

// AdminUser holds the schema definition for the AdminUser entity.
type AdminUser struct {
    ent.Schema
}

// Fields of the AdminUser.
func (AdminUser) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.String("email").Unique(),
        field.String("password"),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the AdminUser.
func (AdminUser) Edges() []ent.Edge {
    return []ent.Edge{}
}
