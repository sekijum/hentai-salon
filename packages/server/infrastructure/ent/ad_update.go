// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/ad"
	"server/infrastructure/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdUpdate is the builder for updating Ad entities.
type AdUpdate struct {
	config
	hooks    []Hook
	mutation *AdMutation
}

// Where appends a list predicates to the AdUpdate builder.
func (au *AdUpdate) Where(ps ...predicate.Ad) *AdUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetContent sets the "content" field.
func (au *AdUpdate) SetContent(s string) *AdUpdate {
	au.mutation.SetContent(s)
	return au
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (au *AdUpdate) SetNillableContent(s *string) *AdUpdate {
	if s != nil {
		au.SetContent(*s)
	}
	return au
}

// SetIsActive sets the "is_active" field.
func (au *AdUpdate) SetIsActive(i int) *AdUpdate {
	au.mutation.ResetIsActive()
	au.mutation.SetIsActive(i)
	return au
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (au *AdUpdate) SetNillableIsActive(i *int) *AdUpdate {
	if i != nil {
		au.SetIsActive(*i)
	}
	return au
}

// AddIsActive adds i to the "is_active" field.
func (au *AdUpdate) AddIsActive(i int) *AdUpdate {
	au.mutation.AddIsActive(i)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AdUpdate) SetCreatedAt(t time.Time) *AdUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AdUpdate) SetNillableCreatedAt(t *time.Time) *AdUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AdUpdate) SetUpdatedAt(t time.Time) *AdUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// Mutation returns the AdMutation object of the builder.
func (au *AdUpdate) Mutation() *AdMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AdUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := ad.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AdUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ad.Table, ad.Columns, sqlgraph.NewFieldSpec(ad.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Content(); ok {
		_spec.SetField(ad.FieldContent, field.TypeString, value)
	}
	if value, ok := au.mutation.IsActive(); ok {
		_spec.SetField(ad.FieldIsActive, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedIsActive(); ok {
		_spec.AddField(ad.FieldIsActive, field.TypeInt, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(ad.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(ad.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ad.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AdUpdateOne is the builder for updating a single Ad entity.
type AdUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdMutation
}

// SetContent sets the "content" field.
func (auo *AdUpdateOne) SetContent(s string) *AdUpdateOne {
	auo.mutation.SetContent(s)
	return auo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (auo *AdUpdateOne) SetNillableContent(s *string) *AdUpdateOne {
	if s != nil {
		auo.SetContent(*s)
	}
	return auo
}

// SetIsActive sets the "is_active" field.
func (auo *AdUpdateOne) SetIsActive(i int) *AdUpdateOne {
	auo.mutation.ResetIsActive()
	auo.mutation.SetIsActive(i)
	return auo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (auo *AdUpdateOne) SetNillableIsActive(i *int) *AdUpdateOne {
	if i != nil {
		auo.SetIsActive(*i)
	}
	return auo
}

// AddIsActive adds i to the "is_active" field.
func (auo *AdUpdateOne) AddIsActive(i int) *AdUpdateOne {
	auo.mutation.AddIsActive(i)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AdUpdateOne) SetCreatedAt(t time.Time) *AdUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AdUpdateOne) SetNillableCreatedAt(t *time.Time) *AdUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AdUpdateOne) SetUpdatedAt(t time.Time) *AdUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// Mutation returns the AdMutation object of the builder.
func (auo *AdUpdateOne) Mutation() *AdMutation {
	return auo.mutation
}

// Where appends a list predicates to the AdUpdate builder.
func (auo *AdUpdateOne) Where(ps ...predicate.Ad) *AdUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdUpdateOne) Select(field string, fields ...string) *AdUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Ad entity.
func (auo *AdUpdateOne) Save(ctx context.Context) (*Ad, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdUpdateOne) SaveX(ctx context.Context) *Ad {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AdUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := ad.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AdUpdateOne) sqlSave(ctx context.Context) (_node *Ad, err error) {
	_spec := sqlgraph.NewUpdateSpec(ad.Table, ad.Columns, sqlgraph.NewFieldSpec(ad.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ad.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ad.FieldID)
		for _, f := range fields {
			if !ad.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ad.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Content(); ok {
		_spec.SetField(ad.FieldContent, field.TypeString, value)
	}
	if value, ok := auo.mutation.IsActive(); ok {
		_spec.SetField(ad.FieldIsActive, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedIsActive(); ok {
		_spec.AddField(ad.FieldIsActive, field.TypeInt, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(ad.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(ad.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Ad{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ad.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
