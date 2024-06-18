package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
    ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.Int("forum_id"),
        field.Int("user_id").Optional(),
        field.String("title"),
        field.Text("content"),
        field.Bool("is_default").Default(false),
        field.Enum("status").Values("Open", "Closed", "Archived", "Disapproved").Default("Open"),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("forum", Forum.Type).Ref("topics").Unique().Field("forum_id").Required(),
        edge.From("user", User.Type).Ref("topics").Unique().Field("user_id"),
        edge.To("comments", Comment.Type),
        edge.To("topic_likes", TopicLike.Type),
        edge.To("user_topic_notifications", UserTopicNotification.Type),
    }
}
