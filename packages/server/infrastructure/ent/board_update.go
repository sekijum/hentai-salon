// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BoardUpdate is the builder for updating Board entities.
type BoardUpdate struct {
	config
	hooks    []Hook
	mutation *BoardMutation
}

// Where appends a list predicates to the BoardUpdate builder.
func (bu *BoardUpdate) Where(ps ...predicate.Board) *BoardUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUserId sets the "userId" field.
func (bu *BoardUpdate) SetUserId(i int) *BoardUpdate {
	bu.mutation.SetUserId(i)
	return bu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableUserId(i *int) *BoardUpdate {
	if i != nil {
		bu.SetUserId(*i)
	}
	return bu
}

// SetTitle sets the "title" field.
func (bu *BoardUpdate) SetTitle(s string) *BoardUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableTitle(s *string) *BoardUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetDescription sets the "description" field.
func (bu *BoardUpdate) SetDescription(s string) *BoardUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableDescription(s *string) *BoardUpdate {
	if s != nil {
		bu.SetDescription(*s)
	}
	return bu
}

// ClearDescription clears the value of the "description" field.
func (bu *BoardUpdate) ClearDescription() *BoardUpdate {
	bu.mutation.ClearDescription()
	return bu
}

// SetThumbnailUrl sets the "thumbnailUrl" field.
func (bu *BoardUpdate) SetThumbnailUrl(s string) *BoardUpdate {
	bu.mutation.SetThumbnailUrl(s)
	return bu
}

// SetNillableThumbnailUrl sets the "thumbnailUrl" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableThumbnailUrl(s *string) *BoardUpdate {
	if s != nil {
		bu.SetThumbnailUrl(*s)
	}
	return bu
}

// ClearThumbnailUrl clears the value of the "thumbnailUrl" field.
func (bu *BoardUpdate) ClearThumbnailUrl() *BoardUpdate {
	bu.mutation.ClearThumbnailUrl()
	return bu
}

// SetStatus sets the "status" field.
func (bu *BoardUpdate) SetStatus(i int) *BoardUpdate {
	bu.mutation.ResetStatus()
	bu.mutation.SetStatus(i)
	return bu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableStatus(i *int) *BoardUpdate {
	if i != nil {
		bu.SetStatus(*i)
	}
	return bu
}

// AddStatus adds i to the "status" field.
func (bu *BoardUpdate) AddStatus(i int) *BoardUpdate {
	bu.mutation.AddStatus(i)
	return bu
}

// SetCreatedAt sets the "createdAt" field.
func (bu *BoardUpdate) SetCreatedAt(t time.Time) *BoardUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableCreatedAt(t *time.Time) *BoardUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updatedAt" field.
func (bu *BoardUpdate) SetUpdatedAt(t time.Time) *BoardUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (bu *BoardUpdate) AddLikedUserIDs(ids ...int) *BoardUpdate {
	bu.mutation.AddLikedUserIDs(ids...)
	return bu
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (bu *BoardUpdate) AddLikedUsers(u ...*User) *BoardUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.AddLikedUserIDs(ids...)
}

// AddSubscribedUserIDs adds the "subscribed_users" edge to the User entity by IDs.
func (bu *BoardUpdate) AddSubscribedUserIDs(ids ...int) *BoardUpdate {
	bu.mutation.AddSubscribedUserIDs(ids...)
	return bu
}

// AddSubscribedUsers adds the "subscribed_users" edges to the User entity.
func (bu *BoardUpdate) AddSubscribedUsers(u ...*User) *BoardUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.AddSubscribedUserIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (bu *BoardUpdate) SetOwnerID(id int) *BoardUpdate {
	bu.mutation.SetOwnerID(id)
	return bu
}

// SetOwner sets the "owner" edge to the User entity.
func (bu *BoardUpdate) SetOwner(u *User) *BoardUpdate {
	return bu.SetOwnerID(u.ID)
}

// AddThreadIDs adds the "threads" edge to the Thread entity by IDs.
func (bu *BoardUpdate) AddThreadIDs(ids ...int) *BoardUpdate {
	bu.mutation.AddThreadIDs(ids...)
	return bu
}

// AddThreads adds the "threads" edges to the Thread entity.
func (bu *BoardUpdate) AddThreads(t ...*Thread) *BoardUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddThreadIDs(ids...)
}

