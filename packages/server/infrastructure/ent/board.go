// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Board is the model entity for the Board schema.
type Board struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// 板名
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// 0: Public, 1: Private, 3: Pending, 3: Archived
	Status int `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BoardQuery when eager-loading is set.
	Edges        BoardEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BoardEdges holds the relations/edges for other nodes in the graph.
type BoardEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Threads holds the value of the threads edge.
	Threads []*Thread `json:"threads,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BoardEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ThreadsOrErr returns the Threads value or an error if the edge
// was not loaded in eager-loading.
func (e BoardEdges) ThreadsOrErr() ([]*Thread, error) {
	if e.loadedTypes[1] {
		return e.Threads, nil
	}
	return nil, &NotLoadedError{edge: "threads"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Board) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case board.FieldID, board.FieldUserID, board.FieldStatus:
			values[i] = new(sql.NullInt64)
		case board.FieldTitle, board.FieldDescription:
			values[i] = new(sql.NullString)
		case board.FieldCreatedAt, board.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Board fields.
func (b *Board) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case board.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case board.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				b.UserID = int(value.Int64)
			}
		case board.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case board.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = new(string)
				*b.Description = value.String
			}
		case board.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = int(value.Int64)
			}
		case board.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case board.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Board.
// This includes values selected through modifiers, order, etc.
func (b *Board) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Board entity.
func (b *Board) QueryOwner() *UserQuery {
	return NewBoardClient(b.config).QueryOwner(b)
}

// QueryThreads queries the "threads" edge of the Board entity.
func (b *Board) QueryThreads() *ThreadQuery {
	return NewBoardClient(b.config).QueryThreads(b)
}

// Update returns a builder for updating this Board.
// Note that you need to call Board.Unwrap() before calling this method if this Board
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Board) Update() *BoardUpdateOne {
	return NewBoardClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Board entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Board) Unwrap() *Board {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Board is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Board) String() string {
	var builder strings.Builder
	builder.WriteString("Board(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", b.UserID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(b.Title)
	builder.WriteString(", ")
	if v := b.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Boards is a parsable slice of Board.
type Boards []*Board
