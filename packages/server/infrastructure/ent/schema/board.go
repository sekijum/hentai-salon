package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.Int("status").Default(0).Comment("0: Public, 1: Private, 3: Pending, 3: Archived"),
		field.Time("createdAt").Default(time.Now).StorageKey("created_at"),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now).StorageKey("updated_at"),
	}
}

func (Board) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("liked_users", User.Type).Ref("subscribed_boards").Through("user_board_like", UserBoardSubscription.Type),
		edge.From("subscribed_users", User.Type).Ref("liked_boards").Through("user_board_subscription", UserBoardLike.Type),
		edge.From("owner", User.Type).Ref("boards").Unique().Field("userId").Required(),
		edge.To("threads", Thread.Type),
	}
}

func (Board) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique(),
	}
}
