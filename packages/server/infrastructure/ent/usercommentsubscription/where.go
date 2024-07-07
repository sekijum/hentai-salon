// Code generated by ent, DO NOT EDIT.

package usercommentsubscription

import (
	"server/infrastructure/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldUserID, v))
}

// CommentID applies equality check predicate on the "comment_id" field. It's identical to CommentIDEQ.
func CommentID(v int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldCommentID, v))
}

// IsNotified applies equality check predicate on the "is_notified" field. It's identical to IsNotifiedEQ.
func IsNotified(v bool) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldIsNotified, v))
}

// IsChecked applies equality check predicate on the "is_checked" field. It's identical to IsCheckedEQ.
func IsChecked(v bool) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldIsChecked, v))
}

// SubscribedAt applies equality check predicate on the "subscribed_at" field. It's identical to SubscribedAtEQ.
func SubscribedAt(v time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldSubscribedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNotIn(FieldUserID, vs...))
}

// CommentIDEQ applies the EQ predicate on the "comment_id" field.
func CommentIDEQ(v int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldCommentID, v))
}

// CommentIDNEQ applies the NEQ predicate on the "comment_id" field.
func CommentIDNEQ(v int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNEQ(FieldCommentID, v))
}

// CommentIDIn applies the In predicate on the "comment_id" field.
func CommentIDIn(vs ...int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldIn(FieldCommentID, vs...))
}

// CommentIDNotIn applies the NotIn predicate on the "comment_id" field.
func CommentIDNotIn(vs ...int) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNotIn(FieldCommentID, vs...))
}

// IsNotifiedEQ applies the EQ predicate on the "is_notified" field.
func IsNotifiedEQ(v bool) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldIsNotified, v))
}

// IsNotifiedNEQ applies the NEQ predicate on the "is_notified" field.
func IsNotifiedNEQ(v bool) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNEQ(FieldIsNotified, v))
}

// IsCheckedEQ applies the EQ predicate on the "is_checked" field.
func IsCheckedEQ(v bool) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldIsChecked, v))
}

// IsCheckedNEQ applies the NEQ predicate on the "is_checked" field.
func IsCheckedNEQ(v bool) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNEQ(FieldIsChecked, v))
}

// SubscribedAtEQ applies the EQ predicate on the "subscribed_at" field.
func SubscribedAtEQ(v time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldEQ(FieldSubscribedAt, v))
}

// SubscribedAtNEQ applies the NEQ predicate on the "subscribed_at" field.
func SubscribedAtNEQ(v time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNEQ(FieldSubscribedAt, v))
}

// SubscribedAtIn applies the In predicate on the "subscribed_at" field.
func SubscribedAtIn(vs ...time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldIn(FieldSubscribedAt, vs...))
}

// SubscribedAtNotIn applies the NotIn predicate on the "subscribed_at" field.
func SubscribedAtNotIn(vs ...time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldNotIn(FieldSubscribedAt, vs...))
}

// SubscribedAtGT applies the GT predicate on the "subscribed_at" field.
func SubscribedAtGT(v time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldGT(FieldSubscribedAt, v))
}

// SubscribedAtGTE applies the GTE predicate on the "subscribed_at" field.
func SubscribedAtGTE(v time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldGTE(FieldSubscribedAt, v))
}

// SubscribedAtLT applies the LT predicate on the "subscribed_at" field.
func SubscribedAtLT(v time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldLT(FieldSubscribedAt, v))
}

// SubscribedAtLTE applies the LTE predicate on the "subscribed_at" field.
func SubscribedAtLTE(v time.Time) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.FieldLTE(FieldSubscribedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, CommentColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.ThreadComment) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(func(s *sql.Selector) {
		step := newCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserCommentSubscription) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserCommentSubscription) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserCommentSubscription) predicate.UserCommentSubscription {
	return predicate.UserCommentSubscription(sql.NotPredicates(p))
}
