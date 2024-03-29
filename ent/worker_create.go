// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/fleetcode"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/tractor"
	"github.com/emoss08/trenova/ent/user"
	"github.com/emoss08/trenova/ent/usstate"
	"github.com/emoss08/trenova/ent/worker"
	"github.com/google/uuid"
)

// WorkerCreate is the builder for creating a Worker entity.
type WorkerCreate struct {
	config
	mutation *WorkerMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (wc *WorkerCreate) SetBusinessUnitID(u uuid.UUID) *WorkerCreate {
	wc.mutation.SetBusinessUnitID(u)
	return wc
}

// SetOrganizationID sets the "organization_id" field.
func (wc *WorkerCreate) SetOrganizationID(u uuid.UUID) *WorkerCreate {
	wc.mutation.SetOrganizationID(u)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkerCreate) SetCreatedAt(t time.Time) *WorkerCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableCreatedAt(t *time.Time) *WorkerCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WorkerCreate) SetUpdatedAt(t time.Time) *WorkerCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableUpdatedAt(t *time.Time) *WorkerCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetStatus sets the "status" field.
func (wc *WorkerCreate) SetStatus(w worker.Status) *WorkerCreate {
	wc.mutation.SetStatus(w)
	return wc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableStatus(w *worker.Status) *WorkerCreate {
	if w != nil {
		wc.SetStatus(*w)
	}
	return wc
}

// SetCode sets the "code" field.
func (wc *WorkerCreate) SetCode(s string) *WorkerCreate {
	wc.mutation.SetCode(s)
	return wc
}

// SetProfilePictureURL sets the "profile_picture_url" field.
func (wc *WorkerCreate) SetProfilePictureURL(s string) *WorkerCreate {
	wc.mutation.SetProfilePictureURL(s)
	return wc
}

// SetNillableProfilePictureURL sets the "profile_picture_url" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableProfilePictureURL(s *string) *WorkerCreate {
	if s != nil {
		wc.SetProfilePictureURL(*s)
	}
	return wc
}

// SetWorkerType sets the "worker_type" field.
func (wc *WorkerCreate) SetWorkerType(wt worker.WorkerType) *WorkerCreate {
	wc.mutation.SetWorkerType(wt)
	return wc
}

// SetNillableWorkerType sets the "worker_type" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableWorkerType(wt *worker.WorkerType) *WorkerCreate {
	if wt != nil {
		wc.SetWorkerType(*wt)
	}
	return wc
}

// SetFirstName sets the "first_name" field.
func (wc *WorkerCreate) SetFirstName(s string) *WorkerCreate {
	wc.mutation.SetFirstName(s)
	return wc
}

// SetLastName sets the "last_name" field.
func (wc *WorkerCreate) SetLastName(s string) *WorkerCreate {
	wc.mutation.SetLastName(s)
	return wc
}

// SetCity sets the "city" field.
func (wc *WorkerCreate) SetCity(s string) *WorkerCreate {
	wc.mutation.SetCity(s)
	return wc
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableCity(s *string) *WorkerCreate {
	if s != nil {
		wc.SetCity(*s)
	}
	return wc
}

// SetPostalCode sets the "postal_code" field.
func (wc *WorkerCreate) SetPostalCode(s string) *WorkerCreate {
	wc.mutation.SetPostalCode(s)
	return wc
}

// SetNillablePostalCode sets the "postal_code" field if the given value is not nil.
func (wc *WorkerCreate) SetNillablePostalCode(s *string) *WorkerCreate {
	if s != nil {
		wc.SetPostalCode(*s)
	}
	return wc
}

// SetStateID sets the "state_id" field.
func (wc *WorkerCreate) SetStateID(u uuid.UUID) *WorkerCreate {
	wc.mutation.SetStateID(u)
	return wc
}

// SetNillableStateID sets the "state_id" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableStateID(u *uuid.UUID) *WorkerCreate {
	if u != nil {
		wc.SetStateID(*u)
	}
	return wc
}

