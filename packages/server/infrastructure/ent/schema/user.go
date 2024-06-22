package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

type User struct {
    ent.Schema
}

func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable(),
        field.String("name").Unique().StorageKey("name").MaxLen(20),
        field.String("email").Unique().MaxLen(254),
        field.String("password"),
        field.String("displayName").Optional().StorageKey("display_name").MaxLen(20),
        field.String("avatarUrl").Optional().StorageKey("avatar_url"),
        field.Enum("status").Values("Active", "Withdrawn", "Suspended", "Inactive").Default("Active"),
        field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
        field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).StorageKey("updated_at"),
    }
}

func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("boards", Board.Type),
        edge.To("threads", Thread.Type),
        edge.To("comments", Comment.Type),
        edge.To("liked_boards", Board.Type).Through("user_board_like", UserBoardLike.Type),
        edge.To("liked_threads", Thread.Type).Through("user_thread_like", UserThreadLike.Type),
        edge.To("liked_comments", Comment.Type).Through("user_comment_like", UserCommentLike.Type),
        edge.To("subscribed_boards", Board.Type).Through("user_board_subscription", UserBoardSubscription.Type),
        edge.To("subscribed_threads", Thread.Type).Through("user_thread_subscription", UserThreadSubscription.Type),
        edge.To("subscribed_comments", Comment.Type).Through("user_comment_subscription", UserCommentSubscription.Type),
    }
}
