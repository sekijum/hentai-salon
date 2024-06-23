// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/infrastructure/ent/comment"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usercommentlike"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCommentLikeCreate is the builder for creating a UserCommentLike entity.
type UserCommentLikeCreate struct {
	config
	mutation *UserCommentLikeMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (uclc *UserCommentLikeCreate) SetUserId(i int) *UserCommentLikeCreate {
	uclc.mutation.SetUserId(i)
	return uclc
}

// SetCommentId sets the "commentId" field.
func (uclc *UserCommentLikeCreate) SetCommentId(i int) *UserCommentLikeCreate {
	uclc.mutation.SetCommentId(i)
	return uclc
}

// SetLikedAt sets the "likedAt" field.
func (uclc *UserCommentLikeCreate) SetLikedAt(t time.Time) *UserCommentLikeCreate {
	uclc.mutation.SetLikedAt(t)
	return uclc
}

// SetNillableLikedAt sets the "likedAt" field if the given value is not nil.
func (uclc *UserCommentLikeCreate) SetNillableLikedAt(t *time.Time) *UserCommentLikeCreate {
	if t != nil {
		uclc.SetLikedAt(*t)
	}
	return uclc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uclc *UserCommentLikeCreate) SetUserID(id int) *UserCommentLikeCreate {
	uclc.mutation.SetUserID(id)
	return uclc
}

// SetUser sets the "user" edge to the User entity.
func (uclc *UserCommentLikeCreate) SetUser(u *User) *UserCommentLikeCreate {
	return uclc.SetUserID(u.ID)
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (uclc *UserCommentLikeCreate) SetCommentID(id int) *UserCommentLikeCreate {
	uclc.mutation.SetCommentID(id)
	return uclc
}

// SetComment sets the "comment" edge to the Comment entity.
func (uclc *UserCommentLikeCreate) SetComment(c *Comment) *UserCommentLikeCreate {
	return uclc.SetCommentID(c.ID)
}

// Mutation returns the UserCommentLikeMutation object of the builder.
func (uclc *UserCommentLikeCreate) Mutation() *UserCommentLikeMutation {
	return uclc.mutation
}

// Save creates the UserCommentLike in the database.
func (uclc *UserCommentLikeCreate) Save(ctx context.Context) (*UserCommentLike, error) {
	uclc.defaults()
	return withHooks(ctx, uclc.sqlSave, uclc.mutation, uclc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uclc *UserCommentLikeCreate) SaveX(ctx context.Context) *UserCommentLike {
	v, err := uclc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uclc *UserCommentLikeCreate) Exec(ctx context.Context) error {
	_, err := uclc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uclc *UserCommentLikeCreate) ExecX(ctx context.Context) {
	if err := uclc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uclc *UserCommentLikeCreate) defaults() {
	if _, ok := uclc.mutation.LikedAt(); !ok {
		v := usercommentlike.DefaultLikedAt()
		uclc.mutation.SetLikedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uclc *UserCommentLikeCreate) check() error {
	if _, ok := uclc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "UserCommentLike.userId"`)}
	}
	if _, ok := uclc.mutation.CommentId(); !ok {
		return &ValidationError{Name: "commentId", err: errors.New(`ent: missing required field "UserCommentLike.commentId"`)}
	}
	if _, ok := uclc.mutation.LikedAt(); !ok {
		return &ValidationError{Name: "likedAt", err: errors.New(`ent: missing required field "UserCommentLike.likedAt"`)}
	}
	if _, ok := uclc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserCommentLike.user"`)}
	}
	if _, ok := uclc.mutation.CommentID(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required edge "UserCommentLike.comment"`)}
	}
	return nil
}

func (uclc *UserCommentLikeCreate) sqlSave(ctx context.Context) (*UserCommentLike, error) {
	if err := uclc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uclc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uclc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (uclc *UserCommentLikeCreate) createSpec() (*UserCommentLike, *sqlgraph.CreateSpec) {
	var (
		_node = &UserCommentLike{config: uclc.config}
		_spec = sqlgraph.NewCreateSpec(usercommentlike.Table, nil)
	)
	if value, ok := uclc.mutation.LikedAt(); ok {
		_spec.SetField(usercommentlike.FieldLikedAt, field.TypeTime, value)
		_node.LikedAt = value
	}
	if nodes := uclc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UserId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uclc.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usercommentlike.CommentTable,
			Columns: []string{usercommentlike.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CommentId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCommentLikeCreateBulk is the builder for creating many UserCommentLike entities in bulk.
type UserCommentLikeCreateBulk struct {
	config
	err      error
	builders []*UserCommentLikeCreate
}

// Save creates the UserCommentLike entities in the database.
func (uclcb *UserCommentLikeCreateBulk) Save(ctx context.Context) ([]*UserCommentLike, error) {
	if uclcb.err != nil {
		return nil, uclcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uclcb.builders))
	nodes := make([]*UserCommentLike, len(uclcb.builders))
	mutators := make([]Mutator, len(uclcb.builders))
	for i := range uclcb.builders {
		func(i int, root context.Context) {
			builder := uclcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserCommentLikeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uclcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uclcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uclcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uclcb *UserCommentLikeCreateBulk) SaveX(ctx context.Context) []*UserCommentLike {
	v, err := uclcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uclcb *UserCommentLikeCreateBulk) Exec(ctx context.Context) error {
	_, err := uclcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uclcb *UserCommentLikeCreateBulk) ExecX(ctx context.Context) {
	if err := uclcb.Exec(ctx); err != nil {
		panic(err)
	}
}