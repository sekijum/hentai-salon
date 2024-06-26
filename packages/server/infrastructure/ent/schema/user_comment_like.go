package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserCommentLike struct {
    ent.Schema
}

func (UserCommentLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("userId", "commentId"),
	}
}

func (UserCommentLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("commentId").StorageKey("comment_id"),
        field.Time("likedAt").Default(time.Now).StorageKey("liked_at"),
    }
}

func (UserCommentLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("comment", ThreadComment.Type).Unique().Required().Field("commentId"),
    }
}
