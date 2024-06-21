package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/index"
)

type TopicTag struct {
    ent.Schema
}

func (TopicTag) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable(),
        field.String("name").Unique().MaxLen(20),
        field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
    }
}

func (TopicTag) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("topics", Topic.Type).Ref("tags").Through("topic_taggings", TopicTagging.Type),
    }
}

func (TopicTag) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("name").Unique(),
    }
}
