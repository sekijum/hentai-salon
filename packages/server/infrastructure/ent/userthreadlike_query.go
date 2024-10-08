// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/thread"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userthreadlike"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// UserThreadLikeQuery is the builder for querying UserThreadLike entities.
type UserThreadLikeQuery struct {
	config
	ctx        *QueryContext
	order      []userthreadlike.OrderOption
	inters     []Interceptor
	predicates []predicate.UserThreadLike
	withUser   *UserQuery
	withThread *ThreadQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserThreadLikeQuery builder.
func (utlq *UserThreadLikeQuery) Where(ps ...predicate.UserThreadLike) *UserThreadLikeQuery {
	utlq.predicates = append(utlq.predicates, ps...)
	return utlq
}

// Limit the number of records to be returned by this query.
func (utlq *UserThreadLikeQuery) Limit(limit int) *UserThreadLikeQuery {
	utlq.ctx.Limit = &limit
	return utlq
}

// Offset to start from.
func (utlq *UserThreadLikeQuery) Offset(offset int) *UserThreadLikeQuery {
	utlq.ctx.Offset = &offset
	return utlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (utlq *UserThreadLikeQuery) Unique(unique bool) *UserThreadLikeQuery {
	utlq.ctx.Unique = &unique
	return utlq
}

// Order specifies how the records should be ordered.
func (utlq *UserThreadLikeQuery) Order(o ...userthreadlike.OrderOption) *UserThreadLikeQuery {
	utlq.order = append(utlq.order, o...)
	return utlq
}

// QueryUser chains the current query on the "user" edge.
func (utlq *UserThreadLikeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: utlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userthreadlike.Table, userthreadlike.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userthreadlike.UserTable, userthreadlike.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(utlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThread chains the current query on the "thread" edge.
func (utlq *UserThreadLikeQuery) QueryThread() *ThreadQuery {
	query := (&ThreadClient{config: utlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userthreadlike.Table, userthreadlike.ThreadColumn, selector),
			sqlgraph.To(thread.Table, thread.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userthreadlike.ThreadTable, userthreadlike.ThreadColumn),
		)
		fromU = sqlgraph.SetNeighbors(utlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserThreadLike entity from the query.
// Returns a *NotFoundError when no UserThreadLike was found.
func (utlq *UserThreadLikeQuery) First(ctx context.Context) (*UserThreadLike, error) {
	nodes, err := utlq.Limit(1).All(setContextOp(ctx, utlq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userthreadlike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utlq *UserThreadLikeQuery) FirstX(ctx context.Context) *UserThreadLike {
	node, err := utlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single UserThreadLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserThreadLike entity is found.
// Returns a *NotFoundError when no UserThreadLike entities are found.
func (utlq *UserThreadLikeQuery) Only(ctx context.Context) (*UserThreadLike, error) {
	nodes, err := utlq.Limit(2).All(setContextOp(ctx, utlq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userthreadlike.Label}
	default:
		return nil, &NotSingularError{userthreadlike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utlq *UserThreadLikeQuery) OnlyX(ctx context.Context) *UserThreadLike {
	node, err := utlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of UserThreadLikes.
func (utlq *UserThreadLikeQuery) All(ctx context.Context) ([]*UserThreadLike, error) {
	ctx = setContextOp(ctx, utlq.ctx, ent.OpQueryAll)
	if err := utlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserThreadLike, *UserThreadLikeQuery]()
	return withInterceptors[[]*UserThreadLike](ctx, utlq, qr, utlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (utlq *UserThreadLikeQuery) AllX(ctx context.Context) []*UserThreadLike {
	nodes, err := utlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (utlq *UserThreadLikeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, utlq.ctx, ent.OpQueryCount)
	if err := utlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, utlq, querierCount[*UserThreadLikeQuery](), utlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (utlq *UserThreadLikeQuery) CountX(ctx context.Context) int {
	count, err := utlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utlq *UserThreadLikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, utlq.ctx, ent.OpQueryExist)
	switch _, err := utlq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (utlq *UserThreadLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := utlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserThreadLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utlq *UserThreadLikeQuery) Clone() *UserThreadLikeQuery {
	if utlq == nil {
		return nil
	}
	return &UserThreadLikeQuery{
		config:     utlq.config,
		ctx:        utlq.ctx.Clone(),
		order:      append([]userthreadlike.OrderOption{}, utlq.order...),
		inters:     append([]Interceptor{}, utlq.inters...),
		predicates: append([]predicate.UserThreadLike{}, utlq.predicates...),
		withUser:   utlq.withUser.Clone(),
		withThread: utlq.withThread.Clone(),
		// clone intermediate query.
		sql:  utlq.sql.Clone(),
		path: utlq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (utlq *UserThreadLikeQuery) WithUser(opts ...func(*UserQuery)) *UserThreadLikeQuery {
	query := (&UserClient{config: utlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utlq.withUser = query
	return utlq
}

// WithThread tells the query-builder to eager-load the nodes that are connected to
// the "thread" edge. The optional arguments are used to configure the query builder of the edge.
func (utlq *UserThreadLikeQuery) WithThread(opts ...func(*ThreadQuery)) *UserThreadLikeQuery {
	query := (&ThreadClient{config: utlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utlq.withThread = query
	return utlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserThreadLike.Query().
//		GroupBy(userthreadlike.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (utlq *UserThreadLikeQuery) GroupBy(field string, fields ...string) *UserThreadLikeGroupBy {
	utlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserThreadLikeGroupBy{build: utlq}
	grbuild.flds = &utlq.ctx.Fields
	grbuild.label = userthreadlike.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.UserThreadLike.Query().
//		Select(userthreadlike.FieldUserID).
//		Scan(ctx, &v)
func (utlq *UserThreadLikeQuery) Select(fields ...string) *UserThreadLikeSelect {
	utlq.ctx.Fields = append(utlq.ctx.Fields, fields...)
	sbuild := &UserThreadLikeSelect{UserThreadLikeQuery: utlq}
	sbuild.label = userthreadlike.Label
	sbuild.flds, sbuild.scan = &utlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserThreadLikeSelect configured with the given aggregations.
func (utlq *UserThreadLikeQuery) Aggregate(fns ...AggregateFunc) *UserThreadLikeSelect {
	return utlq.Select().Aggregate(fns...)
}

func (utlq *UserThreadLikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range utlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, utlq); err != nil {
				return err
			}
		}
	}
	for _, f := range utlq.ctx.Fields {
		if !userthreadlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if utlq.path != nil {
		prev, err := utlq.path(ctx)
		if err != nil {
			return err
		}
		utlq.sql = prev
	}
	return nil
}

func (utlq *UserThreadLikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserThreadLike, error) {
	var (
		nodes       = []*UserThreadLike{}
		_spec       = utlq.querySpec()
		loadedTypes = [2]bool{
			utlq.withUser != nil,
			utlq.withThread != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserThreadLike).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserThreadLike{config: utlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, utlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := utlq.withUser; query != nil {
		if err := utlq.loadUser(ctx, query, nodes, nil,
			func(n *UserThreadLike, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := utlq.withThread; query != nil {
		if err := utlq.loadThread(ctx, query, nodes, nil,
			func(n *UserThreadLike, e *Thread) { n.Edges.Thread = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (utlq *UserThreadLikeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserThreadLike, init func(*UserThreadLike), assign func(*UserThreadLike, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserThreadLike)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (utlq *UserThreadLikeQuery) loadThread(ctx context.Context, query *ThreadQuery, nodes []*UserThreadLike, init func(*UserThreadLike), assign func(*UserThreadLike, *Thread)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserThreadLike)
	for i := range nodes {
		fk := nodes[i].ThreadID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(thread.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "thread_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (utlq *UserThreadLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utlq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, utlq.driver, _spec)
}

func (utlq *UserThreadLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userthreadlike.Table, userthreadlike.Columns, nil)
	_spec.From = utlq.sql
	if unique := utlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if utlq.path != nil {
		_spec.Unique = true
	}
	if fields := utlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if utlq.withUser != nil {
			_spec.Node.AddColumnOnce(userthreadlike.FieldUserID)
		}
		if utlq.withThread != nil {
			_spec.Node.AddColumnOnce(userthreadlike.FieldThreadID)
		}
	}
	if ps := utlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utlq *UserThreadLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(utlq.driver.Dialect())
	t1 := builder.Table(userthreadlike.Table)
	columns := utlq.ctx.Fields
	if len(columns) == 0 {
		columns = userthreadlike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if utlq.sql != nil {
		selector = utlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if utlq.ctx.Unique != nil && *utlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range utlq.predicates {
		p(selector)
	}
	for _, p := range utlq.order {
		p(selector)
	}
	if offset := utlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserThreadLikeGroupBy is the group-by builder for UserThreadLike entities.
type UserThreadLikeGroupBy struct {
	selector
	build *UserThreadLikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utlgb *UserThreadLikeGroupBy) Aggregate(fns ...AggregateFunc) *UserThreadLikeGroupBy {
	utlgb.fns = append(utlgb.fns, fns...)
	return utlgb
}

// Scan applies the selector query and scans the result into the given value.
func (utlgb *UserThreadLikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utlgb.build.ctx, ent.OpQueryGroupBy)
	if err := utlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserThreadLikeQuery, *UserThreadLikeGroupBy](ctx, utlgb.build, utlgb, utlgb.build.inters, v)
}

func (utlgb *UserThreadLikeGroupBy) sqlScan(ctx context.Context, root *UserThreadLikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(utlgb.fns))
	for _, fn := range utlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*utlgb.flds)+len(utlgb.fns))
		for _, f := range *utlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*utlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserThreadLikeSelect is the builder for selecting fields of UserThreadLike entities.
type UserThreadLikeSelect struct {
	*UserThreadLikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (utls *UserThreadLikeSelect) Aggregate(fns ...AggregateFunc) *UserThreadLikeSelect {
	utls.fns = append(utls.fns, fns...)
	return utls
}

// Scan applies the selector query and scans the result into the given value.
func (utls *UserThreadLikeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utls.ctx, ent.OpQuerySelect)
	if err := utls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserThreadLikeQuery, *UserThreadLikeSelect](ctx, utls.UserThreadLikeQuery, utls, utls.inters, v)
}

func (utls *UserThreadLikeSelect) sqlScan(ctx context.Context, root *UserThreadLikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(utls.fns))
	for _, fn := range utls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*utls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
