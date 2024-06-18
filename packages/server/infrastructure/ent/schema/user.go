package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StorageKey("id"),
        field.String("username").Unique(),
        field.String("email").Unique(),
        field.String("password"),
        field.String("display_name").Optional(),
        field.String("avatar").Optional(),
        field.Enum("status").Values("Active", "Withdrawn", "Suspended", "Inactive").Default("Active"),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("forums", Forum.Type),
        edge.To("topics", Topic.Type),
        edge.To("comments", Comment.Type),
        edge.To("user_topic_notifications", UserTopicNotification.Type),
        edge.To("user_comment_notifications", UserCommentNotification.Type),
        edge.To("forum_likes", ForumLike.Type),
        edge.To("topic_likes", TopicLike.Type),
        edge.To("comment_likes", CommentLike.Type),
    }
}
