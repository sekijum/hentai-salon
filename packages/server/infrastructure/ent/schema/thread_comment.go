package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ThreadComment struct {
	ent.Schema
}

func (ThreadComment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("thread_id"),
		field.Int("parent_comment_id").Optional().Nillable().Comment("親コメントID（リプライの場合）"),
		field.Int("user_id").Optional().Nillable().Comment("ログインユーザーの場合"),
		field.String("guest_name").Optional().Nillable().MaxLen(20).Comment("ゲストユーザーの場合"),
		field.Text("content"),
		field.String("ip_address").MaxLen(64).Comment("コメント者のIPアドレス"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (ThreadComment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("thread", Thread.Type).Ref("comments").Unique().Field("thread_id").Required(),
		edge.From("author", User.Type).Ref("comments").Unique().Field("user_id"),
		edge.From("parent_comment", ThreadComment.Type).Ref("replies").Unique().Field("parent_comment_id"),
		edge.To("replies", ThreadComment.Type),
		edge.To("attachments", ThreadCommentAttachment.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("liked_users", User.Type).Ref("liked_comments").Through("user_comment_like", UserCommentLike.Type),
	}
}
