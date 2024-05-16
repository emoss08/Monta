// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/customer"
	"github.com/emoss08/trenova/internal/ent/customercontact"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// CustomerContactQuery is the builder for querying CustomerContact entities.
type CustomerContactQuery struct {
	config
	ctx              *QueryContext
	order            []customercontact.OrderOption
	inters           []Interceptor
	predicates       []predicate.CustomerContact
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withCustomer     *CustomerQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerContactQuery builder.
func (ccq *CustomerContactQuery) Where(ps ...predicate.CustomerContact) *CustomerContactQuery {
	ccq.predicates = append(ccq.predicates, ps...)
	return ccq
}

// Limit the number of records to be returned by this query.
func (ccq *CustomerContactQuery) Limit(limit int) *CustomerContactQuery {
	ccq.ctx.Limit = &limit
	return ccq
}

// Offset to start from.
func (ccq *CustomerContactQuery) Offset(offset int) *CustomerContactQuery {
	ccq.ctx.Offset = &offset
	return ccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccq *CustomerContactQuery) Unique(unique bool) *CustomerContactQuery {
	ccq.ctx.Unique = &unique
	return ccq
}

// Order specifies how the records should be ordered.
func (ccq *CustomerContactQuery) Order(o ...customercontact.OrderOption) *CustomerContactQuery {
	ccq.order = append(ccq.order, o...)
	return ccq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (ccq *CustomerContactQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: ccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customercontact.Table, customercontact.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, customercontact.BusinessUnitTable, customercontact.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(ccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (ccq *CustomerContactQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: ccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customercontact.Table, customercontact.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, customercontact.OrganizationTable, customercontact.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(ccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomer chains the current query on the "customer" edge.
func (ccq *CustomerContactQuery) QueryCustomer() *CustomerQuery {
	query := (&CustomerClient{config: ccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customercontact.Table, customercontact.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, customercontact.CustomerTable, customercontact.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(ccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CustomerContact entity from the query.
// Returns a *NotFoundError when no CustomerContact was found.
func (ccq *CustomerContactQuery) First(ctx context.Context) (*CustomerContact, error) {
	nodes, err := ccq.Limit(1).All(setContextOp(ctx, ccq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customercontact.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccq *CustomerContactQuery) FirstX(ctx context.Context) *CustomerContact {
	node, err := ccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CustomerContact ID from the query.
// Returns a *NotFoundError when no CustomerContact ID was found.
func (ccq *CustomerContactQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ccq.Limit(1).IDs(setContextOp(ctx, ccq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customercontact.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccq *CustomerContactQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CustomerContact entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CustomerContact entity is found.
// Returns a *NotFoundError when no CustomerContact entities are found.
func (ccq *CustomerContactQuery) Only(ctx context.Context) (*CustomerContact, error) {
	nodes, err := ccq.Limit(2).All(setContextOp(ctx, ccq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customercontact.Label}
	default:
		return nil, &NotSingularError{customercontact.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccq *CustomerContactQuery) OnlyX(ctx context.Context) *CustomerContact {
	node, err := ccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CustomerContact ID in the query.
// Returns a *NotSingularError when more than one CustomerContact ID is found.
// Returns a *NotFoundError when no entities are found.
func (ccq *CustomerContactQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ccq.Limit(2).IDs(setContextOp(ctx, ccq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customercontact.Label}
	default:
		err = &NotSingularError{customercontact.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccq *CustomerContactQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CustomerContacts.
func (ccq *CustomerContactQuery) All(ctx context.Context) ([]*CustomerContact, error) {
	ctx = setContextOp(ctx, ccq.ctx, "All")
	if err := ccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CustomerContact, *CustomerContactQuery]()
	return withInterceptors[[]*CustomerContact](ctx, ccq, qr, ccq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ccq *CustomerContactQuery) AllX(ctx context.Context) []*CustomerContact {
	nodes, err := ccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CustomerContact IDs.
func (ccq *CustomerContactQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ccq.ctx.Unique == nil && ccq.path != nil {
		ccq.Unique(true)
	}
	ctx = setContextOp(ctx, ccq.ctx, "IDs")
	if err = ccq.Select(customercontact.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccq *CustomerContactQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccq *CustomerContactQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ccq.ctx, "Count")
	if err := ccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ccq, querierCount[*CustomerContactQuery](), ccq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ccq *CustomerContactQuery) CountX(ctx context.Context) int {
	count, err := ccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccq *CustomerContactQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ccq.ctx, "Exist")
	switch _, err := ccq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ccq *CustomerContactQuery) ExistX(ctx context.Context) bool {
	exist, err := ccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerContactQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccq *CustomerContactQuery) Clone() *CustomerContactQuery {
	if ccq == nil {
		return nil
	}
	return &CustomerContactQuery{
		config:           ccq.config,
		ctx:              ccq.ctx.Clone(),
		order:            append([]customercontact.OrderOption{}, ccq.order...),
		inters:           append([]Interceptor{}, ccq.inters...),
		predicates:       append([]predicate.CustomerContact{}, ccq.predicates...),
		withBusinessUnit: ccq.withBusinessUnit.Clone(),
		withOrganization: ccq.withOrganization.Clone(),
		withCustomer:     ccq.withCustomer.Clone(),
		// clone intermediate query.
		sql:  ccq.sql.Clone(),
		path: ccq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (ccq *CustomerContactQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *CustomerContactQuery {
	query := (&BusinessUnitClient{config: ccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ccq.withBusinessUnit = query
	return ccq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (ccq *CustomerContactQuery) WithOrganization(opts ...func(*OrganizationQuery)) *CustomerContactQuery {
	query := (&OrganizationClient{config: ccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ccq.withOrganization = query
	return ccq
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (ccq *CustomerContactQuery) WithCustomer(opts ...func(*CustomerQuery)) *CustomerContactQuery {
	query := (&CustomerClient{config: ccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ccq.withCustomer = query
	return ccq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CustomerContact.Query().
//		GroupBy(customercontact.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ccq *CustomerContactQuery) GroupBy(field string, fields ...string) *CustomerContactGroupBy {
	ccq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CustomerContactGroupBy{build: ccq}
	grbuild.flds = &ccq.ctx.Fields
	grbuild.label = customercontact.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//	}
//
//	client.CustomerContact.Query().
//		Select(customercontact.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (ccq *CustomerContactQuery) Select(fields ...string) *CustomerContactSelect {
	ccq.ctx.Fields = append(ccq.ctx.Fields, fields...)
	sbuild := &CustomerContactSelect{CustomerContactQuery: ccq}
	sbuild.label = customercontact.Label
	sbuild.flds, sbuild.scan = &ccq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomerContactSelect configured with the given aggregations.
func (ccq *CustomerContactQuery) Aggregate(fns ...AggregateFunc) *CustomerContactSelect {
	return ccq.Select().Aggregate(fns...)
}

func (ccq *CustomerContactQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ccq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ccq); err != nil {
				return err
			}
		}
	}
	for _, f := range ccq.ctx.Fields {
		if !customercontact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccq.path != nil {
		prev, err := ccq.path(ctx)
		if err != nil {
			return err
		}
		ccq.sql = prev
	}
	return nil
}

func (ccq *CustomerContactQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CustomerContact, error) {
	var (
		nodes       = []*CustomerContact{}
		_spec       = ccq.querySpec()
		loadedTypes = [3]bool{
			ccq.withBusinessUnit != nil,
			ccq.withOrganization != nil,
			ccq.withCustomer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CustomerContact).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CustomerContact{config: ccq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ccq.modifiers) > 0 {
		_spec.Modifiers = ccq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ccq.withBusinessUnit; query != nil {
		if err := ccq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *CustomerContact, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := ccq.withOrganization; query != nil {
		if err := ccq.loadOrganization(ctx, query, nodes, nil,
			func(n *CustomerContact, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := ccq.withCustomer; query != nil {
		if err := ccq.loadCustomer(ctx, query, nodes, nil,
			func(n *CustomerContact, e *Customer) { n.Edges.Customer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ccq *CustomerContactQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*CustomerContact, init func(*CustomerContact), assign func(*CustomerContact, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CustomerContact)
	for i := range nodes {
		fk := nodes[i].BusinessUnitID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(businessunit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_unit_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ccq *CustomerContactQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*CustomerContact, init func(*CustomerContact), assign func(*CustomerContact, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CustomerContact)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ccq *CustomerContactQuery) loadCustomer(ctx context.Context, query *CustomerQuery, nodes []*CustomerContact, init func(*CustomerContact), assign func(*CustomerContact, *Customer)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CustomerContact)
	for i := range nodes {
		fk := nodes[i].CustomerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(customer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "customer_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ccq *CustomerContactQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccq.querySpec()
	if len(ccq.modifiers) > 0 {
		_spec.Modifiers = ccq.modifiers
	}
	_spec.Node.Columns = ccq.ctx.Fields
	if len(ccq.ctx.Fields) > 0 {
		_spec.Unique = ccq.ctx.Unique != nil && *ccq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ccq.driver, _spec)
}

func (ccq *CustomerContactQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(customercontact.Table, customercontact.Columns, sqlgraph.NewFieldSpec(customercontact.FieldID, field.TypeUUID))
	_spec.From = ccq.sql
	if unique := ccq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ccq.path != nil {
		_spec.Unique = true
	}
	if fields := ccq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customercontact.FieldID)
		for i := range fields {
			if fields[i] != customercontact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ccq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(customercontact.FieldBusinessUnitID)
		}
		if ccq.withOrganization != nil {
			_spec.Node.AddColumnOnce(customercontact.FieldOrganizationID)
		}
		if ccq.withCustomer != nil {
			_spec.Node.AddColumnOnce(customercontact.FieldCustomerID)
		}
	}
	if ps := ccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccq *CustomerContactQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccq.driver.Dialect())
	t1 := builder.Table(customercontact.Table)
	columns := ccq.ctx.Fields
	if len(columns) == 0 {
		columns = customercontact.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccq.sql != nil {
		selector = ccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccq.ctx.Unique != nil && *ccq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ccq.modifiers {
		m(selector)
	}
	for _, p := range ccq.predicates {
		p(selector)
	}
	for _, p := range ccq.order {
		p(selector)
	}
	if offset := ccq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ccq *CustomerContactQuery) Modify(modifiers ...func(s *sql.Selector)) *CustomerContactSelect {
	ccq.modifiers = append(ccq.modifiers, modifiers...)
	return ccq.Select()
}

// CustomerContactGroupBy is the group-by builder for CustomerContact entities.
type CustomerContactGroupBy struct {
	selector
	build *CustomerContactQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccgb *CustomerContactGroupBy) Aggregate(fns ...AggregateFunc) *CustomerContactGroupBy {
	ccgb.fns = append(ccgb.fns, fns...)
	return ccgb
}

// Scan applies the selector query and scans the result into the given value.
func (ccgb *CustomerContactGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccgb.build.ctx, "GroupBy")
	if err := ccgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerContactQuery, *CustomerContactGroupBy](ctx, ccgb.build, ccgb, ccgb.build.inters, v)
}

func (ccgb *CustomerContactGroupBy) sqlScan(ctx context.Context, root *CustomerContactQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ccgb.fns))
	for _, fn := range ccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ccgb.flds)+len(ccgb.fns))
		for _, f := range *ccgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ccgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomerContactSelect is the builder for selecting fields of CustomerContact entities.
type CustomerContactSelect struct {
	*CustomerContactQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ccs *CustomerContactSelect) Aggregate(fns ...AggregateFunc) *CustomerContactSelect {
	ccs.fns = append(ccs.fns, fns...)
	return ccs
}

// Scan applies the selector query and scans the result into the given value.
func (ccs *CustomerContactSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccs.ctx, "Select")
	if err := ccs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerContactQuery, *CustomerContactSelect](ctx, ccs.CustomerContactQuery, ccs, ccs.inters, v)
}

func (ccs *CustomerContactSelect) sqlScan(ctx context.Context, root *CustomerContactQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ccs.fns))
	for _, fn := range ccs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ccs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ccs *CustomerContactSelect) Modify(modifiers ...func(s *sql.Selector)) *CustomerContactSelect {
	ccs.modifiers = append(ccs.modifiers, modifiers...)
	return ccs
}