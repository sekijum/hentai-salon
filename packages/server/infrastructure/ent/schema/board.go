package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Board struct {
	ent.Schema
}

func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("userId").StorageKey("user_id"),
		field.String("title").Unique().MaxLen(50).Comment("板名"),
		field.String("description").Optional().MaxLen(255),
		field.String("thumbnailUrl").Optional().StorageKey("thumbnail_url"),
		field.Int("order").Default(0),
		field.Enum("status").Values("Public", "Private", "Archived", "Deleted").Default("Public"),
		field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).StorageKey("updated_at"),
	}
}

func (Board) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("liked_users", User.Type).Ref("subscribed_boards").Through("user_board_like", UserBoardSubscription.Type),
		edge.From("subscribed_users", User.Type).Ref("liked_boards").Through("user_board_subscription", UserBoardLike.Type),
		edge.To("threads", Thread.Type),
	}
}
