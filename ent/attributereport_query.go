// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/in-toto/archivista/ent/attributeassertion"
	"github.com/in-toto/archivista/ent/attributereport"
	"github.com/in-toto/archivista/ent/predicate"
	"github.com/in-toto/archivista/ent/statement"
)

// AttributeReportQuery is the builder for querying AttributeReport entities.
type AttributeReportQuery struct {
	config
	ctx                 *QueryContext
	order               []attributereport.OrderOption
	inters              []Interceptor
	predicates          []predicate.AttributeReport
	withAttributes      *AttributeAssertionQuery
	withStatement       *StatementQuery
	withFKs             bool
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*AttributeReport) error
	withNamedAttributes map[string]*AttributeAssertionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeReportQuery builder.
func (arq *AttributeReportQuery) Where(ps ...predicate.AttributeReport) *AttributeReportQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit the number of records to be returned by this query.
func (arq *AttributeReportQuery) Limit(limit int) *AttributeReportQuery {
	arq.ctx.Limit = &limit
	return arq
}

// Offset to start from.
func (arq *AttributeReportQuery) Offset(offset int) *AttributeReportQuery {
	arq.ctx.Offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AttributeReportQuery) Unique(unique bool) *AttributeReportQuery {
	arq.ctx.Unique = &unique
	return arq
}