// SetFleetCodeID sets the "fleet_code_id" field.
func (wc *WorkerCreate) SetFleetCodeID(u uuid.UUID) *WorkerCreate {
	wc.mutation.SetFleetCodeID(u)
	return wc
}

// SetNillableFleetCodeID sets the "fleet_code_id" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableFleetCodeID(u *uuid.UUID) *WorkerCreate {
	if u != nil {
		wc.SetFleetCodeID(*u)
	}
	return wc
}

// SetManagerID sets the "manager_id" field.
func (wc *WorkerCreate) SetManagerID(u uuid.UUID) *WorkerCreate {
	wc.mutation.SetManagerID(u)
	return wc
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableManagerID(u *uuid.UUID) *WorkerCreate {
	if u != nil {
		wc.SetManagerID(*u)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkerCreate) SetID(u uuid.UUID) *WorkerCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkerCreate) SetNillableID(u *uuid.UUID) *WorkerCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (wc *WorkerCreate) SetBusinessUnit(b *BusinessUnit) *WorkerCreate {
	return wc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (wc *WorkerCreate) SetOrganization(o *Organization) *WorkerCreate {
	return wc.SetOrganizationID(o.ID)
}

// SetState sets the "state" edge to the UsState entity.
func (wc *WorkerCreate) SetState(u *UsState) *WorkerCreate {
	return wc.SetStateID(u.ID)
}

// SetFleetCode sets the "fleet_code" edge to the FleetCode entity.
func (wc *WorkerCreate) SetFleetCode(f *FleetCode) *WorkerCreate {
	return wc.SetFleetCodeID(f.ID)
}

// SetManager sets the "manager" edge to the User entity.
func (wc *WorkerCreate) SetManager(u *User) *WorkerCreate {
	return wc.SetManagerID(u.ID)
}

// AddTractorIDs adds the "tractor" edge to the Tractor entity by IDs.
func (wc *WorkerCreate) AddTractorIDs(ids ...uuid.UUID) *WorkerCreate {
	wc.mutation.AddTractorIDs(ids...)
	return wc
}

// AddTractor adds the "tractor" edges to the Tractor entity.
func (wc *WorkerCreate) AddTractor(t ...*Tractor) *WorkerCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTractorIDs(ids...)
}

// AddSecondaryTractorIDs adds the "secondary_tractor" edge to the Tractor entity by IDs.
func (wc *WorkerCreate) AddSecondaryTractorIDs(ids ...uuid.UUID) *WorkerCreate {
	wc.mutation.AddSecondaryTractorIDs(ids...)
	return wc
}

// AddSecondaryTractor adds the "secondary_tractor" edges to the Tractor entity.
func (wc *WorkerCreate) AddSecondaryTractor(t ...*Tractor) *WorkerCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddSecondaryTractorIDs(ids...)
}

// Mutation returns the WorkerMutation object of the builder.
func (wc *WorkerCreate) Mutation() *WorkerMutation {
	return wc.mutation
}

