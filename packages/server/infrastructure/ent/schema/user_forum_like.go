package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema"
)

type UserForumLike struct {
    ent.Schema
}

func (UserForumLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("userId", "forumId"),
	}
}

func (UserForumLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("forumId").StorageKey("forum_id"),
        field.Enum("type").Values("like", "dislike"),
        field.Time("likedAt").Default(time.Now).StorageKey("liked_at"),
    }
}

func (UserForumLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("forum", Forum.Type).Unique().Required().Field("forumId"),
    }
}
