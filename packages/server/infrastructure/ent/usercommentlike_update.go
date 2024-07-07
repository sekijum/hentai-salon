// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usercommentlike"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCommentLikeUpdate is the builder for updating UserCommentLike entities.
type UserCommentLikeUpdate struct {
	config
	hooks    []Hook
	mutation *UserCommentLikeMutation
}

// Where appends a list predicates to the UserCommentLikeUpdate builder.
func (uclu *UserCommentLikeUpdate) Where(ps ...predicate.UserCommentLike) *UserCommentLikeUpdate {
	uclu.mutation.Where(ps...)
	return uclu
}

// SetUserID sets the "user_id" field.
func (uclu *UserCommentLikeUpdate) SetUserID(i int) *UserCommentLikeUpdate {
	uclu.mutation.SetUserID(i)
	return uclu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uclu *UserCommentLikeUpdate) SetNillableUserID(i *int) *UserCommentLikeUpdate {
	if i != nil {
		uclu.SetUserID(*i)
	}
	return uclu
}

// SetCommentID sets the "comment_id" field.
func (uclu *UserCommentLikeUpdate) SetCommentID(i int) *UserCommentLikeUpdate {
	uclu.mutation.SetCommentID(i)
	return uclu
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (uclu *UserCommentLikeUpdate) SetNillableCommentID(i *int) *UserCommentLikeUpdate {
	if i != nil {
		uclu.SetCommentID(*i)
	}
	return uclu
}

// SetLikedAt sets the "liked_at" field.
func (uclu *UserCommentLikeUpdate) SetLikedAt(t time.Time) *UserCommentLikeUpdate {
	uclu.mutation.SetLikedAt(t)
	return uclu
}

// SetNillableLikedAt sets the "liked_at" field if the given value is not nil.
func (uclu *UserCommentLikeUpdate) SetNillableLikedAt(t *time.Time) *UserCommentLikeUpdate {
	if t != nil {
		uclu.SetLikedAt(*t)
	}
	return uclu
}

// SetUser sets the "user" edge to the User entity.
func (uclu *UserCommentLikeUpdate) SetUser(u *User) *UserCommentLikeUpdate {
	return uclu.SetUserID(u.ID)
}

// SetComment sets the "comment" edge to the ThreadComment entity.
func (uclu *UserCommentLikeUpdate) SetComment(t *ThreadComment) *UserCommentLikeUpdate {
	return uclu.SetCommentID(t.ID)
}

// Mutation returns the UserCommentLikeMutation object of the builder.
func (uclu *UserCommentLikeUpdate) Mutation() *UserCommentLikeMutation {
	return uclu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uclu *UserCommentLikeUpdate) ClearUser() *UserCommentLikeUpdate {
	uclu.mutation.ClearUser()
	return uclu
}

// ClearComment clears the "comment" edge to the ThreadComment entity.
func (uclu *UserCommentLikeUpdate) ClearComment() *UserCommentLikeUpdate {
	uclu.mutation.ClearComment()
	return uclu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uclu *UserCommentLikeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uclu.sqlSave, uclu.mutation, uclu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uclu *UserCommentLikeUpdate) SaveX(ctx context.Context) int {
	affected, err := uclu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uclu *UserCommentLikeUpdate) Exec(ctx context.Context) error {
	_, err := uclu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uclu *UserCommentLikeUpdate) ExecX(ctx context.Context) {
	if err := uclu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uclu *UserCommentLikeUpdate) check() error {
	if _, ok := uclu.mutation.UserID(); uclu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.user"`)
	}
	if _, ok := uclu.mutation.CommentID(); uclu.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.comment"`)
	}
	return nil
}

func (uclu *UserCommentLikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uclu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercommentlike.Table, usercommentlike.Columns, sqlgraph.NewFieldSpec(usercommentlike.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(usercommentlike.FieldCommentID, field.TypeInt))
	if ps := uclu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uclu.mutation.LikedAt(); ok {
		_spec.SetField(usercommentlike.FieldLikedAt, field.TypeTime, value)
	}
	if uclu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uclu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uclu.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uclu.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uclu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercommentlike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uclu.mutation.done = true
	return n, nil
}

// UserCommentLikeUpdateOne is the builder for updating a single UserCommentLike entity.
type UserCommentLikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCommentLikeMutation
}

// SetUserID sets the "user_id" field.
func (ucluo *UserCommentLikeUpdateOne) SetUserID(i int) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetUserID(i)
	return ucluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucluo *UserCommentLikeUpdateOne) SetNillableUserID(i *int) *UserCommentLikeUpdateOne {
	if i != nil {
		ucluo.SetUserID(*i)
	}
	return ucluo
}

