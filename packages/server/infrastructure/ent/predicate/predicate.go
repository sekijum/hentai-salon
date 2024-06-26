// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Board is the predicate function for board builders.
type Board func(*sql.Selector)

// Tag is the predicate function for tag builders.
type Tag func(*sql.Selector)

// Thread is the predicate function for thread builders.
type Thread func(*sql.Selector)

// ThreadComment is the predicate function for threadcomment builders.
type ThreadComment func(*sql.Selector)

// ThreadCommentAttachment is the predicate function for threadcommentattachment builders.
type ThreadCommentAttachment func(*sql.Selector)

// ThreadTag is the predicate function for threadtag builders.
type ThreadTag func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)

// UserBoardLike is the predicate function for userboardlike builders.
type UserBoardLike func(*sql.Selector)

// UserBoardSubscription is the predicate function for userboardsubscription builders.
type UserBoardSubscription func(*sql.Selector)

// UserCommentLike is the predicate function for usercommentlike builders.
type UserCommentLike func(*sql.Selector)

// UserCommentSubscription is the predicate function for usercommentsubscription builders.
type UserCommentSubscription func(*sql.Selector)

// UserThreadLike is the predicate function for userthreadlike builders.
type UserThreadLike func(*sql.Selector)

// UserThreadSubscription is the predicate function for userthreadsubscription builders.
type UserThreadSubscription func(*sql.Selector)
