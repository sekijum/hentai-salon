package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Thread struct {
	ent.Schema
}

func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("board_id"),
		field.Int("user_id"),
		field.String("title").Unique().MaxLen(255),
		field.Text("description").Optional().Nillable(),
		field.String("thumbnail_url").Optional().Nillable(),
		field.String("ip_address").MaxLen(64).Comment("スレッド作成者のIPアドレス"),
		field.Int("status").Default(0).Comment("0: Open, 1: Pending, 2: Archived"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("board", Board.Type).Ref("threads").Unique().Field("board_id").Required(),
		edge.From("owner", User.Type).Ref("threads").Unique().Field("user_id").Required(),
		edge.To("comments", ThreadComment.Type),
		edge.To("tags", Tag.Type).Through("thread_tags", ThreadTag.Type),
		edge.From("liked_users", User.Type).Ref("liked_threads").Through("user_thread_like", UserThreadLike.Type),
	}
}

func (Thread) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique(),
	}
}