// SetCommentID sets the "comment_id" field.
func (ucluo *UserCommentLikeUpdateOne) SetCommentID(i int) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetCommentID(i)
	return ucluo
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (ucluo *UserCommentLikeUpdateOne) SetNillableCommentID(i *int) *UserCommentLikeUpdateOne {
	if i != nil {
		ucluo.SetCommentID(*i)
	}
	return ucluo
}

// SetLikedAt sets the "liked_at" field.
func (ucluo *UserCommentLikeUpdateOne) SetLikedAt(t time.Time) *UserCommentLikeUpdateOne {
	ucluo.mutation.SetLikedAt(t)
	return ucluo
}

// SetNillableLikedAt sets the "liked_at" field if the given value is not nil.
func (ucluo *UserCommentLikeUpdateOne) SetNillableLikedAt(t *time.Time) *UserCommentLikeUpdateOne {
	if t != nil {
		ucluo.SetLikedAt(*t)
	}
	return ucluo
}

// SetUser sets the "user" edge to the User entity.
func (ucluo *UserCommentLikeUpdateOne) SetUser(u *User) *UserCommentLikeUpdateOne {
	return ucluo.SetUserID(u.ID)
}

// SetComment sets the "comment" edge to the ThreadComment entity.
func (ucluo *UserCommentLikeUpdateOne) SetComment(t *ThreadComment) *UserCommentLikeUpdateOne {
	return ucluo.SetCommentID(t.ID)
}

// Mutation returns the UserCommentLikeMutation object of the builder.
func (ucluo *UserCommentLikeUpdateOne) Mutation() *UserCommentLikeMutation {
	return ucluo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucluo *UserCommentLikeUpdateOne) ClearUser() *UserCommentLikeUpdateOne {
	ucluo.mutation.ClearUser()
	return ucluo
}

// ClearComment clears the "comment" edge to the ThreadComment entity.
func (ucluo *UserCommentLikeUpdateOne) ClearComment() *UserCommentLikeUpdateOne {
	ucluo.mutation.ClearComment()
	return ucluo
}

// Where appends a list predicates to the UserCommentLikeUpdate builder.
func (ucluo *UserCommentLikeUpdateOne) Where(ps ...predicate.UserCommentLike) *UserCommentLikeUpdateOne {
	ucluo.mutation.Where(ps...)
	return ucluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucluo *UserCommentLikeUpdateOne) Select(field string, fields ...string) *UserCommentLikeUpdateOne {
	ucluo.fields = append([]string{field}, fields...)
	return ucluo
}

// Save executes the query and returns the updated UserCommentLike entity.
func (ucluo *UserCommentLikeUpdateOne) Save(ctx context.Context) (*UserCommentLike, error) {
	return withHooks(ctx, ucluo.sqlSave, ucluo.mutation, ucluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucluo *UserCommentLikeUpdateOne) SaveX(ctx context.Context) *UserCommentLike {
	node, err := ucluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucluo *UserCommentLikeUpdateOne) Exec(ctx context.Context) error {
	_, err := ucluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucluo *UserCommentLikeUpdateOne) ExecX(ctx context.Context) {
	if err := ucluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucluo *UserCommentLikeUpdateOne) check() error {
	if _, ok := ucluo.mutation.UserID(); ucluo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.user"`)
	}
	if _, ok := ucluo.mutation.CommentID(); ucluo.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentLike.comment"`)
	}
	return nil
}

func (ucluo *UserCommentLikeUpdateOne) sqlSave(ctx context.Context) (_node *UserCommentLike, err error) {
	if err := ucluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercommentlike.Table, usercommentlike.Columns, sqlgraph.NewFieldSpec(usercommentlike.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(usercommentlike.FieldCommentID, field.TypeInt))
	if id, ok := ucluo.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "UserCommentLike.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ucluo.mutation.CommentID(); !ok {
		return nil, &ValidationError{Name: "comment_id", err: errors.New(`ent: missing "UserCommentLike.comment_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ucluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !usercommentlike.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ucluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucluo.mutation.LikedAt(); ok {
		_spec.SetField(usercommentlike.FieldLikedAt, field.TypeTime, value)
	}
	if ucluo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucluo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.UserTable,
			Columns: []string{usercommentlike.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ucluo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucluo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserCommentLike{config: ucluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercommentlike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucluo.mutation.done = true
	return _node, nil
}
