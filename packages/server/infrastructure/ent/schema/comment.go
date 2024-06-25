package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("threadId").StorageKey("thread_id"),
		field.Int("parentCommentId").Optional().StorageKey("parent_comment_id").Comment("親コメントID（リプライの場合）"),
		field.Int("userId").Optional().StorageKey("user_id").Comment("ログインユーザーの場合"),
		field.String("guestName").Optional().StorageKey("guest_name").MaxLen(20).Comment("ゲストユーザーの場合"),
		field.Text("message"),
		field.String("ip_address").MaxLen(64).Comment("コメント者のIPアドレス"),
		field.Int("status").Default(0), // 0: Visible, 1: Deleted
		field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).StorageKey("updated_at"),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("thread", Thread.Type).Ref("comments").Unique().Field("threadId").Required(),
		edge.From("author", User.Type).Ref("comments").Unique().Field("userId"),
		edge.From("parent_comment", Comment.Type).Ref("replies").Unique().Field("parentCommentId"),
		edge.To("replies", Comment.Type),
		edge.To("comment_attachments", CommentAttachment.Type),
		edge.From("liked_users", User.Type).Ref("liked_comments").Through("user_comment_like", UserCommentLike.Type),
		edge.From("subscribed_users", User.Type).Ref("subscribed_comments").Through("user_comment_subscription", UserCommentSubscription.Type),
	}
}
