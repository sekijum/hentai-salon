// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BoardsColumns holds the columns for the "boards" table.
	BoardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "thumbnail_url", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
	}
	// BoardsTable holds the schema information for the "boards" table.
	BoardsTable = &schema.Table{
		Name:       "boards",
		Columns:    BoardsColumns,
		PrimaryKey: []*schema.Column{BoardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "boards_users_boards",
				Columns:    []*schema.Column{BoardsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "board_title",
				Unique:  true,
				Columns: []*schema.Column{BoardsColumns[1]},
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tag_name",
				Unique:  true,
				Columns: []*schema.Column{TagsColumns[1]},
			},
		},
	}
	// ThreadsColumns holds the columns for the "threads" table.
	ThreadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "thumbnail_url", Type: field.TypeString, Nullable: true},
		{Name: "ip_address", Type: field.TypeString, Size: 64},
		{Name: "status", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "board_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ThreadsTable holds the schema information for the "threads" table.
	ThreadsTable = &schema.Table{
		Name:       "threads",
		Columns:    ThreadsColumns,
		PrimaryKey: []*schema.Column{ThreadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "threads_boards_threads",
				Columns:    []*schema.Column{ThreadsColumns[8]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "threads_users_threads",
				Columns:    []*schema.Column{ThreadsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "thread_title",
				Unique:  true,
				Columns: []*schema.Column{ThreadsColumns[1]},
			},
		},
	}
	// ThreadCommentsColumns holds the columns for the "thread_comments" table.
	ThreadCommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "guest_name", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "ip_address", Type: field.TypeString, Size: 64},
		{Name: "status", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "thread_id", Type: field.TypeInt},
		{Name: "parent_comment_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// ThreadCommentsTable holds the schema information for the "thread_comments" table.
	ThreadCommentsTable = &schema.Table{
		Name:       "thread_comments",
		Columns:    ThreadCommentsColumns,
		PrimaryKey: []*schema.Column{ThreadCommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_comments_threads_comments",
				Columns:    []*schema.Column{ThreadCommentsColumns[7]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "thread_comments_thread_comments_replies",
				Columns:    []*schema.Column{ThreadCommentsColumns[8]},
				RefColumns: []*schema.Column{ThreadCommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "thread_comments_users_comments",
				Columns:    []*schema.Column{ThreadCommentsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ThreadCommentAttachmentsColumns holds the columns for the "thread_comment_attachments" table.
	ThreadCommentAttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "display_order", Type: field.TypeInt, Default: 0},
		{Name: "type", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// ThreadCommentAttachmentsTable holds the schema information for the "thread_comment_attachments" table.
	ThreadCommentAttachmentsTable = &schema.Table{
		Name:       "thread_comment_attachments",
		Columns:    ThreadCommentAttachmentsColumns,
		PrimaryKey: []*schema.Column{ThreadCommentAttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_comment_attachments_thread_comments_attachments",
				Columns:    []*schema.Column{ThreadCommentAttachmentsColumns[5]},
				RefColumns: []*schema.Column{ThreadCommentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ThreadTagsColumns holds the columns for the "thread_tags" table.
	ThreadTagsColumns = []*schema.Column{
		{Name: "thread_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// ThreadTagsTable holds the schema information for the "thread_tags" table.
	ThreadTagsTable = &schema.Table{
		Name:       "thread_tags",
		Columns:    ThreadTagsColumns,
		PrimaryKey: []*schema.Column{ThreadTagsColumns[0], ThreadTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_tags_threads_thread",
				Columns:    []*schema.Column{ThreadTagsColumns[0]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "thread_tags_tags_tag",
				Columns:    []*schema.Column{ThreadTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 20},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 254},
		{Name: "password", Type: field.TypeString},
		{Name: "profile_link", Type: field.TypeString, Nullable: true},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 0},
		{Name: "role", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserCommentLikesColumns holds the columns for the "user_comment_likes" table.
	UserCommentLikesColumns = []*schema.Column{
		{Name: "liked_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// UserCommentLikesTable holds the schema information for the "user_comment_likes" table.
	UserCommentLikesTable = &schema.Table{
		Name:       "user_comment_likes",
		Columns:    UserCommentLikesColumns,
		PrimaryKey: []*schema.Column{UserCommentLikesColumns[1], UserCommentLikesColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_comment_likes_users_user",
				Columns:    []*schema.Column{UserCommentLikesColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_comment_likes_thread_comments_comment",
				Columns:    []*schema.Column{UserCommentLikesColumns[2]},
				RefColumns: []*schema.Column{ThreadCommentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserCommentSubscriptionsColumns holds the columns for the "user_comment_subscriptions" table.
	UserCommentSubscriptionsColumns = []*schema.Column{
		{Name: "is_notified", Type: field.TypeBool, Default: true},
		{Name: "is_checked", Type: field.TypeBool, Default: false},
		{Name: "subscribed_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// UserCommentSubscriptionsTable holds the schema information for the "user_comment_subscriptions" table.
	UserCommentSubscriptionsTable = &schema.Table{
		Name:       "user_comment_subscriptions",
		Columns:    UserCommentSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserCommentSubscriptionsColumns[3], UserCommentSubscriptionsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_comment_subscriptions_users_user",
				Columns:    []*schema.Column{UserCommentSubscriptionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_comment_subscriptions_thread_comments_comment",
				Columns:    []*schema.Column{UserCommentSubscriptionsColumns[4]},
				RefColumns: []*schema.Column{ThreadCommentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserThreadLikesColumns holds the columns for the "user_thread_likes" table.
	UserThreadLikesColumns = []*schema.Column{
		{Name: "liked_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "thread_id", Type: field.TypeInt},
	}
	// UserThreadLikesTable holds the schema information for the "user_thread_likes" table.
	UserThreadLikesTable = &schema.Table{
		Name:       "user_thread_likes",
		Columns:    UserThreadLikesColumns,
		PrimaryKey: []*schema.Column{UserThreadLikesColumns[1], UserThreadLikesColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_thread_likes_users_user",
				Columns:    []*schema.Column{UserThreadLikesColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_thread_likes_threads_thread",
				Columns:    []*schema.Column{UserThreadLikesColumns[2]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserThreadSubscriptionsColumns holds the columns for the "user_thread_subscriptions" table.
	UserThreadSubscriptionsColumns = []*schema.Column{
		{Name: "is_notified", Type: field.TypeBool, Default: true},
		{Name: "is_checked", Type: field.TypeBool, Default: false},
		{Name: "subscribed_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "thread_id", Type: field.TypeInt},
	}
	// UserThreadSubscriptionsTable holds the schema information for the "user_thread_subscriptions" table.
	UserThreadSubscriptionsTable = &schema.Table{
		Name:       "user_thread_subscriptions",
		Columns:    UserThreadSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserThreadSubscriptionsColumns[3], UserThreadSubscriptionsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_thread_subscriptions_users_user",
				Columns:    []*schema.Column{UserThreadSubscriptionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_thread_subscriptions_threads_thread",
				Columns:    []*schema.Column{UserThreadSubscriptionsColumns[4]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BoardsTable,
		TagsTable,
		ThreadsTable,
		ThreadCommentsTable,
		ThreadCommentAttachmentsTable,
		ThreadTagsTable,
		UsersTable,
		UserCommentLikesTable,
		UserCommentSubscriptionsTable,
		UserThreadLikesTable,
		UserThreadSubscriptionsTable,
	}
)

func init() {
	BoardsTable.ForeignKeys[0].RefTable = UsersTable
	ThreadsTable.ForeignKeys[0].RefTable = BoardsTable
	ThreadsTable.ForeignKeys[1].RefTable = UsersTable
	ThreadCommentsTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadCommentsTable.ForeignKeys[1].RefTable = ThreadCommentsTable
	ThreadCommentsTable.ForeignKeys[2].RefTable = UsersTable
	ThreadCommentAttachmentsTable.ForeignKeys[0].RefTable = ThreadCommentsTable
	ThreadTagsTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadTagsTable.ForeignKeys[1].RefTable = TagsTable
	UserCommentLikesTable.ForeignKeys[0].RefTable = UsersTable
	UserCommentLikesTable.ForeignKeys[1].RefTable = ThreadCommentsTable
	UserCommentSubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	UserCommentSubscriptionsTable.ForeignKeys[1].RefTable = ThreadCommentsTable
	UserThreadLikesTable.ForeignKeys[0].RefTable = UsersTable
	UserThreadLikesTable.ForeignKeys[1].RefTable = ThreadsTable
	UserThreadSubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	UserThreadSubscriptionsTable.ForeignKeys[1].RefTable = ThreadsTable
}
