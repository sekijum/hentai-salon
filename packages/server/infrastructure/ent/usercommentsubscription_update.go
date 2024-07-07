// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usercommentsubscription"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCommentSubscriptionUpdate is the builder for updating UserCommentSubscription entities.
type UserCommentSubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *UserCommentSubscriptionMutation
}

// Where appends a list predicates to the UserCommentSubscriptionUpdate builder.
func (ucsu *UserCommentSubscriptionUpdate) Where(ps ...predicate.UserCommentSubscription) *UserCommentSubscriptionUpdate {
	ucsu.mutation.Where(ps...)
	return ucsu
}

// SetUserID sets the "user_id" field.
func (ucsu *UserCommentSubscriptionUpdate) SetUserID(i int) *UserCommentSubscriptionUpdate {
	ucsu.mutation.SetUserID(i)
	return ucsu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucsu *UserCommentSubscriptionUpdate) SetNillableUserID(i *int) *UserCommentSubscriptionUpdate {
	if i != nil {
		ucsu.SetUserID(*i)
	}
	return ucsu
}

// SetCommentID sets the "comment_id" field.
func (ucsu *UserCommentSubscriptionUpdate) SetCommentID(i int) *UserCommentSubscriptionUpdate {
	ucsu.mutation.SetCommentID(i)
	return ucsu
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (ucsu *UserCommentSubscriptionUpdate) SetNillableCommentID(i *int) *UserCommentSubscriptionUpdate {
	if i != nil {
		ucsu.SetCommentID(*i)
	}
	return ucsu
}

// SetIsNotified sets the "is_notified" field.
func (ucsu *UserCommentSubscriptionUpdate) SetIsNotified(b bool) *UserCommentSubscriptionUpdate {
	ucsu.mutation.SetIsNotified(b)
	return ucsu
}

// SetNillableIsNotified sets the "is_notified" field if the given value is not nil.
func (ucsu *UserCommentSubscriptionUpdate) SetNillableIsNotified(b *bool) *UserCommentSubscriptionUpdate {
	if b != nil {
		ucsu.SetIsNotified(*b)
	}
	return ucsu
}

// SetIsChecked sets the "is_checked" field.
func (ucsu *UserCommentSubscriptionUpdate) SetIsChecked(b bool) *UserCommentSubscriptionUpdate {
	ucsu.mutation.SetIsChecked(b)
	return ucsu
}

// SetNillableIsChecked sets the "is_checked" field if the given value is not nil.
func (ucsu *UserCommentSubscriptionUpdate) SetNillableIsChecked(b *bool) *UserCommentSubscriptionUpdate {
	if b != nil {
		ucsu.SetIsChecked(*b)
	}
	return ucsu
}

// SetSubscribedAt sets the "subscribed_at" field.
func (ucsu *UserCommentSubscriptionUpdate) SetSubscribedAt(t time.Time) *UserCommentSubscriptionUpdate {
	ucsu.mutation.SetSubscribedAt(t)
	return ucsu
}

// SetNillableSubscribedAt sets the "subscribed_at" field if the given value is not nil.
func (ucsu *UserCommentSubscriptionUpdate) SetNillableSubscribedAt(t *time.Time) *UserCommentSubscriptionUpdate {
	if t != nil {
		ucsu.SetSubscribedAt(*t)
	}
	return ucsu
}

// SetUser sets the "user" edge to the User entity.
func (ucsu *UserCommentSubscriptionUpdate) SetUser(u *User) *UserCommentSubscriptionUpdate {
	return ucsu.SetUserID(u.ID)
}

// SetComment sets the "comment" edge to the ThreadComment entity.
func (ucsu *UserCommentSubscriptionUpdate) SetComment(t *ThreadComment) *UserCommentSubscriptionUpdate {
	return ucsu.SetCommentID(t.ID)
}

// Mutation returns the UserCommentSubscriptionMutation object of the builder.
func (ucsu *UserCommentSubscriptionUpdate) Mutation() *UserCommentSubscriptionMutation {
	return ucsu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucsu *UserCommentSubscriptionUpdate) ClearUser() *UserCommentSubscriptionUpdate {
	ucsu.mutation.ClearUser()
	return ucsu
}

// ClearComment clears the "comment" edge to the ThreadComment entity.
func (ucsu *UserCommentSubscriptionUpdate) ClearComment() *UserCommentSubscriptionUpdate {
	ucsu.mutation.ClearComment()
	return ucsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucsu *UserCommentSubscriptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ucsu.sqlSave, ucsu.mutation, ucsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucsu *UserCommentSubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := ucsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucsu *UserCommentSubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := ucsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucsu *UserCommentSubscriptionUpdate) ExecX(ctx context.Context) {
	if err := ucsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucsu *UserCommentSubscriptionUpdate) check() error {
	if _, ok := ucsu.mutation.UserID(); ucsu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentSubscription.user"`)
	}
	if _, ok := ucsu.mutation.CommentID(); ucsu.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentSubscription.comment"`)
	}
	return nil
}

func (ucsu *UserCommentSubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ucsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercommentsubscription.Table, usercommentsubscription.Columns, sqlgraph.NewFieldSpec(usercommentsubscription.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(usercommentsubscription.FieldCommentID, field.TypeInt))
	if ps := ucsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucsu.mutation.IsNotified(); ok {
		_spec.SetField(usercommentsubscription.FieldIsNotified, field.TypeBool, value)
	}
	if value, ok := ucsu.mutation.IsChecked(); ok {
		_spec.SetField(usercommentsubscription.FieldIsChecked, field.TypeBool, value)
	}
	if value, ok := ucsu.mutation.SubscribedAt(); ok {
		_spec.SetField(usercommentsubscription.FieldSubscribedAt, field.TypeTime, value)
	}
	if ucsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.UserTable,
			Columns: []string{usercommentsubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.UserTable,
			Columns: []string{usercommentsubscription.UserColumn},
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
	if ucsu.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.CommentTable,
			Columns: []string{usercommentsubscription.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucsu.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.CommentTable,
			Columns: []string{usercommentsubscription.CommentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ucsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercommentsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucsu.mutation.done = true
	return n, nil
}

// UserCommentSubscriptionUpdateOne is the builder for updating a single UserCommentSubscription entity.
type UserCommentSubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCommentSubscriptionMutation
}

// SetUserID sets the "user_id" field.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetUserID(i int) *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.SetUserID(i)
	return ucsuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetNillableUserID(i *int) *UserCommentSubscriptionUpdateOne {
	if i != nil {
		ucsuo.SetUserID(*i)
	}
	return ucsuo
}

// SetCommentID sets the "comment_id" field.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetCommentID(i int) *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.SetCommentID(i)
	return ucsuo
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetNillableCommentID(i *int) *UserCommentSubscriptionUpdateOne {
	if i != nil {
		ucsuo.SetCommentID(*i)
	}
	return ucsuo
}

// SetIsNotified sets the "is_notified" field.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetIsNotified(b bool) *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.SetIsNotified(b)
	return ucsuo
}

// SetNillableIsNotified sets the "is_notified" field if the given value is not nil.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetNillableIsNotified(b *bool) *UserCommentSubscriptionUpdateOne {
	if b != nil {
		ucsuo.SetIsNotified(*b)
	}
	return ucsuo
}

// SetIsChecked sets the "is_checked" field.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetIsChecked(b bool) *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.SetIsChecked(b)
	return ucsuo
}

// SetNillableIsChecked sets the "is_checked" field if the given value is not nil.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetNillableIsChecked(b *bool) *UserCommentSubscriptionUpdateOne {
	if b != nil {
		ucsuo.SetIsChecked(*b)
	}
	return ucsuo
}

// SetSubscribedAt sets the "subscribed_at" field.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetSubscribedAt(t time.Time) *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.SetSubscribedAt(t)
	return ucsuo
}

// SetNillableSubscribedAt sets the "subscribed_at" field if the given value is not nil.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetNillableSubscribedAt(t *time.Time) *UserCommentSubscriptionUpdateOne {
	if t != nil {
		ucsuo.SetSubscribedAt(*t)
	}
	return ucsuo
}

// SetUser sets the "user" edge to the User entity.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetUser(u *User) *UserCommentSubscriptionUpdateOne {
	return ucsuo.SetUserID(u.ID)
}

// SetComment sets the "comment" edge to the ThreadComment entity.
func (ucsuo *UserCommentSubscriptionUpdateOne) SetComment(t *ThreadComment) *UserCommentSubscriptionUpdateOne {
	return ucsuo.SetCommentID(t.ID)
}

// Mutation returns the UserCommentSubscriptionMutation object of the builder.
func (ucsuo *UserCommentSubscriptionUpdateOne) Mutation() *UserCommentSubscriptionMutation {
	return ucsuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ucsuo *UserCommentSubscriptionUpdateOne) ClearUser() *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.ClearUser()
	return ucsuo
}

// ClearComment clears the "comment" edge to the ThreadComment entity.
func (ucsuo *UserCommentSubscriptionUpdateOne) ClearComment() *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.ClearComment()
	return ucsuo
}

// Where appends a list predicates to the UserCommentSubscriptionUpdate builder.
func (ucsuo *UserCommentSubscriptionUpdateOne) Where(ps ...predicate.UserCommentSubscription) *UserCommentSubscriptionUpdateOne {
	ucsuo.mutation.Where(ps...)
	return ucsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucsuo *UserCommentSubscriptionUpdateOne) Select(field string, fields ...string) *UserCommentSubscriptionUpdateOne {
	ucsuo.fields = append([]string{field}, fields...)
	return ucsuo
}

// Save executes the query and returns the updated UserCommentSubscription entity.
func (ucsuo *UserCommentSubscriptionUpdateOne) Save(ctx context.Context) (*UserCommentSubscription, error) {
	return withHooks(ctx, ucsuo.sqlSave, ucsuo.mutation, ucsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucsuo *UserCommentSubscriptionUpdateOne) SaveX(ctx context.Context) *UserCommentSubscription {
	node, err := ucsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucsuo *UserCommentSubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := ucsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucsuo *UserCommentSubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := ucsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucsuo *UserCommentSubscriptionUpdateOne) check() error {
	if _, ok := ucsuo.mutation.UserID(); ucsuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentSubscription.user"`)
	}
	if _, ok := ucsuo.mutation.CommentID(); ucsuo.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCommentSubscription.comment"`)
	}
	return nil
}

func (ucsuo *UserCommentSubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *UserCommentSubscription, err error) {
	if err := ucsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercommentsubscription.Table, usercommentsubscription.Columns, sqlgraph.NewFieldSpec(usercommentsubscription.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(usercommentsubscription.FieldCommentID, field.TypeInt))
	if id, ok := ucsuo.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "UserCommentSubscription.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ucsuo.mutation.CommentID(); !ok {
		return nil, &ValidationError{Name: "comment_id", err: errors.New(`ent: missing "UserCommentSubscription.comment_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ucsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !usercommentsubscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ucsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucsuo.mutation.IsNotified(); ok {
		_spec.SetField(usercommentsubscription.FieldIsNotified, field.TypeBool, value)
	}
	if value, ok := ucsuo.mutation.IsChecked(); ok {
		_spec.SetField(usercommentsubscription.FieldIsChecked, field.TypeBool, value)
	}
	if value, ok := ucsuo.mutation.SubscribedAt(); ok {
		_spec.SetField(usercommentsubscription.FieldSubscribedAt, field.TypeTime, value)
	}
	if ucsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.UserTable,
			Columns: []string{usercommentsubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.UserTable,
			Columns: []string{usercommentsubscription.UserColumn},
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
	if ucsuo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.CommentTable,
			Columns: []string{usercommentsubscription.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(threadcomment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucsuo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentsubscription.CommentTable,
			Columns: []string{usercommentsubscription.CommentColumn},
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
	_node = &UserCommentSubscription{config: ucsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercommentsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucsuo.mutation.done = true
	return _node, nil
}
