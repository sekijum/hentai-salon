// Code generated by ent, DO NOT EDIT.

package thread

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the thread type in the database.
	Label = "thread"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBoardId holds the string denoting the boardid field in the database.
	FieldBoardId = "board_id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldThumbnailUrl holds the string denoting the thumbnailurl field in the database.
	FieldThumbnailUrl = "thumbnail_url"
	// FieldIsNotifyOnComment holds the string denoting the isnotifyoncomment field in the database.
	FieldIsNotifyOnComment = "is_notify_on_comment"
	// FieldIPAddress holds the string denoting the ip_address field in the database.
	FieldIPAddress = "ip_address"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeBoard holds the string denoting the board edge name in mutations.
	EdgeBoard = "board"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeLikedUsers holds the string denoting the liked_users edge name in mutations.
	EdgeLikedUsers = "liked_users"
	// EdgeSubscribedUsers holds the string denoting the subscribed_users edge name in mutations.
	EdgeSubscribedUsers = "subscribed_users"
	// EdgeThreadTags holds the string denoting the thread_tags edge name in mutations.
	EdgeThreadTags = "thread_tags"
	// EdgeUserThreadLike holds the string denoting the user_thread_like edge name in mutations.
	EdgeUserThreadLike = "user_thread_like"
	// EdgeUserThreadSubscription holds the string denoting the user_thread_subscription edge name in mutations.
	EdgeUserThreadSubscription = "user_thread_subscription"
	// Table holds the table name of the thread in the database.
	Table = "threads"
	// BoardTable is the table that holds the board relation/edge.
	BoardTable = "threads"
	// BoardInverseTable is the table name for the Board entity.
	// It exists in this package in order to avoid circular dependency with the "board" package.
	BoardInverseTable = "boards"
	// BoardColumn is the table column denoting the board relation/edge.
	BoardColumn = "board_id"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "threads"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "thread_comments"
	// CommentsInverseTable is the table name for the ThreadComment entity.
	// It exists in this package in order to avoid circular dependency with the "threadcomment" package.
	CommentsInverseTable = "thread_comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "thread_id"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "thread_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// LikedUsersTable is the table that holds the liked_users relation/edge. The primary key declared below.
	LikedUsersTable = "user_thread_likes"
	// LikedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	LikedUsersInverseTable = "users"
	// SubscribedUsersTable is the table that holds the subscribed_users relation/edge. The primary key declared below.
	SubscribedUsersTable = "user_thread_subscriptions"
	// SubscribedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SubscribedUsersInverseTable = "users"
	// ThreadTagsTable is the table that holds the thread_tags relation/edge.
	ThreadTagsTable = "thread_tags"
	// ThreadTagsInverseTable is the table name for the ThreadTag entity.
	// It exists in this package in order to avoid circular dependency with the "threadtag" package.
	ThreadTagsInverseTable = "thread_tags"
	// ThreadTagsColumn is the table column denoting the thread_tags relation/edge.
	ThreadTagsColumn = "thread_id"
	// UserThreadLikeTable is the table that holds the user_thread_like relation/edge.
	UserThreadLikeTable = "user_thread_likes"
	// UserThreadLikeInverseTable is the table name for the UserThreadLike entity.
	// It exists in this package in order to avoid circular dependency with the "userthreadlike" package.
	UserThreadLikeInverseTable = "user_thread_likes"
	// UserThreadLikeColumn is the table column denoting the user_thread_like relation/edge.
	UserThreadLikeColumn = "thread_id"
	// UserThreadSubscriptionTable is the table that holds the user_thread_subscription relation/edge.
	UserThreadSubscriptionTable = "user_thread_subscriptions"
	// UserThreadSubscriptionInverseTable is the table name for the UserThreadSubscription entity.
	// It exists in this package in order to avoid circular dependency with the "userthreadsubscription" package.
	UserThreadSubscriptionInverseTable = "user_thread_subscriptions"
	// UserThreadSubscriptionColumn is the table column denoting the user_thread_subscription relation/edge.
	UserThreadSubscriptionColumn = "thread_id"
)