// Mutation returns the BoardMutation object of the builder.
func (bu *BoardUpdate) Mutation() *BoardMutation {
	return bu.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (bu *BoardUpdate) ClearLikedUsers() *BoardUpdate {
	bu.mutation.ClearLikedUsers()
	return bu
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (bu *BoardUpdate) RemoveLikedUserIDs(ids ...int) *BoardUpdate {
	bu.mutation.RemoveLikedUserIDs(ids...)
	return bu
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (bu *BoardUpdate) RemoveLikedUsers(u ...*User) *BoardUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.RemoveLikedUserIDs(ids...)
}

// ClearSubscribedUsers clears all "subscribed_users" edges to the User entity.
func (bu *BoardUpdate) ClearSubscribedUsers() *BoardUpdate {
	bu.mutation.ClearSubscribedUsers()
	return bu
}

// RemoveSubscribedUserIDs removes the "subscribed_users" edge to User entities by IDs.
func (bu *BoardUpdate) RemoveSubscribedUserIDs(ids ...int) *BoardUpdate {
	bu.mutation.RemoveSubscribedUserIDs(ids...)
	return bu
}

// RemoveSubscribedUsers removes "subscribed_users" edges to User entities.
func (bu *BoardUpdate) RemoveSubscribedUsers(u ...*User) *BoardUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.RemoveSubscribedUserIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (bu *BoardUpdate) ClearOwner() *BoardUpdate {
	bu.mutation.ClearOwner()
	return bu
}

// ClearThreads clears all "threads" edges to the Thread entity.
func (bu *BoardUpdate) ClearThreads() *BoardUpdate {
	bu.mutation.ClearThreads()
	return bu
}

// RemoveThreadIDs removes the "threads" edge to Thread entities by IDs.
func (bu *BoardUpdate) RemoveThreadIDs(ids ...int) *BoardUpdate {
	bu.mutation.RemoveThreadIDs(ids...)
	return bu
}

// RemoveThreads removes "threads" edges to Thread entities.
func (bu *BoardUpdate) RemoveThreads(t ...*Thread) *BoardUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveThreadIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BoardUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BoardUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BoardUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BoardUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BoardUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := board.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BoardUpdate) check() error {
	if v, ok := bu.mutation.Title(); ok {
		if err := board.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Board.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Description(); ok {
		if err := board.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Board.description": %w`, err)}
		}
	}
	if _, ok := bu.mutation.OwnerID(); bu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Board.owner"`)
	}
	return nil
}

