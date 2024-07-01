// Code generated by ent, DO NOT EDIT.

package threadtag

import (
	"server/infrastructure/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ThreadId applies equality check predicate on the "threadId" field. It's identical to ThreadIdEQ.
func ThreadId(v int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldEQ(FieldThreadId, v))
}

// TagId applies equality check predicate on the "tagId" field. It's identical to TagIdEQ.
func TagId(v int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldEQ(FieldTagId, v))
}

// ThreadIdEQ applies the EQ predicate on the "threadId" field.
func ThreadIdEQ(v int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldEQ(FieldThreadId, v))
}

// ThreadIdNEQ applies the NEQ predicate on the "threadId" field.
func ThreadIdNEQ(v int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldNEQ(FieldThreadId, v))
}

// ThreadIdIn applies the In predicate on the "threadId" field.
func ThreadIdIn(vs ...int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldIn(FieldThreadId, vs...))
}

// ThreadIdNotIn applies the NotIn predicate on the "threadId" field.
func ThreadIdNotIn(vs ...int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldNotIn(FieldThreadId, vs...))
}

// TagIdEQ applies the EQ predicate on the "tagId" field.
func TagIdEQ(v int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldEQ(FieldTagId, v))
}

// TagIdNEQ applies the NEQ predicate on the "tagId" field.
func TagIdNEQ(v int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldNEQ(FieldTagId, v))
}

// TagIdIn applies the In predicate on the "tagId" field.
func TagIdIn(vs ...int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldIn(FieldTagId, vs...))
}

// TagIdNotIn applies the NotIn predicate on the "tagId" field.
func TagIdNotIn(vs ...int) predicate.ThreadTag {
	return predicate.ThreadTag(sql.FieldNotIn(FieldTagId, vs...))
}

// HasThread applies the HasEdge predicate on the "thread" edge.
func HasThread() predicate.ThreadTag {
	return predicate.ThreadTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ThreadColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ThreadTable, ThreadColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThreadWith applies the HasEdge predicate on the "thread" edge with a given conditions (other predicates).
func HasThreadWith(preds ...predicate.Thread) predicate.ThreadTag {
	return predicate.ThreadTag(func(s *sql.Selector) {
		step := newThreadStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.ThreadTag {
	return predicate.ThreadTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, TagColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.ThreadTag {
	return predicate.ThreadTag(func(s *sql.Selector) {
		step := newTagStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ThreadTag) predicate.ThreadTag {
	return predicate.ThreadTag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ThreadTag) predicate.ThreadTag {
	return predicate.ThreadTag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ThreadTag) predicate.ThreadTag {
	return predicate.ThreadTag(sql.NotPredicates(p))
}
