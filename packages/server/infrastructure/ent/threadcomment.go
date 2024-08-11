// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ThreadComment is the model entity for the ThreadComment schema.
type ThreadComment struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// ThreadID holds the value of the "thread_id" field.
	ThreadID int `json:"thread_id,omitempty"`
	// 親コメントID（リプライの場合）
	ParentCommentID *uint64 `json:"parent_comment_id,omitempty"`
	// ログインユーザーの場合
	UserID *int `json:"user_id,omitempty"`
	// ゲストユーザーの場合
	GuestName *string `json:"guest_name,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// コメント者のIPアドレス
	IPAddress string `json:"ip_address,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ThreadCommentQuery when eager-loading is set.
	Edges        ThreadCommentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ThreadCommentEdges holds the relations/edges for other nodes in the graph.
type ThreadCommentEdges struct {
	// Thread holds the value of the thread edge.
	Thread *Thread `json:"thread,omitempty"`
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// ParentComment holds the value of the parent_comment edge.
	ParentComment *ThreadComment `json:"parent_comment,omitempty"`
	// Replies holds the value of the replies edge.
	Replies []*ThreadComment `json:"replies,omitempty"`
	// Attachments holds the value of the attachments edge.
	Attachments []*ThreadCommentAttachment `json:"attachments,omitempty"`
	// LikedUsers holds the value of the liked_users edge.
	LikedUsers []*User `json:"liked_users,omitempty"`
	// UserCommentLike holds the value of the user_comment_like edge.
	UserCommentLike []*UserCommentLike `json:"user_comment_like,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// ThreadOrErr returns the Thread value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadCommentEdges) ThreadOrErr() (*Thread, error) {
	if e.Thread != nil {
		return e.Thread, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: thread.Label}
	}
	return nil, &NotLoadedError{edge: "thread"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadCommentEdges) AuthorOrErr() (*User, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// ParentCommentOrErr returns the ParentComment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadCommentEdges) ParentCommentOrErr() (*ThreadComment, error) {
	if e.ParentComment != nil {
		return e.ParentComment, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: threadcomment.Label}
	}
	return nil, &NotLoadedError{edge: "parent_comment"}
}

// RepliesOrErr returns the Replies value or an error if the edge
// was not loaded in eager-loading.
func (e ThreadCommentEdges) RepliesOrErr() ([]*ThreadComment, error) {
	if e.loadedTypes[3] {
		return e.Replies, nil
	}
	return nil, &NotLoadedError{edge: "replies"}
}

// AttachmentsOrErr returns the Attachments value or an error if the edge
// was not loaded in eager-loading.
func (e ThreadCommentEdges) AttachmentsOrErr() ([]*ThreadCommentAttachment, error) {
	if e.loadedTypes[4] {
		return e.Attachments, nil
	}
	return nil, &NotLoadedError{edge: "attachments"}
}

// LikedUsersOrErr returns the LikedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e ThreadCommentEdges) LikedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.LikedUsers, nil
	}
	return nil, &NotLoadedError{edge: "liked_users"}
}

// UserCommentLikeOrErr returns the UserCommentLike value or an error if the edge
// was not loaded in eager-loading.
func (e ThreadCommentEdges) UserCommentLikeOrErr() ([]*UserCommentLike, error) {
	if e.loadedTypes[6] {
		return e.UserCommentLike, nil
	}
	return nil, &NotLoadedError{edge: "user_comment_like"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ThreadComment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case threadcomment.FieldID, threadcomment.FieldThreadID, threadcomment.FieldParentCommentID, threadcomment.FieldUserID:
			values[i] = new(sql.NullInt64)
		case threadcomment.FieldGuestName, threadcomment.FieldContent, threadcomment.FieldIPAddress:
			values[i] = new(sql.NullString)
		case threadcomment.FieldCreatedAt, threadcomment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ThreadComment fields.
func (tc *ThreadComment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case threadcomment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = uint64(value.Int64)
		case threadcomment.FieldThreadID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field thread_id", values[i])
			} else if value.Valid {
				tc.ThreadID = int(value.Int64)
			}
		case threadcomment.FieldParentCommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_comment_id", values[i])
			} else if value.Valid {
				tc.ParentCommentID = new(uint64)
				*tc.ParentCommentID = uint64(value.Int64)
			}
		case threadcomment.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tc.UserID = new(int)
				*tc.UserID = int(value.Int64)
			}
		case threadcomment.FieldGuestName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field guest_name", values[i])
			} else if value.Valid {
				tc.GuestName = new(string)
				*tc.GuestName = value.String
			}
		case threadcomment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				tc.Content = value.String
			}
		case threadcomment.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				tc.IPAddress = value.String
			}
		case threadcomment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tc.CreatedAt = value.Time
			}
		case threadcomment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tc.UpdatedAt = value.Time
			}
		default:
			tc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ThreadComment.
// This includes values selected through modifiers, order, etc.
func (tc *ThreadComment) Value(name string) (ent.Value, error) {
	return tc.selectValues.Get(name)
}

// QueryThread queries the "thread" edge of the ThreadComment entity.
func (tc *ThreadComment) QueryThread() *ThreadQuery {
	return NewThreadCommentClient(tc.config).QueryThread(tc)
}

// QueryAuthor queries the "author" edge of the ThreadComment entity.
func (tc *ThreadComment) QueryAuthor() *UserQuery {
	return NewThreadCommentClient(tc.config).QueryAuthor(tc)
}

// QueryParentComment queries the "parent_comment" edge of the ThreadComment entity.
func (tc *ThreadComment) QueryParentComment() *ThreadCommentQuery {
	return NewThreadCommentClient(tc.config).QueryParentComment(tc)
}

// QueryReplies queries the "replies" edge of the ThreadComment entity.
func (tc *ThreadComment) QueryReplies() *ThreadCommentQuery {
	return NewThreadCommentClient(tc.config).QueryReplies(tc)
}

// QueryAttachments queries the "attachments" edge of the ThreadComment entity.
func (tc *ThreadComment) QueryAttachments() *ThreadCommentAttachmentQuery {
	return NewThreadCommentClient(tc.config).QueryAttachments(tc)
}

// QueryLikedUsers queries the "liked_users" edge of the ThreadComment entity.
func (tc *ThreadComment) QueryLikedUsers() *UserQuery {
	return NewThreadCommentClient(tc.config).QueryLikedUsers(tc)
}

// QueryUserCommentLike queries the "user_comment_like" edge of the ThreadComment entity.
func (tc *ThreadComment) QueryUserCommentLike() *UserCommentLikeQuery {
	return NewThreadCommentClient(tc.config).QueryUserCommentLike(tc)
}

// Update returns a builder for updating this ThreadComment.
// Note that you need to call ThreadComment.Unwrap() before calling this method if this ThreadComment
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *ThreadComment) Update() *ThreadCommentUpdateOne {
	return NewThreadCommentClient(tc.config).UpdateOne(tc)
}

// Unwrap unwraps the ThreadComment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *ThreadComment) Unwrap() *ThreadComment {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ThreadComment is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *ThreadComment) String() string {
	var builder strings.Builder
	builder.WriteString("ThreadComment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("thread_id=")
	builder.WriteString(fmt.Sprintf("%v", tc.ThreadID))
	builder.WriteString(", ")
	if v := tc.ParentCommentID; v != nil {
		builder.WriteString("parent_comment_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := tc.UserID; v != nil {
		builder.WriteString("user_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := tc.GuestName; v != nil {
		builder.WriteString("guest_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(tc.Content)
	builder.WriteString(", ")
	builder.WriteString("ip_address=")
	builder.WriteString(tc.IPAddress)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(tc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(tc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ThreadComments is a parsable slice of ThreadComment.
type ThreadComments []*ThreadComment
