// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userthreadsubscription"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserThreadSubscription is the model entity for the UserThreadSubscription schema.
type UserThreadSubscription struct {
	config `json:"-"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// ThreadID holds the value of the "thread_id" field.
	ThreadID int `json:"thread_id,omitempty"`
	// メール通知を受け取るかどうかのフラグ
	IsNotified bool `json:"is_notified,omitempty"`
	// 通知画面で確認したかどうかのフラグ
	IsChecked bool `json:"is_checked,omitempty"`
	// SubscribedAt holds the value of the "subscribed_at" field.
	SubscribedAt time.Time `json:"subscribed_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserThreadSubscriptionQuery when eager-loading is set.
	Edges        UserThreadSubscriptionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserThreadSubscriptionEdges holds the relations/edges for other nodes in the graph.
type UserThreadSubscriptionEdges struct {
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
func (e UserThreadSubscriptionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ThreadOrErr returns the Thread value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserThreadSubscriptionEdges) ThreadOrErr() (*Thread, error) {
	if e.Thread != nil {
		return e.Thread, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: thread.Label}
	}
	return nil, &NotLoadedError{edge: "thread"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserThreadSubscription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userthreadsubscription.FieldIsNotified, userthreadsubscription.FieldIsChecked:
			values[i] = new(sql.NullBool)
		case userthreadsubscription.FieldUserID, userthreadsubscription.FieldThreadID:
			values[i] = new(sql.NullInt64)
		case userthreadsubscription.FieldSubscribedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserThreadSubscription fields.
func (uts *UserThreadSubscription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userthreadsubscription.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uts.UserID = int(value.Int64)
			}
		case userthreadsubscription.FieldThreadID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field thread_id", values[i])
			} else if value.Valid {
				uts.ThreadID = int(value.Int64)
			}
		case userthreadsubscription.FieldIsNotified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_notified", values[i])
			} else if value.Valid {
				uts.IsNotified = value.Bool
			}
		case userthreadsubscription.FieldIsChecked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_checked", values[i])
			} else if value.Valid {
				uts.IsChecked = value.Bool
			}
		case userthreadsubscription.FieldSubscribedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field subscribed_at", values[i])
			} else if value.Valid {
				uts.SubscribedAt = value.Time
			}
		default:
			uts.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserThreadSubscription.
// This includes values selected through modifiers, order, etc.
func (uts *UserThreadSubscription) Value(name string) (ent.Value, error) {
	return uts.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserThreadSubscription entity.
func (uts *UserThreadSubscription) QueryUser() *UserQuery {
	return NewUserThreadSubscriptionClient(uts.config).QueryUser(uts)
}

// QueryThread queries the "thread" edge of the UserThreadSubscription entity.
func (uts *UserThreadSubscription) QueryThread() *ThreadQuery {
	return NewUserThreadSubscriptionClient(uts.config).QueryThread(uts)
}

// Update returns a builder for updating this UserThreadSubscription.
// Note that you need to call UserThreadSubscription.Unwrap() before calling this method if this UserThreadSubscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (uts *UserThreadSubscription) Update() *UserThreadSubscriptionUpdateOne {
	return NewUserThreadSubscriptionClient(uts.config).UpdateOne(uts)
}

// Unwrap unwraps the UserThreadSubscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uts *UserThreadSubscription) Unwrap() *UserThreadSubscription {
	_tx, ok := uts.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserThreadSubscription is not a transactional entity")
	}
	uts.config.driver = _tx.drv
	return uts
}

// String implements the fmt.Stringer.
func (uts *UserThreadSubscription) String() string {
	var builder strings.Builder
	builder.WriteString("UserThreadSubscription(")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uts.UserID))
	builder.WriteString(", ")
	builder.WriteString("thread_id=")
	builder.WriteString(fmt.Sprintf("%v", uts.ThreadID))
	builder.WriteString(", ")
	builder.WriteString("is_notified=")
	builder.WriteString(fmt.Sprintf("%v", uts.IsNotified))
	builder.WriteString(", ")
	builder.WriteString("is_checked=")
	builder.WriteString(fmt.Sprintf("%v", uts.IsChecked))
	builder.WriteString(", ")
	builder.WriteString("subscribed_at=")
	builder.WriteString(uts.SubscribedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserThreadSubscriptions is a parsable slice of UserThreadSubscription.
type UserThreadSubscriptions []*UserThreadSubscription
