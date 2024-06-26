package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CommentAttachment struct {
	ent.Schema
}

func (CommentAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("commentId").StorageKey("comment_id"),
		field.String("url"),
		field.Int("order").Default(0),
		field.Enum("type").Values("image", "video"),
	}
}

func (CommentAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comment", Comment.Type).Ref("comment_attachments").Unique().Field("commentId").Required(),
	}
}
