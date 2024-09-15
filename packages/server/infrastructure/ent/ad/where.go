// Code generated by ent, DO NOT EDIT.

package ad

import (
	"server/infrastructure/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ad {
	return predicate.Ad(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ad {
	return predicate.Ad(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Ad {
	return predicate.Ad(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Ad {
	return predicate.Ad(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ad {
	return predicate.Ad(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ad {
	return predicate.Ad(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ad {
	return predicate.Ad(sql.FieldLTE(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldContent, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v int) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldIsActive, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldUpdatedAt, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Ad {
	return predicate.Ad(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Ad {
	return predicate.Ad(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Ad {
	return predicate.Ad(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Ad {
	return predicate.Ad(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Ad {
	return predicate.Ad(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Ad {
	return predicate.Ad(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Ad {
	return predicate.Ad(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Ad {
	return predicate.Ad(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Ad {
	return predicate.Ad(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Ad {
	return predicate.Ad(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Ad {
	return predicate.Ad(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Ad {
	return predicate.Ad(sql.FieldContainsFold(FieldContent, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v int) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v int) predicate.Ad {
	return predicate.Ad(sql.FieldNEQ(FieldIsActive, v))
}

// IsActiveIn applies the In predicate on the "is_active" field.
func IsActiveIn(vs ...int) predicate.Ad {
	return predicate.Ad(sql.FieldIn(FieldIsActive, vs...))
}

// IsActiveNotIn applies the NotIn predicate on the "is_active" field.
func IsActiveNotIn(vs ...int) predicate.Ad {
	return predicate.Ad(sql.FieldNotIn(FieldIsActive, vs...))
}

// IsActiveGT applies the GT predicate on the "is_active" field.
func IsActiveGT(v int) predicate.Ad {
	return predicate.Ad(sql.FieldGT(FieldIsActive, v))
}

// IsActiveGTE applies the GTE predicate on the "is_active" field.
func IsActiveGTE(v int) predicate.Ad {
	return predicate.Ad(sql.FieldGTE(FieldIsActive, v))
}

// IsActiveLT applies the LT predicate on the "is_active" field.
func IsActiveLT(v int) predicate.Ad {
	return predicate.Ad(sql.FieldLT(FieldIsActive, v))
}

// IsActiveLTE applies the LTE predicate on the "is_active" field.
func IsActiveLTE(v int) predicate.Ad {
	return predicate.Ad(sql.FieldLTE(FieldIsActive, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Ad {
	return predicate.Ad(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ad) predicate.Ad {
	return predicate.Ad(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ad) predicate.Ad {
	return predicate.Ad(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ad) predicate.Ad {
	return predicate.Ad(sql.NotPredicates(p))
}
