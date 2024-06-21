package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

type Forum struct {
    ent.Schema
}

func (Forum) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable(),
        field.Int("userId").StorageKey("user_id"),
        field.String("title").MaxLen(50),
        field.String("description").Optional().MaxLen(255),
        field.String("thumbnailUrl").Optional().StorageKey("thumbnail_url"),
        field.Enum("status").Values("Public", "Private", "Archived").Default("Public"),
        field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
        field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).StorageKey("updated_at"),
    }
}

func (Forum) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("liked_users", User.Type).Ref("subscribed_forums").Through("user_forum_like", UserForumSubscription.Type),
        edge.From("subscribed_users", User.Type).Ref("liked_forums").Through("user_forum_subscription", UserForumLike.Type),
        edge.To("topics", Topic.Type),
    }
}
