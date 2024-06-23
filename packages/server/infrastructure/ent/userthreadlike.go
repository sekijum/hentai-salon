// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userthreadlike"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserThreadLike is the model entity for the UserThreadLike schema.
type UserThreadLike struct {
	config `json:"-"`
	// UserId holds the value of the "userId" field.
	UserId int `json:"userId,omitempty"`
	// ThreadId holds the value of the "threadId" field.
	ThreadId int `json:"threadId,omitempty"`
	// LikedAt holds the value of the "likedAt" field.
	LikedAt time.Time `json:"likedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserThreadLikeQuery when eager-loading is set.
	Edges        UserThreadLikeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserThreadLikeEdges holds the relations/edges for other nodes in the graph.
type UserThreadLikeEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Thread holds the value of the thread edge.
	Thread *Thread `json:"thread,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserThreadLikeEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ThreadOrErr returns the Thread value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserThreadLikeEdges) ThreadOrErr() (*Thread, error) {
	if e.Thread != nil {
		return e.Thread, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: thread.Label}
	}
	return nil, &NotLoadedError{edge: "thread"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserThreadLike) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userthreadlike.FieldUserId, userthreadlike.FieldThreadId:
			values[i] = new(sql.NullInt64)
		case userthreadlike.FieldLikedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserThreadLike fields.
func (utl *UserThreadLike) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userthreadlike.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				utl.UserId = int(value.Int64)
			}
		case userthreadlike.FieldThreadId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field threadId", values[i])
			} else if value.Valid {
				utl.ThreadId = int(value.Int64)
			}
		case userthreadlike.FieldLikedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field likedAt", values[i])
			} else if value.Valid {
				utl.LikedAt = value.Time
			}
		default:
			utl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserThreadLike.
// This includes values selected through modifiers, order, etc.
func (utl *UserThreadLike) Value(name string) (ent.Value, error) {
	return utl.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserThreadLike entity.
func (utl *UserThreadLike) QueryUser() *UserQuery {
	return NewUserThreadLikeClient(utl.config).QueryUser(utl)
}

// QueryThread queries the "thread" edge of the UserThreadLike entity.
func (utl *UserThreadLike) QueryThread() *ThreadQuery {
	return NewUserThreadLikeClient(utl.config).QueryThread(utl)
}

// Update returns a builder for updating this UserThreadLike.
// Note that you need to call UserThreadLike.Unwrap() before calling this method if this UserThreadLike
// was returned from a transaction, and the transaction was committed or rolled back.
func (utl *UserThreadLike) Update() *UserThreadLikeUpdateOne {
	return NewUserThreadLikeClient(utl.config).UpdateOne(utl)
}

// Unwrap unwraps the UserThreadLike entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (utl *UserThreadLike) Unwrap() *UserThreadLike {
	_tx, ok := utl.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserThreadLike is not a transactional entity")
	}
	utl.config.driver = _tx.drv
	return utl
}

// String implements the fmt.Stringer.
func (utl *UserThreadLike) String() string {
	var builder strings.Builder
	builder.WriteString("UserThreadLike(")
	builder.WriteString("userId=")
	builder.WriteString(fmt.Sprintf("%v", utl.UserId))
	builder.WriteString(", ")
	builder.WriteString("threadId=")
	builder.WriteString(fmt.Sprintf("%v", utl.ThreadId))
	builder.WriteString(", ")
	builder.WriteString("likedAt=")
	builder.WriteString(utl.LikedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserThreadLikes is a parsable slice of UserThreadLike.
type UserThreadLikes []*UserThreadLike