package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/index"
)

type ThreadTag struct {
    ent.Schema
}

func (ThreadTag) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable(),
        field.String("name").Unique().MaxLen(20),
        field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
    }
}

func (ThreadTag) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("threads", Thread.Type).Ref("tags").Through("thread_taggings", ThreadTagging.Type),
    }
}

func (ThreadTag) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("name").Unique(),
    }
}
