// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usercommentlike"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserCommentLike is the model entity for the UserCommentLike schema.
type UserCommentLike struct {
	config `json:"-"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// CommentID holds the value of the "comment_id" field.
	CommentID uint64 `json:"comment_id,omitempty"`
	// LikedAt holds the value of the "liked_at" field.
	LikedAt time.Time `json:"liked_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserCommentLikeQuery when eager-loading is set.
	Edges        UserCommentLikeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserCommentLikeEdges holds the relations/edges for other nodes in the graph.
type UserCommentLikeEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Comment holds the value of the comment edge.
	Comment *ThreadComment `json:"comment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserCommentLikeEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserCommentLikeEdges) CommentOrErr() (*ThreadComment, error) {
	if e.Comment != nil {
		return e.Comment, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: threadcomment.Label}
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserCommentLike) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usercommentlike.FieldUserID, usercommentlike.FieldCommentID:
			values[i] = new(sql.NullInt64)
		case usercommentlike.FieldLikedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserCommentLike fields.
func (ucl *UserCommentLike) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usercommentlike.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ucl.UserID = int(value.Int64)
			}
		case usercommentlike.FieldCommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment_id", values[i])
			} else if value.Valid {
				ucl.CommentID = uint64(value.Int64)
			}
		case usercommentlike.FieldLikedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field liked_at", values[i])
			} else if value.Valid {
				ucl.LikedAt = value.Time
			}
		default:
			ucl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserCommentLike.
// This includes values selected through modifiers, order, etc.
func (ucl *UserCommentLike) Value(name string) (ent.Value, error) {
	return ucl.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserCommentLike entity.
func (ucl *UserCommentLike) QueryUser() *UserQuery {
	return NewUserCommentLikeClient(ucl.config).QueryUser(ucl)
}

// QueryComment queries the "comment" edge of the UserCommentLike entity.
func (ucl *UserCommentLike) QueryComment() *ThreadCommentQuery {
	return NewUserCommentLikeClient(ucl.config).QueryComment(ucl)
}

// Update returns a builder for updating this UserCommentLike.
// Note that you need to call UserCommentLike.Unwrap() before calling this method if this UserCommentLike
// was returned from a transaction, and the transaction was committed or rolled back.
func (ucl *UserCommentLike) Update() *UserCommentLikeUpdateOne {
	return NewUserCommentLikeClient(ucl.config).UpdateOne(ucl)
}

// Unwrap unwraps the UserCommentLike entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ucl *UserCommentLike) Unwrap() *UserCommentLike {
	_tx, ok := ucl.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserCommentLike is not a transactional entity")
	}
	ucl.config.driver = _tx.drv
	return ucl
}

// String implements the fmt.Stringer.
func (ucl *UserCommentLike) String() string {
	var builder strings.Builder
	builder.WriteString("UserCommentLike(")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ucl.UserID))
	builder.WriteString(", ")
	builder.WriteString("comment_id=")
	builder.WriteString(fmt.Sprintf("%v", ucl.CommentID))
	builder.WriteString(", ")
	builder.WriteString("liked_at=")
	builder.WriteString(ucl.LikedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserCommentLikes is a parsable slice of UserCommentLike.
type UserCommentLikes []*UserCommentLike
