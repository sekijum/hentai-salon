package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ThreadCommentAttachment struct {
	ent.Schema
}

func (ThreadCommentAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("commentId").StorageKey("comment_id"),
		field.String("url"),
		field.Int("display_order").Default(0),
		field.Int("type").Default(0).Comment("0: image, 1: video"),
		field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
	}
}

func (ThreadCommentAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comment", ThreadComment.Type).Ref("attachments").Unique().Field("commentId").Required(),
	}
}
