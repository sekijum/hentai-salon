// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/userthreadlike"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// UserThreadLikeDelete is the builder for deleting a UserThreadLike entity.
type UserThreadLikeDelete struct {
	config
	hooks    []Hook
	mutation *UserThreadLikeMutation
}

// Where appends a list predicates to the UserThreadLikeDelete builder.
func (utld *UserThreadLikeDelete) Where(ps ...predicate.UserThreadLike) *UserThreadLikeDelete {
	utld.mutation.Where(ps...)
	return utld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (utld *UserThreadLikeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, utld.sqlExec, utld.mutation, utld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (utld *UserThreadLikeDelete) ExecX(ctx context.Context) int {
	n, err := utld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (utld *UserThreadLikeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userthreadlike.Table, nil)
	if ps := utld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, utld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	utld.mutation.done = true
	return affected, err
}

// UserThreadLikeDeleteOne is the builder for deleting a single UserThreadLike entity.
type UserThreadLikeDeleteOne struct {
	utld *UserThreadLikeDelete
}

// Where appends a list predicates to the UserThreadLikeDelete builder.
func (utldo *UserThreadLikeDeleteOne) Where(ps ...predicate.UserThreadLike) *UserThreadLikeDeleteOne {
	utldo.utld.mutation.Where(ps...)
	return utldo
}

// Exec executes the deletion query.
func (utldo *UserThreadLikeDeleteOne) Exec(ctx context.Context) error {
	n, err := utldo.utld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userthreadlike.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (utldo *UserThreadLikeDeleteOne) ExecX(ctx context.Context) {
	if err := utldo.Exec(ctx); err != nil {
		panic(err)
	}
}
