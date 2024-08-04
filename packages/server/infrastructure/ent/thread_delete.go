// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/thread"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThreadDelete is the builder for deleting a Thread entity.
type ThreadDelete struct {
	config
	hooks    []Hook
	mutation *ThreadMutation
}

// Where appends a list predicates to the ThreadDelete builder.
func (td *ThreadDelete) Where(ps ...predicate.Thread) *ThreadDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *ThreadDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *ThreadDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *ThreadDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(thread.Table, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// ThreadDeleteOne is the builder for deleting a single Thread entity.
type ThreadDeleteOne struct {
	td *ThreadDelete
}

// Where appends a list predicates to the ThreadDelete builder.
func (tdo *ThreadDeleteOne) Where(ps ...predicate.Thread) *ThreadDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *ThreadDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{thread.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *ThreadDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}