// Order specifies how the records should be ordered.
func (arq *AttributeReportQuery) Order(o ...attributereport.OrderOption) *AttributeReportQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryAttributes chains the current query on the "attributes" edge.
func (arq *AttributeReportQuery) QueryAttributes() *AttributeAssertionQuery {
	query := (&AttributeAssertionClient{config: arq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributereport.Table, attributereport.FieldID, selector),
			sqlgraph.To(attributeassertion.Table, attributeassertion.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, attributereport.AttributesTable, attributereport.AttributesColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatement chains the current query on the "statement" edge.
func (arq *AttributeReportQuery) QueryStatement() *StatementQuery {
	query := (&StatementClient{config: arq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributereport.Table, attributereport.FieldID, selector),
			sqlgraph.To(statement.Table, statement.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, attributereport.StatementTable, attributereport.StatementColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttributeReport entity from the query.
// Returns a *NotFoundError when no AttributeReport was found.
func (arq *AttributeReportQuery) First(ctx context.Context) (*AttributeReport, error) {
	nodes, err := arq.Limit(1).All(setContextOp(ctx, arq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attributereport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AttributeReportQuery) FirstX(ctx context.Context) *AttributeReport {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttributeReport ID from the query.
// Returns a *NotFoundError when no AttributeReport ID was found.
func (arq *AttributeReportQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(1).IDs(setContextOp(ctx, arq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attributereport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AttributeReportQuery) FirstIDX(ctx context.Context) int {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttributeReport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttributeReport entity is found.
// Returns a *NotFoundError when no AttributeReport entities are found.
func (arq *AttributeReportQuery) Only(ctx context.Context) (*AttributeReport, error) {
	nodes, err := arq.Limit(2).All(setContextOp(ctx, arq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attributereport.Label}
	default:
		return nil, &NotSingularError{attributereport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AttributeReportQuery) OnlyX(ctx context.Context) *AttributeReport {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttributeReport ID in the query.
// Returns a *NotSingularError when more than one AttributeReport ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AttributeReportQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(2).IDs(setContextOp(ctx, arq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attributereport.Label}
	default:
		err = &NotSingularError{attributereport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AttributeReportQuery) OnlyIDX(ctx context.Context) int {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttributeReports.
func (arq *AttributeReportQuery) All(ctx context.Context) ([]*AttributeReport, error) {
	ctx = setContextOp(ctx, arq.ctx, "All")
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttributeReport, *AttributeReportQuery]()
	return withInterceptors[[]*AttributeReport](ctx, arq, qr, arq.inters)
}

// AllX is like All, but panics if an error occurs.
func (arq *AttributeReportQuery) AllX(ctx context.Context) []*AttributeReport {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttributeReport IDs.
func (arq *AttributeReportQuery) IDs(ctx context.Context) (ids []int, err error) {
	if arq.ctx.Unique == nil && arq.path != nil {
		arq.Unique(true)
	}
	ctx = setContextOp(ctx, arq.ctx, "IDs")
	if err = arq.Select(attributereport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AttributeReportQuery) IDsX(ctx context.Context) []int {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AttributeReportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, arq.ctx, "Count")
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, arq, querierCount[*AttributeReportQuery](), arq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AttributeReportQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AttributeReportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, arq.ctx, "Exist")
	switch _, err := arq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AttributeReportQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeReportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AttributeReportQuery) Clone() *AttributeReportQuery {
	if arq == nil {
		return nil
	}
	return &AttributeReportQuery{
		config:         arq.config,
		ctx:            arq.ctx.Clone(),
		order:          append([]attributereport.OrderOption{}, arq.order...),
		inters:         append([]Interceptor{}, arq.inters...),
		predicates:     append([]predicate.AttributeReport{}, arq.predicates...),
		withAttributes: arq.withAttributes.Clone(),
		withStatement:  arq.withStatement.Clone(),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// WithAttributes tells the query-builder to eager-load the nodes that are connected to
// the "attributes" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AttributeReportQuery) WithAttributes(opts ...func(*AttributeAssertionQuery)) *AttributeReportQuery {
	query := (&AttributeAssertionClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	arq.withAttributes = query
	return arq
}

// WithStatement tells the query-builder to eager-load the nodes that are connected to
// the "statement" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AttributeReportQuery) WithStatement(opts ...func(*StatementQuery)) *AttributeReportQuery {
	query := (&StatementClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	arq.withStatement = query
	return arq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (arq *AttributeReportQuery) GroupBy(field string, fields ...string) *AttributeReportGroupBy {
	arq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeReportGroupBy{build: arq}
	grbuild.flds = &arq.ctx.Fields
	grbuild.label = attributereport.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (arq *AttributeReportQuery) Select(fields ...string) *AttributeReportSelect {
	arq.ctx.Fields = append(arq.ctx.Fields, fields...)
	sbuild := &AttributeReportSelect{AttributeReportQuery: arq}
	sbuild.label = attributereport.Label
	sbuild.flds, sbuild.scan = &arq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeReportSelect configured with the given aggregations.
func (arq *AttributeReportQuery) Aggregate(fns ...AggregateFunc) *AttributeReportSelect {
	return arq.Select().Aggregate(fns...)
}

func (arq *AttributeReportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range arq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, arq); err != nil {
				return err
			}
		}
	}
	for _, f := range arq.ctx.Fields {
		if !attributereport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AttributeReportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttributeReport, error) {
	var (
		nodes       = []*AttributeReport{}
		withFKs     = arq.withFKs
		_spec       = arq.querySpec()
		loadedTypes = [2]bool{
			arq.withAttributes != nil,
			arq.withStatement != nil,
		}
	)
	if arq.withStatement != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attributereport.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttributeReport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttributeReport{config: arq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := arq.withAttributes; query != nil {
		if err := arq.loadAttributes(ctx, query, nodes,
			func(n *AttributeReport) { n.Edges.Attributes = []*AttributeAssertion{} },
			func(n *AttributeReport, e *AttributeAssertion) { n.Edges.Attributes = append(n.Edges.Attributes, e) }); err != nil {
			return nil, err
		}
	}
	if query := arq.withStatement; query != nil {
		if err := arq.loadStatement(ctx, query, nodes, nil,
			func(n *AttributeReport, e *Statement) { n.Edges.Statement = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range arq.withNamedAttributes {
		if err := arq.loadAttributes(ctx, query, nodes,
			func(n *AttributeReport) { n.appendNamedAttributes(name) },
			func(n *AttributeReport, e *AttributeAssertion) { n.appendNamedAttributes(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range arq.loadTotal {
		if err := arq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (arq *AttributeReportQuery) loadAttributes(ctx context.Context, query *AttributeAssertionQuery, nodes []*AttributeReport, init func(*AttributeReport), assign func(*AttributeReport, *AttributeAssertion)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*AttributeReport)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AttributeAssertion(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(attributereport.AttributesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.attribute_report_attributes
		if fk == nil {
			return fmt.Errorf(`foreign-key "attribute_report_attributes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "attribute_report_attributes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (arq *AttributeReportQuery) loadStatement(ctx context.Context, query *StatementQuery, nodes []*AttributeReport, init func(*AttributeReport), assign func(*AttributeReport, *Statement)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AttributeReport)
	for i := range nodes {
		if nodes[i].statement_attributes_report == nil {
			continue
		}
		fk := *nodes[i].statement_attributes_report
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(statement.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "statement_attributes_report" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (arq *AttributeReportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	_spec.Node.Columns = arq.ctx.Fields
	if len(arq.ctx.Fields) > 0 {
		_spec.Unique = arq.ctx.Unique != nil && *arq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AttributeReportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attributereport.Table, attributereport.Columns, sqlgraph.NewFieldSpec(attributereport.FieldID, field.TypeInt))
	_spec.From = arq.sql
	if unique := arq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if arq.path != nil {
		_spec.Unique = true
	}
	if fields := arq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributereport.FieldID)
		for i := range fields {
			if fields[i] != attributereport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AttributeReportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(attributereport.Table)
	columns := arq.ctx.Fields
	if len(columns) == 0 {
		columns = attributereport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.ctx.Unique != nil && *arq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAttributes tells the query-builder to eager-load the nodes that are connected to the "attributes"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (arq *AttributeReportQuery) WithNamedAttributes(name string, opts ...func(*AttributeAssertionQuery)) *AttributeReportQuery {
	query := (&AttributeAssertionClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if arq.withNamedAttributes == nil {
		arq.withNamedAttributes = make(map[string]*AttributeAssertionQuery)
	}
	arq.withNamedAttributes[name] = query
	return arq
}

// AttributeReportGroupBy is the group-by builder for AttributeReport entities.
type AttributeReportGroupBy struct {
	selector
	build *AttributeReportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AttributeReportGroupBy) Aggregate(fns ...AggregateFunc) *AttributeReportGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the selector query and scans the result into the given value.
func (argb *AttributeReportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, argb.build.ctx, "GroupBy")
	if err := argb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeReportQuery, *AttributeReportGroupBy](ctx, argb.build, argb, argb.build.inters, v)
}

func (argb *AttributeReportGroupBy) sqlScan(ctx context.Context, root *AttributeReportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*argb.flds)+len(argb.fns))
		for _, f := range *argb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*argb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeReportSelect is the builder for selecting fields of AttributeReport entities.
type AttributeReportSelect struct {
	*AttributeReportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ars *AttributeReportSelect) Aggregate(fns ...AggregateFunc) *AttributeReportSelect {
	ars.fns = append(ars.fns, fns...)
	return ars
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AttributeReportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ars.ctx, "Select")
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeReportQuery, *AttributeReportSelect](ctx, ars.AttributeReportQuery, ars, ars.inters, v)
}

func (ars *AttributeReportSelect) sqlScan(ctx context.Context, root *AttributeReportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ars.fns))
	for _, fn := range ars.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ars.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}