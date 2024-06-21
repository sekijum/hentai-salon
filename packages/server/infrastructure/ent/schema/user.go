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
        field.String("userName").Unique().StorageKey("user_name").MaxLen(20),
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
        edge.To("forums", Forum.Type),
        edge.To("topics", Topic.Type),
        edge.To("comments", Comment.Type),
        edge.To("liked_forums", Forum.Type).Through("user_forum_like", UserForumLike.Type),
        edge.To("liked_topics", Topic.Type).Through("user_topic_like", UserTopicLike.Type),
        edge.To("liked_comments", Comment.Type).Through("user_comment_like", UserCommentLike.Type),
        edge.To("subscribed_forums", Forum.Type).Through("user_forum_subscription", UserForumSubscription.Type),
        edge.To("subscribed_topics", Topic.Type).Through("user_topic_subscription", UserTopicSubscription.Type),
        edge.To("subscribed_comments", Comment.Type).Through("user_comment_subscription", UserCommentSubscription.Type),
    }
}
