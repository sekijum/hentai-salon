// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/tag"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/threadtag"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ThreadTag is the model entity for the ThreadTag schema.
type ThreadTag struct {
	config `json:"-"`
	// ThreadID holds the value of the "thread_id" field.
	ThreadID int `json:"thread_id,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID int `json:"tag_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ThreadTagQuery when eager-loading is set.
	Edges        ThreadTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ThreadTagEdges holds the relations/edges for other nodes in the graph.
type ThreadTagEdges struct {
	// Thread holds the value of the thread edge.
	Thread *Thread `json:"thread,omitempty"`
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ThreadOrErr returns the Thread value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadTagEdges) ThreadOrErr() (*Thread, error) {
	if e.Thread != nil {
		return e.Thread, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: thread.Label}
	}
	return nil, &NotLoadedError{edge: "thread"}
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreadTagEdges) TagOrErr() (*Tag, error) {
	if e.Tag != nil {
		return e.Tag, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: tag.Label}
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ThreadTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case threadtag.FieldThreadID, threadtag.FieldTagID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ThreadTag fields.
func (tt *ThreadTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case threadtag.FieldThreadID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field thread_id", values[i])
			} else if value.Valid {
				tt.ThreadID = int(value.Int64)
			}
		case threadtag.FieldTagID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value.Valid {
				tt.TagID = int(value.Int64)
			}
		default:
			tt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ThreadTag.
// This includes values selected through modifiers, order, etc.
func (tt *ThreadTag) Value(name string) (ent.Value, error) {
	return tt.selectValues.Get(name)
}

// QueryThread queries the "thread" edge of the ThreadTag entity.
func (tt *ThreadTag) QueryThread() *ThreadQuery {
	return NewThreadTagClient(tt.config).QueryThread(tt)
}

// QueryTag queries the "tag" edge of the ThreadTag entity.
func (tt *ThreadTag) QueryTag() *TagQuery {
	return NewThreadTagClient(tt.config).QueryTag(tt)
}

// Update returns a builder for updating this ThreadTag.
// Note that you need to call ThreadTag.Unwrap() before calling this method if this ThreadTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *ThreadTag) Update() *ThreadTagUpdateOne {
	return NewThreadTagClient(tt.config).UpdateOne(tt)
}

// Unwrap unwraps the ThreadTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *ThreadTag) Unwrap() *ThreadTag {
	_tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ThreadTag is not a transactional entity")
	}
	tt.config.driver = _tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *ThreadTag) String() string {
	var builder strings.Builder
	builder.WriteString("ThreadTag(")
	builder.WriteString("thread_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.ThreadID))
	builder.WriteString(", ")
	builder.WriteString("tag_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.TagID))
	builder.WriteByte(')')
	return builder.String()
}

// ThreadTags is a parsable slice of ThreadTag.
type ThreadTags []*ThreadTag
