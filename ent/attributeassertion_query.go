// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/in-toto/archivista/ent/attributeassertion"
	"github.com/in-toto/archivista/ent/attributereport"
	"github.com/in-toto/archivista/ent/predicate"
)

// AttributeAssertionQuery is the builder for querying AttributeAssertion entities.
type AttributeAssertionQuery struct {
	config
	ctx                  *QueryContext
	order                []attributeassertion.OrderOption
	inters               []Interceptor
	predicates           []predicate.AttributeAssertion
	withAttributesReport *AttributeReportQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*AttributeAssertion) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeAssertionQuery builder.
func (aaq *AttributeAssertionQuery) Where(ps ...predicate.AttributeAssertion) *AttributeAssertionQuery {
	aaq.predicates = append(aaq.predicates, ps...)
	return aaq
}

// Limit the number of records to be returned by this query.
func (aaq *AttributeAssertionQuery) Limit(limit int) *AttributeAssertionQuery {
	aaq.ctx.Limit = &limit
	return aaq
}

// Offset to start from.
func (aaq *AttributeAssertionQuery) Offset(offset int) *AttributeAssertionQuery {
	aaq.ctx.Offset = &offset
	return aaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aaq *AttributeAssertionQuery) Unique(unique bool) *AttributeAssertionQuery {
	aaq.ctx.Unique = &unique
	return aaq
}

// Order specifies how the records should be ordered.
func (aaq *AttributeAssertionQuery) Order(o ...attributeassertion.OrderOption) *AttributeAssertionQuery {
	aaq.order = append(aaq.order, o...)
	return aaq
}

