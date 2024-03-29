// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/LintaoAmons/undercontrol/ent/account"
	"github.com/LintaoAmons/undercontrol/ent/accounthistory"
	"github.com/LintaoAmons/undercontrol/ent/predicate"
)

// AccountHistoryQuery is the builder for querying AccountHistory entities.
type AccountHistoryQuery struct {
	config
	ctx         *QueryContext
	order       []accounthistory.OrderOption
	inters      []Interceptor
	predicates  []predicate.AccountHistory
	withAccount *AccountQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountHistoryQuery builder.
func (ahq *AccountHistoryQuery) Where(ps ...predicate.AccountHistory) *AccountHistoryQuery {
	ahq.predicates = append(ahq.predicates, ps...)
	return ahq
}

// Limit the number of records to be returned by this query.
func (ahq *AccountHistoryQuery) Limit(limit int) *AccountHistoryQuery {
	ahq.ctx.Limit = &limit
	return ahq
}

// Offset to start from.
func (ahq *AccountHistoryQuery) Offset(offset int) *AccountHistoryQuery {
	ahq.ctx.Offset = &offset
	return ahq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ahq *AccountHistoryQuery) Unique(unique bool) *AccountHistoryQuery {
	ahq.ctx.Unique = &unique
	return ahq
}

// Order specifies how the records should be ordered.
func (ahq *AccountHistoryQuery) Order(o ...accounthistory.OrderOption) *AccountHistoryQuery {
	ahq.order = append(ahq.order, o...)
	return ahq
}

// QueryAccount chains the current query on the "account" edge.
func (ahq *AccountHistoryQuery) QueryAccount() *AccountQuery {
	query := (&AccountClient{config: ahq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ahq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ahq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accounthistory.Table, accounthistory.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accounthistory.AccountTable, accounthistory.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(ahq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccountHistory entity from the query.
// Returns a *NotFoundError when no AccountHistory was found.
func (ahq *AccountHistoryQuery) First(ctx context.Context) (*AccountHistory, error) {
	nodes, err := ahq.Limit(1).All(setContextOp(ctx, ahq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accounthistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ahq *AccountHistoryQuery) FirstX(ctx context.Context) *AccountHistory {
	node, err := ahq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountHistory ID from the query.
// Returns a *NotFoundError when no AccountHistory ID was found.
func (ahq *AccountHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ahq.Limit(1).IDs(setContextOp(ctx, ahq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accounthistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ahq *AccountHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := ahq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccountHistory entity is found.
// Returns a *NotFoundError when no AccountHistory entities are found.
func (ahq *AccountHistoryQuery) Only(ctx context.Context) (*AccountHistory, error) {
	nodes, err := ahq.Limit(2).All(setContextOp(ctx, ahq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accounthistory.Label}
	default:
		return nil, &NotSingularError{accounthistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ahq *AccountHistoryQuery) OnlyX(ctx context.Context) *AccountHistory {
	node, err := ahq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountHistory ID in the query.
// Returns a *NotSingularError when more than one AccountHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ahq *AccountHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ahq.Limit(2).IDs(setContextOp(ctx, ahq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accounthistory.Label}
	default:
		err = &NotSingularError{accounthistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ahq *AccountHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := ahq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountHistories.
func (ahq *AccountHistoryQuery) All(ctx context.Context) ([]*AccountHistory, error) {
	ctx = setContextOp(ctx, ahq.ctx, "All")
	if err := ahq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccountHistory, *AccountHistoryQuery]()
	return withInterceptors[[]*AccountHistory](ctx, ahq, qr, ahq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ahq *AccountHistoryQuery) AllX(ctx context.Context) []*AccountHistory {
	nodes, err := ahq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountHistory IDs.
func (ahq *AccountHistoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ahq.ctx.Unique == nil && ahq.path != nil {
		ahq.Unique(true)
	}
	ctx = setContextOp(ctx, ahq.ctx, "IDs")
	if err = ahq.Select(accounthistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ahq *AccountHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := ahq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ahq *AccountHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ahq.ctx, "Count")
	if err := ahq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ahq, querierCount[*AccountHistoryQuery](), ahq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ahq *AccountHistoryQuery) CountX(ctx context.Context) int {
	count, err := ahq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ahq *AccountHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ahq.ctx, "Exist")
	switch _, err := ahq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ahq *AccountHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ahq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ahq *AccountHistoryQuery) Clone() *AccountHistoryQuery {
	if ahq == nil {
		return nil
	}
	return &AccountHistoryQuery{
		config:      ahq.config,
		ctx:         ahq.ctx.Clone(),
		order:       append([]accounthistory.OrderOption{}, ahq.order...),
		inters:      append([]Interceptor{}, ahq.inters...),
		predicates:  append([]predicate.AccountHistory{}, ahq.predicates...),
		withAccount: ahq.withAccount.Clone(),
		// clone intermediate query.
		sql:  ahq.sql.Clone(),
		path: ahq.path,
	}
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (ahq *AccountHistoryQuery) WithAccount(opts ...func(*AccountQuery)) *AccountHistoryQuery {
	query := (&AccountClient{config: ahq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ahq.withAccount = query
	return ahq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AccountHistory.Query().
//		GroupBy(accounthistory.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ahq *AccountHistoryQuery) GroupBy(field string, fields ...string) *AccountHistoryGroupBy {
	ahq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccountHistoryGroupBy{build: ahq}
	grbuild.flds = &ahq.ctx.Fields
	grbuild.label = accounthistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.AccountHistory.Query().
//		Select(accounthistory.FieldName).
//		Scan(ctx, &v)
func (ahq *AccountHistoryQuery) Select(fields ...string) *AccountHistorySelect {
	ahq.ctx.Fields = append(ahq.ctx.Fields, fields...)
	sbuild := &AccountHistorySelect{AccountHistoryQuery: ahq}
	sbuild.label = accounthistory.Label
	sbuild.flds, sbuild.scan = &ahq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccountHistorySelect configured with the given aggregations.
func (ahq *AccountHistoryQuery) Aggregate(fns ...AggregateFunc) *AccountHistorySelect {
	return ahq.Select().Aggregate(fns...)
}

func (ahq *AccountHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ahq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ahq); err != nil {
				return err
			}
		}
	}
	for _, f := range ahq.ctx.Fields {
		if !accounthistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ahq.path != nil {
		prev, err := ahq.path(ctx)
		if err != nil {
			return err
		}
		ahq.sql = prev
	}
	return nil
}

func (ahq *AccountHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccountHistory, error) {
	var (
		nodes       = []*AccountHistory{}
		withFKs     = ahq.withFKs
		_spec       = ahq.querySpec()
		loadedTypes = [1]bool{
			ahq.withAccount != nil,
		}
	)
	if ahq.withAccount != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, accounthistory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccountHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccountHistory{config: ahq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ahq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ahq.withAccount; query != nil {
		if err := ahq.loadAccount(ctx, query, nodes, nil,
			func(n *AccountHistory, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ahq *AccountHistoryQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*AccountHistory, init func(*AccountHistory), assign func(*AccountHistory, *Account)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AccountHistory)
	for i := range nodes {
		if nodes[i].account_accounthistory == nil {
			continue
		}
		fk := *nodes[i].account_accounthistory
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_accounthistory" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ahq *AccountHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ahq.querySpec()
	_spec.Node.Columns = ahq.ctx.Fields
	if len(ahq.ctx.Fields) > 0 {
		_spec.Unique = ahq.ctx.Unique != nil && *ahq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ahq.driver, _spec)
}

func (ahq *AccountHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accounthistory.Table, accounthistory.Columns, sqlgraph.NewFieldSpec(accounthistory.FieldID, field.TypeInt))
	_spec.From = ahq.sql
	if unique := ahq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ahq.path != nil {
		_spec.Unique = true
	}
	if fields := ahq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accounthistory.FieldID)
		for i := range fields {
			if fields[i] != accounthistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ahq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ahq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ahq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ahq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ahq *AccountHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ahq.driver.Dialect())
	t1 := builder.Table(accounthistory.Table)
	columns := ahq.ctx.Fields
	if len(columns) == 0 {
		columns = accounthistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ahq.sql != nil {
		selector = ahq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ahq.ctx.Unique != nil && *ahq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ahq.predicates {
		p(selector)
	}
	for _, p := range ahq.order {
		p(selector)
	}
	if offset := ahq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ahq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountHistoryGroupBy is the group-by builder for AccountHistory entities.
type AccountHistoryGroupBy struct {
	selector
	build *AccountHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ahgb *AccountHistoryGroupBy) Aggregate(fns ...AggregateFunc) *AccountHistoryGroupBy {
	ahgb.fns = append(ahgb.fns, fns...)
	return ahgb
}

// Scan applies the selector query and scans the result into the given value.
func (ahgb *AccountHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ahgb.build.ctx, "GroupBy")
	if err := ahgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountHistoryQuery, *AccountHistoryGroupBy](ctx, ahgb.build, ahgb, ahgb.build.inters, v)
}

func (ahgb *AccountHistoryGroupBy) sqlScan(ctx context.Context, root *AccountHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ahgb.fns))
	for _, fn := range ahgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ahgb.flds)+len(ahgb.fns))
		for _, f := range *ahgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ahgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ahgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccountHistorySelect is the builder for selecting fields of AccountHistory entities.
type AccountHistorySelect struct {
	*AccountHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ahs *AccountHistorySelect) Aggregate(fns ...AggregateFunc) *AccountHistorySelect {
	ahs.fns = append(ahs.fns, fns...)
	return ahs
}

// Scan applies the selector query and scans the result into the given value.
func (ahs *AccountHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ahs.ctx, "Select")
	if err := ahs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountHistoryQuery, *AccountHistorySelect](ctx, ahs.AccountHistoryQuery, ahs, ahs.inters, v)
}

func (ahs *AccountHistorySelect) sqlScan(ctx context.Context, root *AccountHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ahs.fns))
	for _, fn := range ahs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ahs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ahs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