// Columns holds all SQL columns for thread fields.
var Columns = []string{
	FieldID,
	FieldBoardId,
	FieldUserId,
	FieldTitle,
	FieldDescription,
	FieldThumbnailUrl,
	FieldIsNotifyOnComment,
	FieldIPAddress,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"threadId", "tagId"}
	// LikedUsersPrimaryKey and LikedUsersColumn2 are the table columns denoting the
	// primary key for the liked_users relation (M2M).
	LikedUsersPrimaryKey = []string{"userId", "threadId"}
	// SubscribedUsersPrimaryKey and SubscribedUsersColumn2 are the table columns denoting the
	// primary key for the subscribed_users relation (M2M).
	SubscribedUsersPrimaryKey = []string{"userId", "threadId"}
)

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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultIsNotifyOnComment holds the default value on creation for the "isNotifyOnComment" field.
	DefaultIsNotifyOnComment bool
	// IPAddressValidator is a validator for the "ip_address" field. It is called by the builders before save.
	IPAddressValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Thread queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBoardId orders the results by the boardId field.
func ByBoardId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBoardId, opts...).ToFunc()
}

// ByUserId orders the results by the userId field.
func ByUserId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserId, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByThumbnailUrl orders the results by the thumbnailUrl field.
func ByThumbnailUrl(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThumbnailUrl, opts...).ToFunc()
}

// ByIsNotifyOnComment orders the results by the isNotifyOnComment field.
func ByIsNotifyOnComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsNotifyOnComment, opts...).ToFunc()
}

// ByIPAddress orders the results by the ip_address field.
func ByIPAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPAddress, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBoardField orders the results by board field.
func ByBoardField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBoardStep(), sql.OrderByField(field, opts...))
	}
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentsCount orders the results by comments count.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsStep(), opts...)
	}
}

// ByComments orders the results by comments terms.
func ByComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikedUsersCount orders the results by liked_users count.
func ByLikedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikedUsersStep(), opts...)
	}
}

// ByLikedUsers orders the results by liked_users terms.
func ByLikedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubscribedUsersCount orders the results by subscribed_users count.
func BySubscribedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscribedUsersStep(), opts...)
	}
}

// BySubscribedUsers orders the results by subscribed_users terms.
func BySubscribedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscribedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByThreadTagsCount orders the results by thread_tags count.
func ByThreadTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newThreadTagsStep(), opts...)
	}
}

// ByThreadTags orders the results by thread_tags terms.
func ByThreadTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserThreadLikeCount orders the results by user_thread_like count.
func ByUserThreadLikeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserThreadLikeStep(), opts...)
	}
}

// ByUserThreadLike orders the results by user_thread_like terms.
func ByUserThreadLike(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserThreadLikeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserThreadSubscriptionCount orders the results by user_thread_subscription count.
func ByUserThreadSubscriptionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserThreadSubscriptionStep(), opts...)
	}
}

// ByUserThreadSubscription orders the results by user_thread_subscription terms.
func ByUserThreadSubscription(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserThreadSubscriptionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBoardStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BoardInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BoardTable, BoardColumn),
	)
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
func newLikedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, LikedUsersTable, LikedUsersPrimaryKey...),
	)
}
func newSubscribedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscribedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SubscribedUsersTable, SubscribedUsersPrimaryKey...),
	)
}
func newThreadTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadTagsInverseTable, ThreadTagsColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, ThreadTagsTable, ThreadTagsColumn),
	)
}
func newUserThreadLikeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserThreadLikeInverseTable, UserThreadLikeColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserThreadLikeTable, UserThreadLikeColumn),
	)
}
func newUserThreadSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserThreadSubscriptionInverseTable, UserThreadSubscriptionColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserThreadSubscriptionTable, UserThreadSubscriptionColumn),
	)
}
