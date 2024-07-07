// Code generated by ent, DO NOT EDIT.

package userthreadsubscription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the userthreadsubscription type in the database.
	Label = "user_thread_subscription"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldThreadID holds the string denoting the thread_id field in the database.
	FieldThreadID = "thread_id"
	// FieldIsNotified holds the string denoting the is_notified field in the database.
	FieldIsNotified = "is_notified"
	// FieldIsChecked holds the string denoting the is_checked field in the database.
	FieldIsChecked = "is_checked"
	// FieldSubscribedAt holds the string denoting the subscribed_at field in the database.
	FieldSubscribedAt = "subscribed_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeThread holds the string denoting the thread edge name in mutations.
	EdgeThread = "thread"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// ThreadFieldID holds the string denoting the ID field of the Thread.
	ThreadFieldID = "id"
	// Table holds the table name of the userthreadsubscription in the database.
	Table = "user_thread_subscriptions"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_thread_subscriptions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// ThreadTable is the table that holds the thread relation/edge.
	ThreadTable = "user_thread_subscriptions"
	// ThreadInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadInverseTable = "threads"
	// ThreadColumn is the table column denoting the thread relation/edge.
	ThreadColumn = "thread_id"
)

// Columns holds all SQL columns for userthreadsubscription fields.
var Columns = []string{
	FieldUserID,
	FieldThreadID,
	FieldIsNotified,
	FieldIsChecked,
	FieldSubscribedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsNotified holds the default value on creation for the "is_notified" field.
	DefaultIsNotified bool
	// DefaultIsChecked holds the default value on creation for the "is_checked" field.
	DefaultIsChecked bool
	// DefaultSubscribedAt holds the default value on creation for the "subscribed_at" field.
	DefaultSubscribedAt func() time.Time
)

// OrderOption defines the ordering options for the UserThreadSubscription queries.
type OrderOption func(*sql.Selector)

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByThreadID orders the results by the thread_id field.
func ByThreadID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreadID, opts...).ToFunc()
}

// ByIsNotified orders the results by the is_notified field.
func ByIsNotified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsNotified, opts...).ToFunc()
}

// ByIsChecked orders the results by the is_checked field.
func ByIsChecked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsChecked, opts...).ToFunc()
}

// BySubscribedAt orders the results by the subscribed_at field.
func BySubscribedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscribedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByThreadField orders the results by thread field.
func ByThreadField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, UserColumn),
		sqlgraph.To(UserInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
func newThreadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, ThreadColumn),
		sqlgraph.To(ThreadInverseTable, ThreadFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ThreadTable, ThreadColumn),
	)
}
