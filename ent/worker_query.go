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
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/fleetcode"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/tractor"
	"github.com/emoss08/trenova/ent/user"
	"github.com/emoss08/trenova/ent/usstate"
	"github.com/emoss08/trenova/ent/worker"
	"github.com/emoss08/trenova/ent/workerprofile"
	"github.com/google/uuid"
)

// WorkerQuery is the builder for querying Worker entities.
type WorkerQuery struct {
	config
	ctx                  *QueryContext
	order                []worker.OrderOption
	inters               []Interceptor
	predicates           []predicate.Worker
	withBusinessUnit     *BusinessUnitQuery
	withOrganization     *OrganizationQuery
	withState            *UsStateQuery
	withFleetCode        *FleetCodeQuery
	withManager          *UserQuery
	withPrimaryTractor   *TractorQuery
	withSecondaryTractor *TractorQuery
	withWorkerProfile    *WorkerProfileQuery
	modifiers            []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkerQuery builder.
func (wq *WorkerQuery) Where(ps ...predicate.Worker) *WorkerQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WorkerQuery) Limit(limit int) *WorkerQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WorkerQuery) Offset(offset int) *WorkerQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WorkerQuery) Unique(unique bool) *WorkerQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WorkerQuery) Order(o ...worker.OrderOption) *WorkerQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (wq *WorkerQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, worker.BusinessUnitTable, worker.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (wq *WorkerQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, worker.OrganizationTable, worker.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryState chains the current query on the "state" edge.
func (wq *WorkerQuery) QueryState() *UsStateQuery {
	query := (&UsStateClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(usstate.Table, usstate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, worker.StateTable, worker.StateColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFleetCode chains the current query on the "fleet_code" edge.
func (wq *WorkerQuery) QueryFleetCode() *FleetCodeQuery {
	query := (&FleetCodeClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(fleetcode.Table, fleetcode.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, worker.FleetCodeTable, worker.FleetCodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryManager chains the current query on the "manager" edge.
func (wq *WorkerQuery) QueryManager() *UserQuery {
	query := (&UserClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, worker.ManagerTable, worker.ManagerColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPrimaryTractor chains the current query on the "primary_tractor" edge.
func (wq *WorkerQuery) QueryPrimaryTractor() *TractorQuery {
	query := (&TractorClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(tractor.Table, tractor.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, worker.PrimaryTractorTable, worker.PrimaryTractorColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySecondaryTractor chains the current query on the "secondary_tractor" edge.
func (wq *WorkerQuery) QuerySecondaryTractor() *TractorQuery {
	query := (&TractorClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(tractor.Table, tractor.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, worker.SecondaryTractorTable, worker.SecondaryTractorColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkerProfile chains the current query on the "worker_profile" edge.
func (wq *WorkerQuery) QueryWorkerProfile() *WorkerProfileQuery {
	query := (&WorkerProfileClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(worker.Table, worker.FieldID, selector),
			sqlgraph.To(workerprofile.Table, workerprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, worker.WorkerProfileTable, worker.WorkerProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Worker entity from the query.
// Returns a *NotFoundError when no Worker was found.
func (wq *WorkerQuery) First(ctx context.Context) (*Worker, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{worker.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WorkerQuery) FirstX(ctx context.Context) *Worker {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Worker ID from the query.
// Returns a *NotFoundError when no Worker ID was found.
func (wq *WorkerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{worker.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WorkerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Worker entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Worker entity is found.
// Returns a *NotFoundError when no Worker entities are found.
func (wq *WorkerQuery) Only(ctx context.Context) (*Worker, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{worker.Label}
	default:
		return nil, &NotSingularError{worker.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WorkerQuery) OnlyX(ctx context.Context) *Worker {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Worker ID in the query.
// Returns a *NotSingularError when more than one Worker ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WorkerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{worker.Label}
	default:
		err = &NotSingularError{worker.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WorkerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Workers.
func (wq *WorkerQuery) All(ctx context.Context) ([]*Worker, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Worker, *WorkerQuery]()
	return withInterceptors[[]*Worker](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WorkerQuery) AllX(ctx context.Context) []*Worker {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Worker IDs.
func (wq *WorkerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(worker.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WorkerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WorkerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WorkerQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WorkerQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WorkerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WorkerQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WorkerQuery) Clone() *WorkerQuery {
	if wq == nil {
		return nil
	}
	return &WorkerQuery{
		config:               wq.config,
		ctx:                  wq.ctx.Clone(),
		order:                append([]worker.OrderOption{}, wq.order...),
		inters:               append([]Interceptor{}, wq.inters...),
		predicates:           append([]predicate.Worker{}, wq.predicates...),
		withBusinessUnit:     wq.withBusinessUnit.Clone(),
		withOrganization:     wq.withOrganization.Clone(),
		withState:            wq.withState.Clone(),
		withFleetCode:        wq.withFleetCode.Clone(),
		withManager:          wq.withManager.Clone(),
		withPrimaryTractor:   wq.withPrimaryTractor.Clone(),
		withSecondaryTractor: wq.withSecondaryTractor.Clone(),
		withWorkerProfile:    wq.withWorkerProfile.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *WorkerQuery {
	query := (&BusinessUnitClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withBusinessUnit = query
	return wq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithOrganization(opts ...func(*OrganizationQuery)) *WorkerQuery {
	query := (&OrganizationClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withOrganization = query
	return wq
}

// WithState tells the query-builder to eager-load the nodes that are connected to
// the "state" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithState(opts ...func(*UsStateQuery)) *WorkerQuery {
	query := (&UsStateClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withState = query
	return wq
}

// WithFleetCode tells the query-builder to eager-load the nodes that are connected to
// the "fleet_code" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithFleetCode(opts ...func(*FleetCodeQuery)) *WorkerQuery {
	query := (&FleetCodeClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withFleetCode = query
	return wq
}

// WithManager tells the query-builder to eager-load the nodes that are connected to
// the "manager" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithManager(opts ...func(*UserQuery)) *WorkerQuery {
	query := (&UserClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withManager = query
	return wq
}

// WithPrimaryTractor tells the query-builder to eager-load the nodes that are connected to
// the "primary_tractor" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithPrimaryTractor(opts ...func(*TractorQuery)) *WorkerQuery {
	query := (&TractorClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withPrimaryTractor = query
	return wq
}

// WithSecondaryTractor tells the query-builder to eager-load the nodes that are connected to
// the "secondary_tractor" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithSecondaryTractor(opts ...func(*TractorQuery)) *WorkerQuery {
	query := (&TractorClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withSecondaryTractor = query
	return wq
}

// WithWorkerProfile tells the query-builder to eager-load the nodes that are connected to
// the "worker_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkerQuery) WithWorkerProfile(opts ...func(*WorkerProfileQuery)) *WorkerQuery {
	query := (&WorkerProfileClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkerProfile = query
	return wq
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
//	client.Worker.Query().
//		GroupBy(worker.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WorkerQuery) GroupBy(field string, fields ...string) *WorkerGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkerGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = worker.Label
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
//	client.Worker.Query().
//		Select(worker.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (wq *WorkerQuery) Select(fields ...string) *WorkerSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WorkerSelect{WorkerQuery: wq}
	sbuild.label = worker.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkerSelect configured with the given aggregations.
func (wq *WorkerQuery) Aggregate(fns ...AggregateFunc) *WorkerSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WorkerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !worker.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WorkerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Worker, error) {
	var (
		nodes       = []*Worker{}
		_spec       = wq.querySpec()
		loadedTypes = [8]bool{
			wq.withBusinessUnit != nil,
			wq.withOrganization != nil,
			wq.withState != nil,
			wq.withFleetCode != nil,
			wq.withManager != nil,
			wq.withPrimaryTractor != nil,
			wq.withSecondaryTractor != nil,
			wq.withWorkerProfile != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Worker).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Worker{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withBusinessUnit; query != nil {
		if err := wq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *Worker, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withOrganization; query != nil {
		if err := wq.loadOrganization(ctx, query, nodes, nil,
			func(n *Worker, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withState; query != nil {
		if err := wq.loadState(ctx, query, nodes, nil,
			func(n *Worker, e *UsState) { n.Edges.State = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withFleetCode; query != nil {
		if err := wq.loadFleetCode(ctx, query, nodes, nil,
			func(n *Worker, e *FleetCode) { n.Edges.FleetCode = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withManager; query != nil {
		if err := wq.loadManager(ctx, query, nodes, nil,
			func(n *Worker, e *User) { n.Edges.Manager = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withPrimaryTractor; query != nil {
		if err := wq.loadPrimaryTractor(ctx, query, nodes, nil,
			func(n *Worker, e *Tractor) { n.Edges.PrimaryTractor = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withSecondaryTractor; query != nil {
		if err := wq.loadSecondaryTractor(ctx, query, nodes, nil,
			func(n *Worker, e *Tractor) { n.Edges.SecondaryTractor = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWorkerProfile; query != nil {
		if err := wq.loadWorkerProfile(ctx, query, nodes, nil,
			func(n *Worker, e *WorkerProfile) { n.Edges.WorkerProfile = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WorkerQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Worker)
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
func (wq *WorkerQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Worker)
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
func (wq *WorkerQuery) loadState(ctx context.Context, query *UsStateQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *UsState)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Worker)
	for i := range nodes {
		if nodes[i].StateID == nil {
			continue
		}
		fk := *nodes[i].StateID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(usstate.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "state_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkerQuery) loadFleetCode(ctx context.Context, query *FleetCodeQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *FleetCode)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Worker)
	for i := range nodes {
		if nodes[i].FleetCodeID == nil {
			continue
		}
		fk := *nodes[i].FleetCodeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(fleetcode.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "fleet_code_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkerQuery) loadManager(ctx context.Context, query *UserQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Worker)
	for i := range nodes {
		if nodes[i].ManagerID == nil {
			continue
		}
		fk := *nodes[i].ManagerID
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
			return fmt.Errorf(`unexpected foreign-key "manager_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkerQuery) loadPrimaryTractor(ctx context.Context, query *TractorQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *Tractor)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Worker)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tractor.FieldPrimaryWorkerID)
	}
	query.Where(predicate.Tractor(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(worker.PrimaryTractorColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PrimaryWorkerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "primary_worker_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WorkerQuery) loadSecondaryTractor(ctx context.Context, query *TractorQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *Tractor)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Worker)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tractor.FieldSecondaryWorkerID)
	}
	query.Where(predicate.Tractor(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(worker.SecondaryTractorColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SecondaryWorkerID
		if fk == nil {
			return fmt.Errorf(`foreign-key "secondary_worker_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "secondary_worker_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WorkerQuery) loadWorkerProfile(ctx context.Context, query *WorkerProfileQuery, nodes []*Worker, init func(*Worker), assign func(*Worker, *WorkerProfile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Worker)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(workerprofile.FieldWorkerID)
	}
	query.Where(predicate.WorkerProfile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(worker.WorkerProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.WorkerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "worker_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WorkerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WorkerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(worker.Table, worker.Columns, sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, worker.FieldID)
		for i := range fields {
			if fields[i] != worker.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(worker.FieldBusinessUnitID)
		}
		if wq.withOrganization != nil {
			_spec.Node.AddColumnOnce(worker.FieldOrganizationID)
		}
		if wq.withState != nil {
			_spec.Node.AddColumnOnce(worker.FieldStateID)
		}
		if wq.withFleetCode != nil {
			_spec.Node.AddColumnOnce(worker.FieldFleetCodeID)
		}
		if wq.withManager != nil {
			_spec.Node.AddColumnOnce(worker.FieldManagerID)
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WorkerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(worker.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = worker.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wq.modifiers {
		m(selector)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wq *WorkerQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkerSelect {
	wq.modifiers = append(wq.modifiers, modifiers...)
	return wq.Select()
}

// WorkerGroupBy is the group-by builder for Worker entities.
type WorkerGroupBy struct {
	selector
	build *WorkerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WorkerGroupBy) Aggregate(fns ...AggregateFunc) *WorkerGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WorkerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkerQuery, *WorkerGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WorkerGroupBy) sqlScan(ctx context.Context, root *WorkerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkerSelect is the builder for selecting fields of Worker entities.
type WorkerSelect struct {
	*WorkerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WorkerSelect) Aggregate(fns ...AggregateFunc) *WorkerSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WorkerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkerQuery, *WorkerSelect](ctx, ws.WorkerQuery, ws, ws.inters, v)
}

func (ws *WorkerSelect) sqlScan(ctx context.Context, root *WorkerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ws *WorkerSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkerSelect {
	ws.modifiers = append(ws.modifiers, modifiers...)
	return ws
}