func (bu *BoardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.SetField(board.FieldDescription, field.TypeString, value)
	}
	if bu.mutation.DescriptionCleared() {
		_spec.ClearField(board.FieldDescription, field.TypeString)
	}
	if value, ok := bu.mutation.ThumbnailUrl(); ok {
		_spec.SetField(board.FieldThumbnailUrl, field.TypeString, value)
	}
	if bu.mutation.ThumbnailUrlCleared() {
		_spec.ClearField(board.FieldThumbnailUrl, field.TypeString)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(board.FieldStatus, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedStatus(); ok {
		_spec.AddField(board.FieldStatus, field.TypeInt, value)
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(board.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(board.FieldUpdatedAt, field.TypeTime, value)
	}
	if bu.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.LikedUsersTable,
			Columns: board.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &UserBoardSubscriptionCreate{config: bu.config, mutation: newUserBoardSubscriptionMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !bu.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.LikedUsersTable,
			Columns: board.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardSubscriptionCreate{config: bu.config, mutation: newUserBoardSubscriptionMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.LikedUsersTable,
			Columns: board.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardSubscriptionCreate{config: bu.config, mutation: newUserBoardSubscriptionMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.SubscribedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.SubscribedUsersTable,
			Columns: board.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &UserBoardLikeCreate{config: bu.config, mutation: newUserBoardLikeMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedSubscribedUsersIDs(); len(nodes) > 0 && !bu.mutation.SubscribedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.SubscribedUsersTable,
			Columns: board.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardLikeCreate{config: bu.config, mutation: newUserBoardLikeMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.SubscribedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.SubscribedUsersTable,
			Columns: board.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardLikeCreate{config: bu.config, mutation: newUserBoardLikeMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.OwnerTable,
			Columns: []string{board.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.OwnerTable,
			Columns: []string{board.OwnerColumn},
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
	if bu.mutation.ThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.ThreadsTable,
			Columns: []string{board.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedThreadsIDs(); len(nodes) > 0 && !bu.mutation.ThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.ThreadsTable,
			Columns: []string{board.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.ThreadsTable,
			Columns: []string{board.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BoardUpdateOne is the builder for updating a single Board entity.
type BoardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BoardMutation
}

// SetUserId sets the "userId" field.
func (buo *BoardUpdateOne) SetUserId(i int) *BoardUpdateOne {
	buo.mutation.SetUserId(i)
	return buo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableUserId(i *int) *BoardUpdateOne {
	if i != nil {
		buo.SetUserId(*i)
	}
	return buo
}

// SetTitle sets the "title" field.
func (buo *BoardUpdateOne) SetTitle(s string) *BoardUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableTitle(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetDescription sets the "description" field.
func (buo *BoardUpdateOne) SetDescription(s string) *BoardUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableDescription(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetDescription(*s)
	}
	return buo
}

// ClearDescription clears the value of the "description" field.
func (buo *BoardUpdateOne) ClearDescription() *BoardUpdateOne {
	buo.mutation.ClearDescription()
	return buo
}

// SetThumbnailUrl sets the "thumbnailUrl" field.
func (buo *BoardUpdateOne) SetThumbnailUrl(s string) *BoardUpdateOne {
	buo.mutation.SetThumbnailUrl(s)
	return buo
}

// SetNillableThumbnailUrl sets the "thumbnailUrl" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableThumbnailUrl(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetThumbnailUrl(*s)
	}
	return buo
}

// ClearThumbnailUrl clears the value of the "thumbnailUrl" field.
func (buo *BoardUpdateOne) ClearThumbnailUrl() *BoardUpdateOne {
	buo.mutation.ClearThumbnailUrl()
	return buo
}

// SetStatus sets the "status" field.
func (buo *BoardUpdateOne) SetStatus(i int) *BoardUpdateOne {
	buo.mutation.ResetStatus()
	buo.mutation.SetStatus(i)
	return buo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableStatus(i *int) *BoardUpdateOne {
	if i != nil {
		buo.SetStatus(*i)
	}
	return buo
}

// AddStatus adds i to the "status" field.
func (buo *BoardUpdateOne) AddStatus(i int) *BoardUpdateOne {
	buo.mutation.AddStatus(i)
	return buo
}

// SetCreatedAt sets the "createdAt" field.
func (buo *BoardUpdateOne) SetCreatedAt(t time.Time) *BoardUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableCreatedAt(t *time.Time) *BoardUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updatedAt" field.
func (buo *BoardUpdateOne) SetUpdatedAt(t time.Time) *BoardUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (buo *BoardUpdateOne) AddLikedUserIDs(ids ...int) *BoardUpdateOne {
	buo.mutation.AddLikedUserIDs(ids...)
	return buo
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (buo *BoardUpdateOne) AddLikedUsers(u ...*User) *BoardUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.AddLikedUserIDs(ids...)
}

// AddSubscribedUserIDs adds the "subscribed_users" edge to the User entity by IDs.
func (buo *BoardUpdateOne) AddSubscribedUserIDs(ids ...int) *BoardUpdateOne {
	buo.mutation.AddSubscribedUserIDs(ids...)
	return buo
}

// AddSubscribedUsers adds the "subscribed_users" edges to the User entity.
func (buo *BoardUpdateOne) AddSubscribedUsers(u ...*User) *BoardUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.AddSubscribedUserIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (buo *BoardUpdateOne) SetOwnerID(id int) *BoardUpdateOne {
	buo.mutation.SetOwnerID(id)
	return buo
}

// SetOwner sets the "owner" edge to the User entity.
func (buo *BoardUpdateOne) SetOwner(u *User) *BoardUpdateOne {
	return buo.SetOwnerID(u.ID)
}

// AddThreadIDs adds the "threads" edge to the Thread entity by IDs.
func (buo *BoardUpdateOne) AddThreadIDs(ids ...int) *BoardUpdateOne {
	buo.mutation.AddThreadIDs(ids...)
	return buo
}

// AddThreads adds the "threads" edges to the Thread entity.
func (buo *BoardUpdateOne) AddThreads(t ...*Thread) *BoardUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddThreadIDs(ids...)
}

// Mutation returns the BoardMutation object of the builder.
func (buo *BoardUpdateOne) Mutation() *BoardMutation {
	return buo.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (buo *BoardUpdateOne) ClearLikedUsers() *BoardUpdateOne {
	buo.mutation.ClearLikedUsers()
	return buo
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (buo *BoardUpdateOne) RemoveLikedUserIDs(ids ...int) *BoardUpdateOne {
	buo.mutation.RemoveLikedUserIDs(ids...)
	return buo
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (buo *BoardUpdateOne) RemoveLikedUsers(u ...*User) *BoardUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.RemoveLikedUserIDs(ids...)
}

// ClearSubscribedUsers clears all "subscribed_users" edges to the User entity.
func (buo *BoardUpdateOne) ClearSubscribedUsers() *BoardUpdateOne {
	buo.mutation.ClearSubscribedUsers()
	return buo
}

// RemoveSubscribedUserIDs removes the "subscribed_users" edge to User entities by IDs.
func (buo *BoardUpdateOne) RemoveSubscribedUserIDs(ids ...int) *BoardUpdateOne {
	buo.mutation.RemoveSubscribedUserIDs(ids...)
	return buo
}

// RemoveSubscribedUsers removes "subscribed_users" edges to User entities.
func (buo *BoardUpdateOne) RemoveSubscribedUsers(u ...*User) *BoardUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.RemoveSubscribedUserIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (buo *BoardUpdateOne) ClearOwner() *BoardUpdateOne {
	buo.mutation.ClearOwner()
	return buo
}

// ClearThreads clears all "threads" edges to the Thread entity.
func (buo *BoardUpdateOne) ClearThreads() *BoardUpdateOne {
	buo.mutation.ClearThreads()
	return buo
}

// RemoveThreadIDs removes the "threads" edge to Thread entities by IDs.
func (buo *BoardUpdateOne) RemoveThreadIDs(ids ...int) *BoardUpdateOne {
	buo.mutation.RemoveThreadIDs(ids...)
	return buo
}

// RemoveThreads removes "threads" edges to Thread entities.
func (buo *BoardUpdateOne) RemoveThreads(t ...*Thread) *BoardUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveThreadIDs(ids...)
}

// Where appends a list predicates to the BoardUpdate builder.
func (buo *BoardUpdateOne) Where(ps ...predicate.Board) *BoardUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BoardUpdateOne) Select(field string, fields ...string) *BoardUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Board entity.
func (buo *BoardUpdateOne) Save(ctx context.Context) (*Board, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BoardUpdateOne) SaveX(ctx context.Context) *Board {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BoardUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BoardUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BoardUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := board.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BoardUpdateOne) check() error {
	if v, ok := buo.mutation.Title(); ok {
		if err := board.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Board.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Description(); ok {
		if err := board.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Board.description": %w`, err)}
		}
	}
	if _, ok := buo.mutation.OwnerID(); buo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Board.owner"`)
	}
	return nil
}

func (buo *BoardUpdateOne) sqlSave(ctx context.Context) (_node *Board, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Board.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, board.FieldID)
		for _, f := range fields {
			if !board.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != board.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.SetField(board.FieldDescription, field.TypeString, value)
	}
	if buo.mutation.DescriptionCleared() {
		_spec.ClearField(board.FieldDescription, field.TypeString)
	}
	if value, ok := buo.mutation.ThumbnailUrl(); ok {
		_spec.SetField(board.FieldThumbnailUrl, field.TypeString, value)
	}
	if buo.mutation.ThumbnailUrlCleared() {
		_spec.ClearField(board.FieldThumbnailUrl, field.TypeString)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(board.FieldStatus, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedStatus(); ok {
		_spec.AddField(board.FieldStatus, field.TypeInt, value)
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(board.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(board.FieldUpdatedAt, field.TypeTime, value)
	}
	if buo.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.LikedUsersTable,
			Columns: board.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &UserBoardSubscriptionCreate{config: buo.config, mutation: newUserBoardSubscriptionMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !buo.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.LikedUsersTable,
			Columns: board.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardSubscriptionCreate{config: buo.config, mutation: newUserBoardSubscriptionMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.LikedUsersTable,
			Columns: board.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardSubscriptionCreate{config: buo.config, mutation: newUserBoardSubscriptionMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.SubscribedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.SubscribedUsersTable,
			Columns: board.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &UserBoardLikeCreate{config: buo.config, mutation: newUserBoardLikeMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedSubscribedUsersIDs(); len(nodes) > 0 && !buo.mutation.SubscribedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.SubscribedUsersTable,
			Columns: board.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardLikeCreate{config: buo.config, mutation: newUserBoardLikeMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.SubscribedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   board.SubscribedUsersTable,
			Columns: board.SubscribedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserBoardLikeCreate{config: buo.config, mutation: newUserBoardLikeMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.OwnerTable,
			Columns: []string{board.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.OwnerTable,
			Columns: []string{board.OwnerColumn},
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
	if buo.mutation.ThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.ThreadsTable,
			Columns: []string{board.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedThreadsIDs(); len(nodes) > 0 && !buo.mutation.ThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.ThreadsTable,
			Columns: []string{board.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.ThreadsTable,
			Columns: []string{board.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Board{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