// Save creates the Worker in the database.
func (wc *WorkerCreate) Save(ctx context.Context) (*Worker, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkerCreate) SaveX(ctx context.Context) *Worker {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkerCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkerCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkerCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := worker.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := worker.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.Status(); !ok {
		v := worker.DefaultStatus
		wc.mutation.SetStatus(v)
	}
	if _, ok := wc.mutation.WorkerType(); !ok {
		v := worker.DefaultWorkerType
		wc.mutation.SetWorkerType(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := worker.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkerCreate) check() error {
	if _, ok := wc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "Worker.business_unit_id"`)}
	}
	if _, ok := wc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "Worker.organization_id"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Worker.created_at"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Worker.updated_at"`)}
	}
	if _, ok := wc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Worker.status"`)}
	}
	if v, ok := wc.mutation.Status(); ok {
		if err := worker.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Worker.status": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Worker.code"`)}
	}
	if v, ok := wc.mutation.Code(); ok {
		if err := worker.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Worker.code": %w`, err)}
		}
	}
	if _, ok := wc.mutation.WorkerType(); !ok {
		return &ValidationError{Name: "worker_type", err: errors.New(`ent: missing required field "Worker.worker_type"`)}
	}
	if v, ok := wc.mutation.WorkerType(); ok {
		if err := worker.WorkerTypeValidator(v); err != nil {
			return &ValidationError{Name: "worker_type", err: fmt.Errorf(`ent: validator failed for field "Worker.worker_type": %w`, err)}
		}
	}
	if _, ok := wc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Worker.first_name"`)}
	}
	if v, ok := wc.mutation.FirstName(); ok {
		if err := worker.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Worker.first_name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Worker.last_name"`)}
	}
	if v, ok := wc.mutation.LastName(); ok {
		if err := worker.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Worker.last_name": %w`, err)}
		}
	}
	if v, ok := wc.mutation.City(); ok {
		if err := worker.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "Worker.city": %w`, err)}
		}
	}
	if v, ok := wc.mutation.PostalCode(); ok {
		if err := worker.PostalCodeValidator(v); err != nil {
			return &ValidationError{Name: "postal_code", err: fmt.Errorf(`ent: validator failed for field "Worker.postal_code": %w`, err)}
		}
	}
	if _, ok := wc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "Worker.business_unit"`)}
	}
	if _, ok := wc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "Worker.organization"`)}
	}
	return nil
}

func (wc *WorkerCreate) sqlSave(ctx context.Context) (*Worker, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkerCreate) createSpec() (*Worker, *sqlgraph.CreateSpec) {
	var (
		_node = &Worker{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(worker.Table, sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(worker.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.SetField(worker.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wc.mutation.Status(); ok {
		_spec.SetField(worker.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := wc.mutation.Code(); ok {
		_spec.SetField(worker.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := wc.mutation.ProfilePictureURL(); ok {
		_spec.SetField(worker.FieldProfilePictureURL, field.TypeString, value)
		_node.ProfilePictureURL = value
	}
	if value, ok := wc.mutation.WorkerType(); ok {
		_spec.SetField(worker.FieldWorkerType, field.TypeEnum, value)
		_node.WorkerType = value
	}
	if value, ok := wc.mutation.FirstName(); ok {
		_spec.SetField(worker.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := wc.mutation.LastName(); ok {
		_spec.SetField(worker.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := wc.mutation.City(); ok {
		_spec.SetField(worker.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := wc.mutation.PostalCode(); ok {
		_spec.SetField(worker.FieldPostalCode, field.TypeString, value)
		_node.PostalCode = value
	}
	if nodes := wc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   worker.BusinessUnitTable,
			Columns: []string{worker.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BusinessUnitID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   worker.OrganizationTable,
			Columns: []string{worker.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.StateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   worker.StateTable,
			Columns: []string{worker.StateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usstate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StateID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.FleetCodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   worker.FleetCodeTable,
			Columns: []string{worker.FleetCodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fleetcode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FleetCodeID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ManagerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   worker.ManagerTable,
			Columns: []string{worker.ManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ManagerID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TractorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   worker.TractorTable,
			Columns: []string{worker.TractorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tractor.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.SecondaryTractorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   worker.SecondaryTractorTable,
			Columns: []string{worker.SecondaryTractorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tractor.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkerCreateBulk is the builder for creating many Worker entities in bulk.
type WorkerCreateBulk struct {
	config
	err      error
	builders []*WorkerCreate
}

// Save creates the Worker entities in the database.
func (wcb *WorkerCreateBulk) Save(ctx context.Context) ([]*Worker, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Worker, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkerMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkerCreateBulk) SaveX(ctx context.Context) []*Worker {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkerCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkerCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
