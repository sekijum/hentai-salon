package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema"
)

type UserBoardLike struct {
    ent.Schema
}

func (UserBoardLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("userId", "boardId"),
	}
}

func (UserBoardLike) Fields() []ent.Field {
    return []ent.Field{
        field.Int("userId").StorageKey("user_id"),
        field.Int("boardId").StorageKey("board_id"),
        field.Time("likedAt").Default(time.Now).StorageKey("liked_at"),
    }
}

func (UserBoardLike) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("user", User.Type).Unique().Required().Field("userId"),
        edge.To("board", Board.Type).Unique().Required().Field("boardId"),
    }
}
