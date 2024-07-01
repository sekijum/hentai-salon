// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/threadcomment"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/usercommentlike"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// UserCommentLikeQuery is the builder for querying UserCommentLike entities.
type UserCommentLikeQuery struct {
	config
	ctx         *QueryContext
	order       []usercommentlike.OrderOption
	inters      []Interceptor
	predicates  []predicate.UserCommentLike
	withUser    *UserQuery
	withComment *ThreadCommentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCommentLikeQuery builder.
func (uclq *UserCommentLikeQuery) Where(ps ...predicate.UserCommentLike) *UserCommentLikeQuery {
	uclq.predicates = append(uclq.predicates, ps...)
	return uclq
}

// Limit the number of records to be returned by this query.
func (uclq *UserCommentLikeQuery) Limit(limit int) *UserCommentLikeQuery {
	uclq.ctx.Limit = &limit
	return uclq
}

// Offset to start from.
func (uclq *UserCommentLikeQuery) Offset(offset int) *UserCommentLikeQuery {
	uclq.ctx.Offset = &offset
	return uclq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uclq *UserCommentLikeQuery) Unique(unique bool) *UserCommentLikeQuery {
	uclq.ctx.Unique = &unique
	return uclq
}

// Order specifies how the records should be ordered.
func (uclq *UserCommentLikeQuery) Order(o ...usercommentlike.OrderOption) *UserCommentLikeQuery {
	uclq.order = append(uclq.order, o...)
	return uclq
}

// QueryUser chains the current query on the "user" edge.
func (uclq *UserCommentLikeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: uclq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uclq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uclq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercommentlike.Table, usercommentlike.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usercommentlike.UserTable, usercommentlike.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uclq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComment chains the current query on the "comment" edge.
func (uclq *UserCommentLikeQuery) QueryComment() *ThreadCommentQuery {
	query := (&ThreadCommentClient{config: uclq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uclq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uclq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercommentlike.Table, usercommentlike.CommentColumn, selector),
			sqlgraph.To(threadcomment.Table, threadcomment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usercommentlike.CommentTable, usercommentlike.CommentColumn),
		)
		fromU = sqlgraph.SetNeighbors(uclq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserCommentLike entity from the query.
// Returns a *NotFoundError when no UserCommentLike was found.
func (uclq *UserCommentLikeQuery) First(ctx context.Context) (*UserCommentLike, error) {
	nodes, err := uclq.Limit(1).All(setContextOp(ctx, uclq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercommentlike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uclq *UserCommentLikeQuery) FirstX(ctx context.Context) *UserCommentLike {
	node, err := uclq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single UserCommentLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserCommentLike entity is found.
// Returns a *NotFoundError when no UserCommentLike entities are found.
func (uclq *UserCommentLikeQuery) Only(ctx context.Context) (*UserCommentLike, error) {
	nodes, err := uclq.Limit(2).All(setContextOp(ctx, uclq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercommentlike.Label}
	default:
		return nil, &NotSingularError{usercommentlike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uclq *UserCommentLikeQuery) OnlyX(ctx context.Context) *UserCommentLike {
	node, err := uclq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of UserCommentLikes.
func (uclq *UserCommentLikeQuery) All(ctx context.Context) ([]*UserCommentLike, error) {
	ctx = setContextOp(ctx, uclq.ctx, "All")
	if err := uclq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserCommentLike, *UserCommentLikeQuery]()
	return withInterceptors[[]*UserCommentLike](ctx, uclq, qr, uclq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uclq *UserCommentLikeQuery) AllX(ctx context.Context) []*UserCommentLike {
	nodes, err := uclq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (uclq *UserCommentLikeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uclq.ctx, "Count")
	if err := uclq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uclq, querierCount[*UserCommentLikeQuery](), uclq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uclq *UserCommentLikeQuery) CountX(ctx context.Context) int {
	count, err := uclq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uclq *UserCommentLikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uclq.ctx, "Exist")
	switch _, err := uclq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uclq *UserCommentLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := uclq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCommentLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uclq *UserCommentLikeQuery) Clone() *UserCommentLikeQuery {
	if uclq == nil {
		return nil
	}
	return &UserCommentLikeQuery{
		config:      uclq.config,
		ctx:         uclq.ctx.Clone(),
		order:       append([]usercommentlike.OrderOption{}, uclq.order...),
		inters:      append([]Interceptor{}, uclq.inters...),
		predicates:  append([]predicate.UserCommentLike{}, uclq.predicates...),
		withUser:    uclq.withUser.Clone(),
		withComment: uclq.withComment.Clone(),
		// clone intermediate query.
		sql:  uclq.sql.Clone(),
		path: uclq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (uclq *UserCommentLikeQuery) WithUser(opts ...func(*UserQuery)) *UserCommentLikeQuery {
	query := (&UserClient{config: uclq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uclq.withUser = query
	return uclq
}

// WithComment tells the query-builder to eager-load the nodes that are connected to
// the "comment" edge. The optional arguments are used to configure the query builder of the edge.
func (uclq *UserCommentLikeQuery) WithComment(opts ...func(*ThreadCommentQuery)) *UserCommentLikeQuery {
	query := (&ThreadCommentClient{config: uclq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uclq.withComment = query
	return uclq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserId int `json:"userId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserCommentLike.Query().
//		GroupBy(usercommentlike.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uclq *UserCommentLikeQuery) GroupBy(field string, fields ...string) *UserCommentLikeGroupBy {
	uclq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserCommentLikeGroupBy{build: uclq}
	grbuild.flds = &uclq.ctx.Fields
	grbuild.label = usercommentlike.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserId int `json:"userId,omitempty"`
//	}
//
//	client.UserCommentLike.Query().
//		Select(usercommentlike.FieldUserId).
//		Scan(ctx, &v)
func (uclq *UserCommentLikeQuery) Select(fields ...string) *UserCommentLikeSelect {
	uclq.ctx.Fields = append(uclq.ctx.Fields, fields...)
	sbuild := &UserCommentLikeSelect{UserCommentLikeQuery: uclq}
	sbuild.label = usercommentlike.Label
	sbuild.flds, sbuild.scan = &uclq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserCommentLikeSelect configured with the given aggregations.
func (uclq *UserCommentLikeQuery) Aggregate(fns ...AggregateFunc) *UserCommentLikeSelect {
	return uclq.Select().Aggregate(fns...)
}

func (uclq *UserCommentLikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uclq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uclq); err != nil {
				return err
			}
		}
	}
	for _, f := range uclq.ctx.Fields {
		if !usercommentlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uclq.path != nil {
		prev, err := uclq.path(ctx)
		if err != nil {
			return err
		}
		uclq.sql = prev
	}
	return nil
}

func (uclq *UserCommentLikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserCommentLike, error) {
	var (
		nodes       = []*UserCommentLike{}
		_spec       = uclq.querySpec()
		loadedTypes = [2]bool{
			uclq.withUser != nil,
			uclq.withComment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserCommentLike).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserCommentLike{config: uclq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uclq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uclq.withUser; query != nil {
		if err := uclq.loadUser(ctx, query, nodes, nil,
			func(n *UserCommentLike, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := uclq.withComment; query != nil {
		if err := uclq.loadComment(ctx, query, nodes, nil,
			func(n *UserCommentLike, e *ThreadComment) { n.Edges.Comment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uclq *UserCommentLikeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserCommentLike, init func(*UserCommentLike), assign func(*UserCommentLike, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserCommentLike)
	for i := range nodes {
		fk := nodes[i].UserId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "userId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uclq *UserCommentLikeQuery) loadComment(ctx context.Context, query *ThreadCommentQuery, nodes []*UserCommentLike, init func(*UserCommentLike), assign func(*UserCommentLike, *ThreadComment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserCommentLike)
	for i := range nodes {
		fk := nodes[i].CommentId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(threadcomment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "commentId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uclq *UserCommentLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uclq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, uclq.driver, _spec)
}

func (uclq *UserCommentLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usercommentlike.Table, usercommentlike.Columns, nil)
	_spec.From = uclq.sql
	if unique := uclq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uclq.path != nil {
		_spec.Unique = true
	}
	if fields := uclq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if uclq.withUser != nil {
			_spec.Node.AddColumnOnce(usercommentlike.FieldUserId)
		}
		if uclq.withComment != nil {
			_spec.Node.AddColumnOnce(usercommentlike.FieldCommentId)
		}
	}
	if ps := uclq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uclq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uclq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uclq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uclq *UserCommentLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uclq.driver.Dialect())
	t1 := builder.Table(usercommentlike.Table)
	columns := uclq.ctx.Fields
	if len(columns) == 0 {
		columns = usercommentlike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uclq.sql != nil {
		selector = uclq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uclq.ctx.Unique != nil && *uclq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uclq.predicates {
		p(selector)
	}
	for _, p := range uclq.order {
		p(selector)
	}
	if offset := uclq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uclq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserCommentLikeGroupBy is the group-by builder for UserCommentLike entities.
type UserCommentLikeGroupBy struct {
	selector
	build *UserCommentLikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uclgb *UserCommentLikeGroupBy) Aggregate(fns ...AggregateFunc) *UserCommentLikeGroupBy {
	uclgb.fns = append(uclgb.fns, fns...)
	return uclgb
}

// Scan applies the selector query and scans the result into the given value.
func (uclgb *UserCommentLikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uclgb.build.ctx, "GroupBy")
	if err := uclgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCommentLikeQuery, *UserCommentLikeGroupBy](ctx, uclgb.build, uclgb, uclgb.build.inters, v)
}

func (uclgb *UserCommentLikeGroupBy) sqlScan(ctx context.Context, root *UserCommentLikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uclgb.fns))
	for _, fn := range uclgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uclgb.flds)+len(uclgb.fns))
		for _, f := range *uclgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uclgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uclgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserCommentLikeSelect is the builder for selecting fields of UserCommentLike entities.
type UserCommentLikeSelect struct {
	*UserCommentLikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ucls *UserCommentLikeSelect) Aggregate(fns ...AggregateFunc) *UserCommentLikeSelect {
	ucls.fns = append(ucls.fns, fns...)
	return ucls
}

// Scan applies the selector query and scans the result into the given value.
func (ucls *UserCommentLikeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucls.ctx, "Select")
	if err := ucls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCommentLikeQuery, *UserCommentLikeSelect](ctx, ucls.UserCommentLikeQuery, ucls, ucls.inters, v)
}

func (ucls *UserCommentLikeSelect) sqlScan(ctx context.Context, root *UserCommentLikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ucls.fns))
	for _, fn := range ucls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ucls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
