// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/comment"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ThreadId holds the value of the "threadId" field.
	ThreadId int `json:"threadId,omitempty"`
	// 親コメントID（リプライの場合）
	ParentCommentId int `json:"parentCommentId,omitempty"`
	// ログインユーザーの場合
	UserId int `json:"userId,omitempty"`
	// ゲストユーザーの場合
	GuestName string `json:"guestName,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// コメント者のIPアドレス
	IPAddress string `json:"ip_address,omitempty"`
	// Status holds the value of the "status" field.
	Status comment.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges        CommentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Thread holds the value of the thread edge.
	Thread *Thread `json:"thread,omitempty"`
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// ParentComment holds the value of the parent_comment edge.
	ParentComment *Comment `json:"parent_comment,omitempty"`
	// Replies holds the value of the replies edge.
	Replies []*Comment `json:"replies,omitempty"`
	// CommentAttachments holds the value of the comment_attachments edge.
	CommentAttachments []*CommentAttachment `json:"comment_attachments,omitempty"`
	// LikedUsers holds the value of the liked_users edge.
	LikedUsers []*User `json:"liked_users,omitempty"`
	// SubscribedUsers holds the value of the subscribed_users edge.
	SubscribedUsers []*User `json:"subscribed_users,omitempty"`
	// UserCommentLike holds the value of the user_comment_like edge.
	UserCommentLike []*UserCommentLike `json:"user_comment_like,omitempty"`
	// UserCommentSubscription holds the value of the user_comment_subscription edge.
	UserCommentSubscription []*UserCommentSubscription `json:"user_comment_subscription,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// ThreadOrErr returns the Thread value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) ThreadOrErr() (*Thread, error) {
	if e.Thread != nil {
		return e.Thread, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: thread.Label}
	}
	return nil, &NotLoadedError{edge: "thread"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) AuthorOrErr() (*User, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// ParentCommentOrErr returns the ParentComment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) ParentCommentOrErr() (*Comment, error) {
	if e.ParentComment != nil {
		return e.ParentComment, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: comment.Label}
	}
	return nil, &NotLoadedError{edge: "parent_comment"}
}

// RepliesOrErr returns the Replies value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) RepliesOrErr() ([]*Comment, error) {
	if e.loadedTypes[3] {
		return e.Replies, nil
	}
	return nil, &NotLoadedError{edge: "replies"}
}

// CommentAttachmentsOrErr returns the CommentAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) CommentAttachmentsOrErr() ([]*CommentAttachment, error) {
	if e.loadedTypes[4] {
		return e.CommentAttachments, nil
	}
	return nil, &NotLoadedError{edge: "comment_attachments"}
}

// LikedUsersOrErr returns the LikedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) LikedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.LikedUsers, nil
	}
	return nil, &NotLoadedError{edge: "liked_users"}
}

// SubscribedUsersOrErr returns the SubscribedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) SubscribedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[6] {
		return e.SubscribedUsers, nil
	}
	return nil, &NotLoadedError{edge: "subscribed_users"}
}

// UserCommentLikeOrErr returns the UserCommentLike value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) UserCommentLikeOrErr() ([]*UserCommentLike, error) {
	if e.loadedTypes[7] {
		return e.UserCommentLike, nil
	}
	return nil, &NotLoadedError{edge: "user_comment_like"}
}

// UserCommentSubscriptionOrErr returns the UserCommentSubscription value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) UserCommentSubscriptionOrErr() ([]*UserCommentSubscription, error) {
	if e.loadedTypes[8] {
		return e.UserCommentSubscription, nil
	}
	return nil, &NotLoadedError{edge: "user_comment_subscription"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldThreadId, comment.FieldParentCommentId, comment.FieldUserId:
			values[i] = new(sql.NullInt64)
		case comment.FieldGuestName, comment.FieldMessage, comment.FieldIPAddress, comment.FieldStatus:
			values[i] = new(sql.NullString)
		case comment.FieldCreatedAt, comment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comment.FieldThreadId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field threadId", values[i])
			} else if value.Valid {
				c.ThreadId = int(value.Int64)
			}
		case comment.FieldParentCommentId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parentCommentId", values[i])
			} else if value.Valid {
				c.ParentCommentId = int(value.Int64)
			}
		case comment.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				c.UserId = int(value.Int64)
			}
		case comment.FieldGuestName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field guestName", values[i])
			} else if value.Valid {
				c.GuestName = value.String
			}
		case comment.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				c.Message = value.String
			}
		case comment.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				c.IPAddress = value.String
			}
		case comment.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = comment.Status(value.String)
			}
		case comment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Comment.
// This includes values selected through modifiers, order, etc.
func (c *Comment) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryThread queries the "thread" edge of the Comment entity.
func (c *Comment) QueryThread() *ThreadQuery {
	return NewCommentClient(c.config).QueryThread(c)
}

// QueryAuthor queries the "author" edge of the Comment entity.
func (c *Comment) QueryAuthor() *UserQuery {
	return NewCommentClient(c.config).QueryAuthor(c)
}

// QueryParentComment queries the "parent_comment" edge of the Comment entity.
func (c *Comment) QueryParentComment() *CommentQuery {
	return NewCommentClient(c.config).QueryParentComment(c)
}

// QueryReplies queries the "replies" edge of the Comment entity.
func (c *Comment) QueryReplies() *CommentQuery {
	return NewCommentClient(c.config).QueryReplies(c)
}

// QueryCommentAttachments queries the "comment_attachments" edge of the Comment entity.
func (c *Comment) QueryCommentAttachments() *CommentAttachmentQuery {
	return NewCommentClient(c.config).QueryCommentAttachments(c)
}

// QueryLikedUsers queries the "liked_users" edge of the Comment entity.
func (c *Comment) QueryLikedUsers() *UserQuery {
	return NewCommentClient(c.config).QueryLikedUsers(c)
}

// QuerySubscribedUsers queries the "subscribed_users" edge of the Comment entity.
func (c *Comment) QuerySubscribedUsers() *UserQuery {
	return NewCommentClient(c.config).QuerySubscribedUsers(c)
}

// QueryUserCommentLike queries the "user_comment_like" edge of the Comment entity.
func (c *Comment) QueryUserCommentLike() *UserCommentLikeQuery {
	return NewCommentClient(c.config).QueryUserCommentLike(c)
}

// QueryUserCommentSubscription queries the "user_comment_subscription" edge of the Comment entity.
func (c *Comment) QueryUserCommentSubscription() *UserCommentSubscriptionQuery {
	return NewCommentClient(c.config).QueryUserCommentSubscription(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return NewCommentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("threadId=")
	builder.WriteString(fmt.Sprintf("%v", c.ThreadId))
	builder.WriteString(", ")
	builder.WriteString("parentCommentId=")
	builder.WriteString(fmt.Sprintf("%v", c.ParentCommentId))
	builder.WriteString(", ")
	builder.WriteString("userId=")
	builder.WriteString(fmt.Sprintf("%v", c.UserId))
	builder.WriteString(", ")
	builder.WriteString("guestName=")
	builder.WriteString(c.GuestName)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(c.Message)
	builder.WriteString(", ")
	builder.WriteString("ip_address=")
	builder.WriteString(c.IPAddress)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment