// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userboardsubscription"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBoardSubscriptionUpdate is the builder for updating UserBoardSubscription entities.
type UserBoardSubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *UserBoardSubscriptionMutation
}

// Where appends a list predicates to the UserBoardSubscriptionUpdate builder.
func (ubsu *UserBoardSubscriptionUpdate) Where(ps ...predicate.UserBoardSubscription) *UserBoardSubscriptionUpdate {
	ubsu.mutation.Where(ps...)
	return ubsu
}

// SetUserId sets the "userId" field.
func (ubsu *UserBoardSubscriptionUpdate) SetUserId(i int) *UserBoardSubscriptionUpdate {
	ubsu.mutation.SetUserId(i)
	return ubsu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ubsu *UserBoardSubscriptionUpdate) SetNillableUserId(i *int) *UserBoardSubscriptionUpdate {
	if i != nil {
		ubsu.SetUserId(*i)
	}
	return ubsu
}

// SetBoardId sets the "boardId" field.
func (ubsu *UserBoardSubscriptionUpdate) SetBoardId(i int) *UserBoardSubscriptionUpdate {
	ubsu.mutation.SetBoardId(i)
	return ubsu
}

// SetNillableBoardId sets the "boardId" field if the given value is not nil.
func (ubsu *UserBoardSubscriptionUpdate) SetNillableBoardId(i *int) *UserBoardSubscriptionUpdate {
	if i != nil {
		ubsu.SetBoardId(*i)
	}
	return ubsu
}

// SetIsNotified sets the "isNotified" field.
func (ubsu *UserBoardSubscriptionUpdate) SetIsNotified(b bool) *UserBoardSubscriptionUpdate {
	ubsu.mutation.SetIsNotified(b)
	return ubsu
}

// SetNillableIsNotified sets the "isNotified" field if the given value is not nil.
func (ubsu *UserBoardSubscriptionUpdate) SetNillableIsNotified(b *bool) *UserBoardSubscriptionUpdate {
	if b != nil {
		ubsu.SetIsNotified(*b)
	}
	return ubsu
}

// SetIsChecked sets the "isChecked" field.
func (ubsu *UserBoardSubscriptionUpdate) SetIsChecked(b bool) *UserBoardSubscriptionUpdate {
	ubsu.mutation.SetIsChecked(b)
	return ubsu
}

// SetNillableIsChecked sets the "isChecked" field if the given value is not nil.
func (ubsu *UserBoardSubscriptionUpdate) SetNillableIsChecked(b *bool) *UserBoardSubscriptionUpdate {
	if b != nil {
		ubsu.SetIsChecked(*b)
	}
	return ubsu
}

// SetSubscribedAt sets the "subscribedAt" field.
func (ubsu *UserBoardSubscriptionUpdate) SetSubscribedAt(t time.Time) *UserBoardSubscriptionUpdate {
	ubsu.mutation.SetSubscribedAt(t)
	return ubsu
}

// SetNillableSubscribedAt sets the "subscribedAt" field if the given value is not nil.
func (ubsu *UserBoardSubscriptionUpdate) SetNillableSubscribedAt(t *time.Time) *UserBoardSubscriptionUpdate {
	if t != nil {
		ubsu.SetSubscribedAt(*t)
	}
	return ubsu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ubsu *UserBoardSubscriptionUpdate) SetUserID(id int) *UserBoardSubscriptionUpdate {
	ubsu.mutation.SetUserID(id)
	return ubsu
}

// SetUser sets the "user" edge to the User entity.
func (ubsu *UserBoardSubscriptionUpdate) SetUser(u *User) *UserBoardSubscriptionUpdate {
	return ubsu.SetUserID(u.ID)
}

// SetBoardID sets the "board" edge to the Board entity by ID.
func (ubsu *UserBoardSubscriptionUpdate) SetBoardID(id int) *UserBoardSubscriptionUpdate {
	ubsu.mutation.SetBoardID(id)
	return ubsu
}

// SetBoard sets the "board" edge to the Board entity.
func (ubsu *UserBoardSubscriptionUpdate) SetBoard(b *Board) *UserBoardSubscriptionUpdate {
	return ubsu.SetBoardID(b.ID)
}

// Mutation returns the UserBoardSubscriptionMutation object of the builder.
func (ubsu *UserBoardSubscriptionUpdate) Mutation() *UserBoardSubscriptionMutation {
	return ubsu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ubsu *UserBoardSubscriptionUpdate) ClearUser() *UserBoardSubscriptionUpdate {
	ubsu.mutation.ClearUser()
	return ubsu
}

// ClearBoard clears the "board" edge to the Board entity.
func (ubsu *UserBoardSubscriptionUpdate) ClearBoard() *UserBoardSubscriptionUpdate {
	ubsu.mutation.ClearBoard()
	return ubsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ubsu *UserBoardSubscriptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ubsu.sqlSave, ubsu.mutation, ubsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ubsu *UserBoardSubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := ubsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ubsu *UserBoardSubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := ubsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubsu *UserBoardSubscriptionUpdate) ExecX(ctx context.Context) {
	if err := ubsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ubsu *UserBoardSubscriptionUpdate) check() error {
	if _, ok := ubsu.mutation.UserID(); ubsu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardSubscription.user"`)
	}
	if _, ok := ubsu.mutation.BoardID(); ubsu.mutation.BoardCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardSubscription.board"`)
	}
	return nil
}

func (ubsu *UserBoardSubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ubsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userboardsubscription.Table, userboardsubscription.Columns, sqlgraph.NewFieldSpec(userboardsubscription.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(userboardsubscription.FieldBoardId, field.TypeInt))
	if ps := ubsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ubsu.mutation.IsNotified(); ok {
		_spec.SetField(userboardsubscription.FieldIsNotified, field.TypeBool, value)
	}
	if value, ok := ubsu.mutation.IsChecked(); ok {
		_spec.SetField(userboardsubscription.FieldIsChecked, field.TypeBool, value)
	}
	if value, ok := ubsu.mutation.SubscribedAt(); ok {
		_spec.SetField(userboardsubscription.FieldSubscribedAt, field.TypeTime, value)
	}
	if ubsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.UserTable,
			Columns: []string{userboardsubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.UserTable,
			Columns: []string{userboardsubscription.UserColumn},
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
	if ubsu.mutation.BoardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.BoardTable,
			Columns: []string{userboardsubscription.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubsu.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.BoardTable,
			Columns: []string{userboardsubscription.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ubsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userboardsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ubsu.mutation.done = true
	return n, nil
}

// UserBoardSubscriptionUpdateOne is the builder for updating a single UserBoardSubscription entity.
type UserBoardSubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserBoardSubscriptionMutation
}

// SetUserId sets the "userId" field.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetUserId(i int) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.SetUserId(i)
	return ubsuo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetNillableUserId(i *int) *UserBoardSubscriptionUpdateOne {
	if i != nil {
		ubsuo.SetUserId(*i)
	}
	return ubsuo
}

// SetBoardId sets the "boardId" field.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetBoardId(i int) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.SetBoardId(i)
	return ubsuo
}

// SetNillableBoardId sets the "boardId" field if the given value is not nil.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetNillableBoardId(i *int) *UserBoardSubscriptionUpdateOne {
	if i != nil {
		ubsuo.SetBoardId(*i)
	}
	return ubsuo
}

// SetIsNotified sets the "isNotified" field.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetIsNotified(b bool) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.SetIsNotified(b)
	return ubsuo
}

// SetNillableIsNotified sets the "isNotified" field if the given value is not nil.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetNillableIsNotified(b *bool) *UserBoardSubscriptionUpdateOne {
	if b != nil {
		ubsuo.SetIsNotified(*b)
	}
	return ubsuo
}

// SetIsChecked sets the "isChecked" field.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetIsChecked(b bool) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.SetIsChecked(b)
	return ubsuo
}

// SetNillableIsChecked sets the "isChecked" field if the given value is not nil.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetNillableIsChecked(b *bool) *UserBoardSubscriptionUpdateOne {
	if b != nil {
		ubsuo.SetIsChecked(*b)
	}
	return ubsuo
}

// SetSubscribedAt sets the "subscribedAt" field.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetSubscribedAt(t time.Time) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.SetSubscribedAt(t)
	return ubsuo
}

// SetNillableSubscribedAt sets the "subscribedAt" field if the given value is not nil.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetNillableSubscribedAt(t *time.Time) *UserBoardSubscriptionUpdateOne {
	if t != nil {
		ubsuo.SetSubscribedAt(*t)
	}
	return ubsuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetUserID(id int) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.SetUserID(id)
	return ubsuo
}

// SetUser sets the "user" edge to the User entity.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetUser(u *User) *UserBoardSubscriptionUpdateOne {
	return ubsuo.SetUserID(u.ID)
}

// SetBoardID sets the "board" edge to the Board entity by ID.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetBoardID(id int) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.SetBoardID(id)
	return ubsuo
}

// SetBoard sets the "board" edge to the Board entity.
func (ubsuo *UserBoardSubscriptionUpdateOne) SetBoard(b *Board) *UserBoardSubscriptionUpdateOne {
	return ubsuo.SetBoardID(b.ID)
}

// Mutation returns the UserBoardSubscriptionMutation object of the builder.
func (ubsuo *UserBoardSubscriptionUpdateOne) Mutation() *UserBoardSubscriptionMutation {
	return ubsuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ubsuo *UserBoardSubscriptionUpdateOne) ClearUser() *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.ClearUser()
	return ubsuo
}

// ClearBoard clears the "board" edge to the Board entity.
func (ubsuo *UserBoardSubscriptionUpdateOne) ClearBoard() *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.ClearBoard()
	return ubsuo
}

// Where appends a list predicates to the UserBoardSubscriptionUpdate builder.
func (ubsuo *UserBoardSubscriptionUpdateOne) Where(ps ...predicate.UserBoardSubscription) *UserBoardSubscriptionUpdateOne {
	ubsuo.mutation.Where(ps...)
	return ubsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ubsuo *UserBoardSubscriptionUpdateOne) Select(field string, fields ...string) *UserBoardSubscriptionUpdateOne {
	ubsuo.fields = append([]string{field}, fields...)
	return ubsuo
}

// Save executes the query and returns the updated UserBoardSubscription entity.
func (ubsuo *UserBoardSubscriptionUpdateOne) Save(ctx context.Context) (*UserBoardSubscription, error) {
	return withHooks(ctx, ubsuo.sqlSave, ubsuo.mutation, ubsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ubsuo *UserBoardSubscriptionUpdateOne) SaveX(ctx context.Context) *UserBoardSubscription {
	node, err := ubsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ubsuo *UserBoardSubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := ubsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubsuo *UserBoardSubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := ubsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ubsuo *UserBoardSubscriptionUpdateOne) check() error {
	if _, ok := ubsuo.mutation.UserID(); ubsuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardSubscription.user"`)
	}
	if _, ok := ubsuo.mutation.BoardID(); ubsuo.mutation.BoardCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserBoardSubscription.board"`)
	}
	return nil
}

func (ubsuo *UserBoardSubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *UserBoardSubscription, err error) {
	if err := ubsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userboardsubscription.Table, userboardsubscription.Columns, sqlgraph.NewFieldSpec(userboardsubscription.FieldUserId, field.TypeInt), sqlgraph.NewFieldSpec(userboardsubscription.FieldBoardId, field.TypeInt))
	if id, ok := ubsuo.mutation.UserId(); !ok {
		return nil, &ValidationError{Name: "userId", err: errors.New(`ent: missing "UserBoardSubscription.userId" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ubsuo.mutation.BoardId(); !ok {
		return nil, &ValidationError{Name: "boardId", err: errors.New(`ent: missing "UserBoardSubscription.boardId" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ubsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !userboardsubscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ubsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ubsuo.mutation.IsNotified(); ok {
		_spec.SetField(userboardsubscription.FieldIsNotified, field.TypeBool, value)
	}
	if value, ok := ubsuo.mutation.IsChecked(); ok {
		_spec.SetField(userboardsubscription.FieldIsChecked, field.TypeBool, value)
	}
	if value, ok := ubsuo.mutation.SubscribedAt(); ok {
		_spec.SetField(userboardsubscription.FieldSubscribedAt, field.TypeTime, value)
	}
	if ubsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.UserTable,
			Columns: []string{userboardsubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.UserTable,
			Columns: []string{userboardsubscription.UserColumn},
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
	if ubsuo.mutation.BoardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.BoardTable,
			Columns: []string{userboardsubscription.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ubsuo.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userboardsubscription.BoardTable,
			Columns: []string{userboardsubscription.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserBoardSubscription{config: ubsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ubsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userboardsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ubsuo.mutation.done = true
	return _node, nil
}