// QueryAttributesReport chains the current query on the "attributes_report" edge.
func (aaq *AttributeAssertionQuery) QueryAttributesReport() *AttributeReportQuery {
	query := (&AttributeReportClient{config: aaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributeassertion.Table, attributeassertion.FieldID, selector),
			sqlgraph.To(attributereport.Table, attributereport.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attributeassertion.AttributesReportTable, attributeassertion.AttributesReportColumn),
		)
		fromU = sqlgraph.SetNeighbors(aaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttributeAssertion entity from the query.
// Returns a *NotFoundError when no AttributeAssertion was found.
func (aaq *AttributeAssertionQuery) First(ctx context.Context) (*AttributeAssertion, error) {
	nodes, err := aaq.Limit(1).All(setContextOp(ctx, aaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attributeassertion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) FirstX(ctx context.Context) *AttributeAssertion {
	node, err := aaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttributeAssertion ID from the query.
// Returns a *NotFoundError when no AttributeAssertion ID was found.
func (aaq *AttributeAssertionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aaq.Limit(1).IDs(setContextOp(ctx, aaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attributeassertion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) FirstIDX(ctx context.Context) int {
	id, err := aaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttributeAssertion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttributeAssertion entity is found.
// Returns a *NotFoundError when no AttributeAssertion entities are found.
func (aaq *AttributeAssertionQuery) Only(ctx context.Context) (*AttributeAssertion, error) {
	nodes, err := aaq.Limit(2).All(setContextOp(ctx, aaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attributeassertion.Label}
	default:
		return nil, &NotSingularError{attributeassertion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) OnlyX(ctx context.Context) *AttributeAssertion {
	node, err := aaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttributeAssertion ID in the query.
// Returns a *NotSingularError when more than one AttributeAssertion ID is found.
// Returns a *NotFoundError when no entities are found.
func (aaq *AttributeAssertionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aaq.Limit(2).IDs(setContextOp(ctx, aaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attributeassertion.Label}
	default:
		err = &NotSingularError{attributeassertion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) OnlyIDX(ctx context.Context) int {
	id, err := aaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttributeAssertions.
func (aaq *AttributeAssertionQuery) All(ctx context.Context) ([]*AttributeAssertion, error) {
	ctx = setContextOp(ctx, aaq.ctx, "All")
	if err := aaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttributeAssertion, *AttributeAssertionQuery]()
	return withInterceptors[[]*AttributeAssertion](ctx, aaq, qr, aaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) AllX(ctx context.Context) []*AttributeAssertion {
	nodes, err := aaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttributeAssertion IDs.
func (aaq *AttributeAssertionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aaq.ctx.Unique == nil && aaq.path != nil {
		aaq.Unique(true)
	}
	ctx = setContextOp(ctx, aaq.ctx, "IDs")
	if err = aaq.Select(attributeassertion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) IDsX(ctx context.Context) []int {
	ids, err := aaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aaq *AttributeAssertionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aaq.ctx, "Count")
	if err := aaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aaq, querierCount[*AttributeAssertionQuery](), aaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) CountX(ctx context.Context) int {
	count, err := aaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aaq *AttributeAssertionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aaq.ctx, "Exist")
	switch _, err := aaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aaq *AttributeAssertionQuery) ExistX(ctx context.Context) bool {
	exist, err := aaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeAssertionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aaq *AttributeAssertionQuery) Clone() *AttributeAssertionQuery {
	if aaq == nil {
		return nil
	}
	return &AttributeAssertionQuery{
		config:               aaq.config,
		ctx:                  aaq.ctx.Clone(),
		order:                append([]attributeassertion.OrderOption{}, aaq.order...),
		inters:               append([]Interceptor{}, aaq.inters...),
		predicates:           append([]predicate.AttributeAssertion{}, aaq.predicates...),
		withAttributesReport: aaq.withAttributesReport.Clone(),
		// clone intermediate query.
		sql:  aaq.sql.Clone(),
		path: aaq.path,
	}
}

// WithAttributesReport tells the query-builder to eager-load the nodes that are connected to
// the "attributes_report" edge. The optional arguments are used to configure the query builder of the edge.
func (aaq *AttributeAssertionQuery) WithAttributesReport(opts ...func(*AttributeReportQuery)) *AttributeAssertionQuery {
	query := (&AttributeReportClient{config: aaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aaq.withAttributesReport = query
	return aaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Attribute string `json:"attribute,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttributeAssertion.Query().
//		GroupBy(attributeassertion.FieldAttribute).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aaq *AttributeAssertionQuery) GroupBy(field string, fields ...string) *AttributeAssertionGroupBy {
	aaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeAssertionGroupBy{build: aaq}
	grbuild.flds = &aaq.ctx.Fields
	grbuild.label = attributeassertion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Attribute string `json:"attribute,omitempty"`
//	}
//
//	client.AttributeAssertion.Query().
//		Select(attributeassertion.FieldAttribute).
//		Scan(ctx, &v)
func (aaq *AttributeAssertionQuery) Select(fields ...string) *AttributeAssertionSelect {
	aaq.ctx.Fields = append(aaq.ctx.Fields, fields...)
	sbuild := &AttributeAssertionSelect{AttributeAssertionQuery: aaq}
	sbuild.label = attributeassertion.Label
	sbuild.flds, sbuild.scan = &aaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeAssertionSelect configured with the given aggregations.
func (aaq *AttributeAssertionQuery) Aggregate(fns ...AggregateFunc) *AttributeAssertionSelect {
	return aaq.Select().Aggregate(fns...)
}

func (aaq *AttributeAssertionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aaq); err != nil {
				return err
			}
		}
	}
	for _, f := range aaq.ctx.Fields {
		if !attributeassertion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aaq.path != nil {
		prev, err := aaq.path(ctx)
		if err != nil {
			return err
		}
		aaq.sql = prev
	}
	return nil
}

func (aaq *AttributeAssertionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttributeAssertion, error) {
	var (
		nodes       = []*AttributeAssertion{}
		withFKs     = aaq.withFKs
		_spec       = aaq.querySpec()
		loadedTypes = [1]bool{
			aaq.withAttributesReport != nil,
		}
	)
	if aaq.withAttributesReport != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attributeassertion.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttributeAssertion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttributeAssertion{config: aaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aaq.modifiers) > 0 {
		_spec.Modifiers = aaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aaq.withAttributesReport; query != nil {
		if err := aaq.loadAttributesReport(ctx, query, nodes, nil,
			func(n *AttributeAssertion, e *AttributeReport) { n.Edges.AttributesReport = e }); err != nil {
			return nil, err
		}
	}
	for i := range aaq.loadTotal {
		if err := aaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aaq *AttributeAssertionQuery) loadAttributesReport(ctx context.Context, query *AttributeReportQuery, nodes []*AttributeAssertion, init func(*AttributeAssertion), assign func(*AttributeAssertion, *AttributeReport)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AttributeAssertion)
	for i := range nodes {
		if nodes[i].attribute_report_attributes == nil {
			continue
		}
		fk := *nodes[i].attribute_report_attributes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(attributereport.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "attribute_report_attributes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aaq *AttributeAssertionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aaq.querySpec()
	if len(aaq.modifiers) > 0 {
		_spec.Modifiers = aaq.modifiers
	}
	_spec.Node.Columns = aaq.ctx.Fields
	if len(aaq.ctx.Fields) > 0 {
		_spec.Unique = aaq.ctx.Unique != nil && *aaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aaq.driver, _spec)
}

func (aaq *AttributeAssertionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attributeassertion.Table, attributeassertion.Columns, sqlgraph.NewFieldSpec(attributeassertion.FieldID, field.TypeInt))
	_spec.From = aaq.sql
	if unique := aaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aaq.path != nil {
		_spec.Unique = true
	}
	if fields := aaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributeassertion.FieldID)
		for i := range fields {
			if fields[i] != attributeassertion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aaq *AttributeAssertionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aaq.driver.Dialect())
	t1 := builder.Table(attributeassertion.Table)
	columns := aaq.ctx.Fields
	if len(columns) == 0 {
		columns = attributeassertion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aaq.sql != nil {
		selector = aaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aaq.ctx.Unique != nil && *aaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aaq.predicates {
		p(selector)
	}
	for _, p := range aaq.order {
		p(selector)
	}
	if offset := aaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttributeAssertionGroupBy is the group-by builder for AttributeAssertion entities.
type AttributeAssertionGroupBy struct {
	selector
	build *AttributeAssertionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aagb *AttributeAssertionGroupBy) Aggregate(fns ...AggregateFunc) *AttributeAssertionGroupBy {
	aagb.fns = append(aagb.fns, fns...)
	return aagb
}

// Scan applies the selector query and scans the result into the given value.
func (aagb *AttributeAssertionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aagb.build.ctx, "GroupBy")
	if err := aagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeAssertionQuery, *AttributeAssertionGroupBy](ctx, aagb.build, aagb, aagb.build.inters, v)
}

func (aagb *AttributeAssertionGroupBy) sqlScan(ctx context.Context, root *AttributeAssertionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aagb.fns))
	for _, fn := range aagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aagb.flds)+len(aagb.fns))
		for _, f := range *aagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeAssertionSelect is the builder for selecting fields of AttributeAssertion entities.
type AttributeAssertionSelect struct {
	*AttributeAssertionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aas *AttributeAssertionSelect) Aggregate(fns ...AggregateFunc) *AttributeAssertionSelect {
	aas.fns = append(aas.fns, fns...)
	return aas
}

// Scan applies the selector query and scans the result into the given value.
func (aas *AttributeAssertionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aas.ctx, "Select")
	if err := aas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeAssertionQuery, *AttributeAssertionSelect](ctx, aas.AttributeAssertionQuery, aas, aas.inters, v)
}

func (aas *AttributeAssertionSelect) sqlScan(ctx context.Context, root *AttributeAssertionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aas.fns))
	for _, fn := range aas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}