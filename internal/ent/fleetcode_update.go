// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/fleetcode"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/user"
	"github.com/google/uuid"
)

// FleetCodeUpdate is the builder for updating FleetCode entities.
type FleetCodeUpdate struct {
	config
	hooks     []Hook
	mutation  *FleetCodeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FleetCodeUpdate builder.
func (fcu *FleetCodeUpdate) Where(ps ...predicate.FleetCode) *FleetCodeUpdate {
	fcu.mutation.Where(ps...)
	return fcu
}

// SetOrganizationID sets the "organization_id" field.
func (fcu *FleetCodeUpdate) SetOrganizationID(u uuid.UUID) *FleetCodeUpdate {
	fcu.mutation.SetOrganizationID(u)
	return fcu
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableOrganizationID(u *uuid.UUID) *FleetCodeUpdate {
	if u != nil {
		fcu.SetOrganizationID(*u)
	}
	return fcu
}

// SetUpdatedAt sets the "updated_at" field.
func (fcu *FleetCodeUpdate) SetUpdatedAt(t time.Time) *FleetCodeUpdate {
	fcu.mutation.SetUpdatedAt(t)
	return fcu
}

// SetVersion sets the "version" field.
func (fcu *FleetCodeUpdate) SetVersion(i int) *FleetCodeUpdate {
	fcu.mutation.ResetVersion()
	fcu.mutation.SetVersion(i)
	return fcu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableVersion(i *int) *FleetCodeUpdate {
	if i != nil {
		fcu.SetVersion(*i)
	}
	return fcu
}

// AddVersion adds i to the "version" field.
func (fcu *FleetCodeUpdate) AddVersion(i int) *FleetCodeUpdate {
	fcu.mutation.AddVersion(i)
	return fcu
}

// SetStatus sets the "status" field.
func (fcu *FleetCodeUpdate) SetStatus(f fleetcode.Status) *FleetCodeUpdate {
	fcu.mutation.SetStatus(f)
	return fcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableStatus(f *fleetcode.Status) *FleetCodeUpdate {
	if f != nil {
		fcu.SetStatus(*f)
	}
	return fcu
}

// SetCode sets the "code" field.
func (fcu *FleetCodeUpdate) SetCode(s string) *FleetCodeUpdate {
	fcu.mutation.SetCode(s)
	return fcu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableCode(s *string) *FleetCodeUpdate {
	if s != nil {
		fcu.SetCode(*s)
	}
	return fcu
}

// SetDescription sets the "description" field.
func (fcu *FleetCodeUpdate) SetDescription(s string) *FleetCodeUpdate {
	fcu.mutation.SetDescription(s)
	return fcu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableDescription(s *string) *FleetCodeUpdate {
	if s != nil {
		fcu.SetDescription(*s)
	}
	return fcu
}

// ClearDescription clears the value of the "description" field.
func (fcu *FleetCodeUpdate) ClearDescription() *FleetCodeUpdate {
	fcu.mutation.ClearDescription()
	return fcu
}

// SetRevenueGoal sets the "revenue_goal" field.
func (fcu *FleetCodeUpdate) SetRevenueGoal(f float64) *FleetCodeUpdate {
	fcu.mutation.ResetRevenueGoal()
	fcu.mutation.SetRevenueGoal(f)
	return fcu
}

// SetNillableRevenueGoal sets the "revenue_goal" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableRevenueGoal(f *float64) *FleetCodeUpdate {
	if f != nil {
		fcu.SetRevenueGoal(*f)
	}
	return fcu
}

// AddRevenueGoal adds f to the "revenue_goal" field.
func (fcu *FleetCodeUpdate) AddRevenueGoal(f float64) *FleetCodeUpdate {
	fcu.mutation.AddRevenueGoal(f)
	return fcu
}

// ClearRevenueGoal clears the value of the "revenue_goal" field.
func (fcu *FleetCodeUpdate) ClearRevenueGoal() *FleetCodeUpdate {
	fcu.mutation.ClearRevenueGoal()
	return fcu
}

// SetDeadheadGoal sets the "deadhead_goal" field.
func (fcu *FleetCodeUpdate) SetDeadheadGoal(f float64) *FleetCodeUpdate {
	fcu.mutation.ResetDeadheadGoal()
	fcu.mutation.SetDeadheadGoal(f)
	return fcu
}

// SetNillableDeadheadGoal sets the "deadhead_goal" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableDeadheadGoal(f *float64) *FleetCodeUpdate {
	if f != nil {
		fcu.SetDeadheadGoal(*f)
	}
	return fcu
}

// AddDeadheadGoal adds f to the "deadhead_goal" field.
func (fcu *FleetCodeUpdate) AddDeadheadGoal(f float64) *FleetCodeUpdate {
	fcu.mutation.AddDeadheadGoal(f)
	return fcu
}

// ClearDeadheadGoal clears the value of the "deadhead_goal" field.
func (fcu *FleetCodeUpdate) ClearDeadheadGoal() *FleetCodeUpdate {
	fcu.mutation.ClearDeadheadGoal()
	return fcu
}

// SetMileageGoal sets the "mileage_goal" field.
func (fcu *FleetCodeUpdate) SetMileageGoal(f float64) *FleetCodeUpdate {
	fcu.mutation.ResetMileageGoal()
	fcu.mutation.SetMileageGoal(f)
	return fcu
}

// SetNillableMileageGoal sets the "mileage_goal" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableMileageGoal(f *float64) *FleetCodeUpdate {
	if f != nil {
		fcu.SetMileageGoal(*f)
	}
	return fcu
}

// AddMileageGoal adds f to the "mileage_goal" field.
func (fcu *FleetCodeUpdate) AddMileageGoal(f float64) *FleetCodeUpdate {
	fcu.mutation.AddMileageGoal(f)
	return fcu
}

// ClearMileageGoal clears the value of the "mileage_goal" field.
func (fcu *FleetCodeUpdate) ClearMileageGoal() *FleetCodeUpdate {
	fcu.mutation.ClearMileageGoal()
	return fcu
}

// SetManagerID sets the "manager_id" field.
func (fcu *FleetCodeUpdate) SetManagerID(u uuid.UUID) *FleetCodeUpdate {
	fcu.mutation.SetManagerID(u)
	return fcu
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableManagerID(u *uuid.UUID) *FleetCodeUpdate {
	if u != nil {
		fcu.SetManagerID(*u)
	}
	return fcu
}

// ClearManagerID clears the value of the "manager_id" field.
func (fcu *FleetCodeUpdate) ClearManagerID() *FleetCodeUpdate {
	fcu.mutation.ClearManagerID()
	return fcu
}

// SetColor sets the "color" field.
func (fcu *FleetCodeUpdate) SetColor(s string) *FleetCodeUpdate {
	fcu.mutation.SetColor(s)
	return fcu
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (fcu *FleetCodeUpdate) SetNillableColor(s *string) *FleetCodeUpdate {
	if s != nil {
		fcu.SetColor(*s)
	}
	return fcu
}

// ClearColor clears the value of the "color" field.
func (fcu *FleetCodeUpdate) ClearColor() *FleetCodeUpdate {
	fcu.mutation.ClearColor()
	return fcu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (fcu *FleetCodeUpdate) SetOrganization(o *Organization) *FleetCodeUpdate {
	return fcu.SetOrganizationID(o.ID)
}

// SetManager sets the "manager" edge to the User entity.
func (fcu *FleetCodeUpdate) SetManager(u *User) *FleetCodeUpdate {
	return fcu.SetManagerID(u.ID)
}

// Mutation returns the FleetCodeMutation object of the builder.
func (fcu *FleetCodeUpdate) Mutation() *FleetCodeMutation {
	return fcu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (fcu *FleetCodeUpdate) ClearOrganization() *FleetCodeUpdate {
	fcu.mutation.ClearOrganization()
	return fcu
}

// ClearManager clears the "manager" edge to the User entity.
func (fcu *FleetCodeUpdate) ClearManager() *FleetCodeUpdate {
	fcu.mutation.ClearManager()
	return fcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fcu *FleetCodeUpdate) Save(ctx context.Context) (int, error) {
	fcu.defaults()
	return withHooks(ctx, fcu.sqlSave, fcu.mutation, fcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fcu *FleetCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := fcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fcu *FleetCodeUpdate) Exec(ctx context.Context) error {
	_, err := fcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcu *FleetCodeUpdate) ExecX(ctx context.Context) {
	if err := fcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcu *FleetCodeUpdate) defaults() {
	if _, ok := fcu.mutation.UpdatedAt(); !ok {
		v := fleetcode.UpdateDefaultUpdatedAt()
		fcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fcu *FleetCodeUpdate) check() error {
	if v, ok := fcu.mutation.Status(); ok {
		if err := fleetcode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "FleetCode.status": %w`, err)}
		}
	}
	if v, ok := fcu.mutation.Code(); ok {
		if err := fleetcode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "FleetCode.code": %w`, err)}
		}
	}
	if _, ok := fcu.mutation.BusinessUnitID(); fcu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FleetCode.business_unit"`)
	}
	if _, ok := fcu.mutation.OrganizationID(); fcu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FleetCode.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fcu *FleetCodeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FleetCodeUpdate {
	fcu.modifiers = append(fcu.modifiers, modifiers...)
	return fcu
}

func (fcu *FleetCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(fleetcode.Table, fleetcode.Columns, sqlgraph.NewFieldSpec(fleetcode.FieldID, field.TypeUUID))
	if ps := fcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fcu.mutation.UpdatedAt(); ok {
		_spec.SetField(fleetcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fcu.mutation.Version(); ok {
		_spec.SetField(fleetcode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := fcu.mutation.AddedVersion(); ok {
		_spec.AddField(fleetcode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := fcu.mutation.Status(); ok {
		_spec.SetField(fleetcode.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := fcu.mutation.Code(); ok {
		_spec.SetField(fleetcode.FieldCode, field.TypeString, value)
	}
	if value, ok := fcu.mutation.Description(); ok {
		_spec.SetField(fleetcode.FieldDescription, field.TypeString, value)
	}
	if fcu.mutation.DescriptionCleared() {
		_spec.ClearField(fleetcode.FieldDescription, field.TypeString)
	}
	if value, ok := fcu.mutation.RevenueGoal(); ok {
		_spec.SetField(fleetcode.FieldRevenueGoal, field.TypeFloat64, value)
	}
	if value, ok := fcu.mutation.AddedRevenueGoal(); ok {
		_spec.AddField(fleetcode.FieldRevenueGoal, field.TypeFloat64, value)
	}
	if fcu.mutation.RevenueGoalCleared() {
		_spec.ClearField(fleetcode.FieldRevenueGoal, field.TypeFloat64)
	}
	if value, ok := fcu.mutation.DeadheadGoal(); ok {
		_spec.SetField(fleetcode.FieldDeadheadGoal, field.TypeFloat64, value)
	}
	if value, ok := fcu.mutation.AddedDeadheadGoal(); ok {
		_spec.AddField(fleetcode.FieldDeadheadGoal, field.TypeFloat64, value)
	}
	if fcu.mutation.DeadheadGoalCleared() {
		_spec.ClearField(fleetcode.FieldDeadheadGoal, field.TypeFloat64)
	}
	if value, ok := fcu.mutation.MileageGoal(); ok {
		_spec.SetField(fleetcode.FieldMileageGoal, field.TypeFloat64, value)
	}
	if value, ok := fcu.mutation.AddedMileageGoal(); ok {
		_spec.AddField(fleetcode.FieldMileageGoal, field.TypeFloat64, value)
	}
	if fcu.mutation.MileageGoalCleared() {
		_spec.ClearField(fleetcode.FieldMileageGoal, field.TypeFloat64)
	}
	if value, ok := fcu.mutation.Color(); ok {
		_spec.SetField(fleetcode.FieldColor, field.TypeString, value)
	}
	if fcu.mutation.ColorCleared() {
		_spec.ClearField(fleetcode.FieldColor, field.TypeString)
	}
	if fcu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.OrganizationTable,
			Columns: []string{fleetcode.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.OrganizationTable,
			Columns: []string{fleetcode.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fcu.mutation.ManagerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.ManagerTable,
			Columns: []string{fleetcode.ManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcu.mutation.ManagerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.ManagerTable,
			Columns: []string{fleetcode.ManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(fcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, fcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fleetcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fcu.mutation.done = true
	return n, nil
}

// FleetCodeUpdateOne is the builder for updating a single FleetCode entity.
type FleetCodeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FleetCodeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetOrganizationID sets the "organization_id" field.
func (fcuo *FleetCodeUpdateOne) SetOrganizationID(u uuid.UUID) *FleetCodeUpdateOne {
	fcuo.mutation.SetOrganizationID(u)
	return fcuo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableOrganizationID(u *uuid.UUID) *FleetCodeUpdateOne {
	if u != nil {
		fcuo.SetOrganizationID(*u)
	}
	return fcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fcuo *FleetCodeUpdateOne) SetUpdatedAt(t time.Time) *FleetCodeUpdateOne {
	fcuo.mutation.SetUpdatedAt(t)
	return fcuo
}

// SetVersion sets the "version" field.
func (fcuo *FleetCodeUpdateOne) SetVersion(i int) *FleetCodeUpdateOne {
	fcuo.mutation.ResetVersion()
	fcuo.mutation.SetVersion(i)
	return fcuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableVersion(i *int) *FleetCodeUpdateOne {
	if i != nil {
		fcuo.SetVersion(*i)
	}
	return fcuo
}

// AddVersion adds i to the "version" field.
func (fcuo *FleetCodeUpdateOne) AddVersion(i int) *FleetCodeUpdateOne {
	fcuo.mutation.AddVersion(i)
	return fcuo
}

// SetStatus sets the "status" field.
func (fcuo *FleetCodeUpdateOne) SetStatus(f fleetcode.Status) *FleetCodeUpdateOne {
	fcuo.mutation.SetStatus(f)
	return fcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableStatus(f *fleetcode.Status) *FleetCodeUpdateOne {
	if f != nil {
		fcuo.SetStatus(*f)
	}
	return fcuo
}

// SetCode sets the "code" field.
func (fcuo *FleetCodeUpdateOne) SetCode(s string) *FleetCodeUpdateOne {
	fcuo.mutation.SetCode(s)
	return fcuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableCode(s *string) *FleetCodeUpdateOne {
	if s != nil {
		fcuo.SetCode(*s)
	}
	return fcuo
}

// SetDescription sets the "description" field.
func (fcuo *FleetCodeUpdateOne) SetDescription(s string) *FleetCodeUpdateOne {
	fcuo.mutation.SetDescription(s)
	return fcuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableDescription(s *string) *FleetCodeUpdateOne {
	if s != nil {
		fcuo.SetDescription(*s)
	}
	return fcuo
}

// ClearDescription clears the value of the "description" field.
func (fcuo *FleetCodeUpdateOne) ClearDescription() *FleetCodeUpdateOne {
	fcuo.mutation.ClearDescription()
	return fcuo
}

// SetRevenueGoal sets the "revenue_goal" field.
func (fcuo *FleetCodeUpdateOne) SetRevenueGoal(f float64) *FleetCodeUpdateOne {
	fcuo.mutation.ResetRevenueGoal()
	fcuo.mutation.SetRevenueGoal(f)
	return fcuo
}

// SetNillableRevenueGoal sets the "revenue_goal" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableRevenueGoal(f *float64) *FleetCodeUpdateOne {
	if f != nil {
		fcuo.SetRevenueGoal(*f)
	}
	return fcuo
}

// AddRevenueGoal adds f to the "revenue_goal" field.
func (fcuo *FleetCodeUpdateOne) AddRevenueGoal(f float64) *FleetCodeUpdateOne {
	fcuo.mutation.AddRevenueGoal(f)
	return fcuo
}

// ClearRevenueGoal clears the value of the "revenue_goal" field.
func (fcuo *FleetCodeUpdateOne) ClearRevenueGoal() *FleetCodeUpdateOne {
	fcuo.mutation.ClearRevenueGoal()
	return fcuo
}

// SetDeadheadGoal sets the "deadhead_goal" field.
func (fcuo *FleetCodeUpdateOne) SetDeadheadGoal(f float64) *FleetCodeUpdateOne {
	fcuo.mutation.ResetDeadheadGoal()
	fcuo.mutation.SetDeadheadGoal(f)
	return fcuo
}

// SetNillableDeadheadGoal sets the "deadhead_goal" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableDeadheadGoal(f *float64) *FleetCodeUpdateOne {
	if f != nil {
		fcuo.SetDeadheadGoal(*f)
	}
	return fcuo
}

// AddDeadheadGoal adds f to the "deadhead_goal" field.
func (fcuo *FleetCodeUpdateOne) AddDeadheadGoal(f float64) *FleetCodeUpdateOne {
	fcuo.mutation.AddDeadheadGoal(f)
	return fcuo
}

// ClearDeadheadGoal clears the value of the "deadhead_goal" field.
func (fcuo *FleetCodeUpdateOne) ClearDeadheadGoal() *FleetCodeUpdateOne {
	fcuo.mutation.ClearDeadheadGoal()
	return fcuo
}

// SetMileageGoal sets the "mileage_goal" field.
func (fcuo *FleetCodeUpdateOne) SetMileageGoal(f float64) *FleetCodeUpdateOne {
	fcuo.mutation.ResetMileageGoal()
	fcuo.mutation.SetMileageGoal(f)
	return fcuo
}

// SetNillableMileageGoal sets the "mileage_goal" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableMileageGoal(f *float64) *FleetCodeUpdateOne {
	if f != nil {
		fcuo.SetMileageGoal(*f)
	}
	return fcuo
}

// AddMileageGoal adds f to the "mileage_goal" field.
func (fcuo *FleetCodeUpdateOne) AddMileageGoal(f float64) *FleetCodeUpdateOne {
	fcuo.mutation.AddMileageGoal(f)
	return fcuo
}

// ClearMileageGoal clears the value of the "mileage_goal" field.
func (fcuo *FleetCodeUpdateOne) ClearMileageGoal() *FleetCodeUpdateOne {
	fcuo.mutation.ClearMileageGoal()
	return fcuo
}

// SetManagerID sets the "manager_id" field.
func (fcuo *FleetCodeUpdateOne) SetManagerID(u uuid.UUID) *FleetCodeUpdateOne {
	fcuo.mutation.SetManagerID(u)
	return fcuo
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableManagerID(u *uuid.UUID) *FleetCodeUpdateOne {
	if u != nil {
		fcuo.SetManagerID(*u)
	}
	return fcuo
}

// ClearManagerID clears the value of the "manager_id" field.
func (fcuo *FleetCodeUpdateOne) ClearManagerID() *FleetCodeUpdateOne {
	fcuo.mutation.ClearManagerID()
	return fcuo
}

// SetColor sets the "color" field.
func (fcuo *FleetCodeUpdateOne) SetColor(s string) *FleetCodeUpdateOne {
	fcuo.mutation.SetColor(s)
	return fcuo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (fcuo *FleetCodeUpdateOne) SetNillableColor(s *string) *FleetCodeUpdateOne {
	if s != nil {
		fcuo.SetColor(*s)
	}
	return fcuo
}

// ClearColor clears the value of the "color" field.
func (fcuo *FleetCodeUpdateOne) ClearColor() *FleetCodeUpdateOne {
	fcuo.mutation.ClearColor()
	return fcuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (fcuo *FleetCodeUpdateOne) SetOrganization(o *Organization) *FleetCodeUpdateOne {
	return fcuo.SetOrganizationID(o.ID)
}

// SetManager sets the "manager" edge to the User entity.
func (fcuo *FleetCodeUpdateOne) SetManager(u *User) *FleetCodeUpdateOne {
	return fcuo.SetManagerID(u.ID)
}

// Mutation returns the FleetCodeMutation object of the builder.
func (fcuo *FleetCodeUpdateOne) Mutation() *FleetCodeMutation {
	return fcuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (fcuo *FleetCodeUpdateOne) ClearOrganization() *FleetCodeUpdateOne {
	fcuo.mutation.ClearOrganization()
	return fcuo
}

// ClearManager clears the "manager" edge to the User entity.
func (fcuo *FleetCodeUpdateOne) ClearManager() *FleetCodeUpdateOne {
	fcuo.mutation.ClearManager()
	return fcuo
}

// Where appends a list predicates to the FleetCodeUpdate builder.
func (fcuo *FleetCodeUpdateOne) Where(ps ...predicate.FleetCode) *FleetCodeUpdateOne {
	fcuo.mutation.Where(ps...)
	return fcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fcuo *FleetCodeUpdateOne) Select(field string, fields ...string) *FleetCodeUpdateOne {
	fcuo.fields = append([]string{field}, fields...)
	return fcuo
}

// Save executes the query and returns the updated FleetCode entity.
func (fcuo *FleetCodeUpdateOne) Save(ctx context.Context) (*FleetCode, error) {
	fcuo.defaults()
	return withHooks(ctx, fcuo.sqlSave, fcuo.mutation, fcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fcuo *FleetCodeUpdateOne) SaveX(ctx context.Context) *FleetCode {
	node, err := fcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fcuo *FleetCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := fcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcuo *FleetCodeUpdateOne) ExecX(ctx context.Context) {
	if err := fcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcuo *FleetCodeUpdateOne) defaults() {
	if _, ok := fcuo.mutation.UpdatedAt(); !ok {
		v := fleetcode.UpdateDefaultUpdatedAt()
		fcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fcuo *FleetCodeUpdateOne) check() error {
	if v, ok := fcuo.mutation.Status(); ok {
		if err := fleetcode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "FleetCode.status": %w`, err)}
		}
	}
	if v, ok := fcuo.mutation.Code(); ok {
		if err := fleetcode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "FleetCode.code": %w`, err)}
		}
	}
	if _, ok := fcuo.mutation.BusinessUnitID(); fcuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FleetCode.business_unit"`)
	}
	if _, ok := fcuo.mutation.OrganizationID(); fcuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FleetCode.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fcuo *FleetCodeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FleetCodeUpdateOne {
	fcuo.modifiers = append(fcuo.modifiers, modifiers...)
	return fcuo
}

func (fcuo *FleetCodeUpdateOne) sqlSave(ctx context.Context) (_node *FleetCode, err error) {
	if err := fcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(fleetcode.Table, fleetcode.Columns, sqlgraph.NewFieldSpec(fleetcode.FieldID, field.TypeUUID))
	id, ok := fcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FleetCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fleetcode.FieldID)
		for _, f := range fields {
			if !fleetcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fleetcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(fleetcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fcuo.mutation.Version(); ok {
		_spec.SetField(fleetcode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := fcuo.mutation.AddedVersion(); ok {
		_spec.AddField(fleetcode.FieldVersion, field.TypeInt, value)
	}
	if value, ok := fcuo.mutation.Status(); ok {
		_spec.SetField(fleetcode.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := fcuo.mutation.Code(); ok {
		_spec.SetField(fleetcode.FieldCode, field.TypeString, value)
	}
	if value, ok := fcuo.mutation.Description(); ok {
		_spec.SetField(fleetcode.FieldDescription, field.TypeString, value)
	}
	if fcuo.mutation.DescriptionCleared() {
		_spec.ClearField(fleetcode.FieldDescription, field.TypeString)
	}
	if value, ok := fcuo.mutation.RevenueGoal(); ok {
		_spec.SetField(fleetcode.FieldRevenueGoal, field.TypeFloat64, value)
	}
	if value, ok := fcuo.mutation.AddedRevenueGoal(); ok {
		_spec.AddField(fleetcode.FieldRevenueGoal, field.TypeFloat64, value)
	}
	if fcuo.mutation.RevenueGoalCleared() {
		_spec.ClearField(fleetcode.FieldRevenueGoal, field.TypeFloat64)
	}
	if value, ok := fcuo.mutation.DeadheadGoal(); ok {
		_spec.SetField(fleetcode.FieldDeadheadGoal, field.TypeFloat64, value)
	}
	if value, ok := fcuo.mutation.AddedDeadheadGoal(); ok {
		_spec.AddField(fleetcode.FieldDeadheadGoal, field.TypeFloat64, value)
	}
	if fcuo.mutation.DeadheadGoalCleared() {
		_spec.ClearField(fleetcode.FieldDeadheadGoal, field.TypeFloat64)
	}
	if value, ok := fcuo.mutation.MileageGoal(); ok {
		_spec.SetField(fleetcode.FieldMileageGoal, field.TypeFloat64, value)
	}
	if value, ok := fcuo.mutation.AddedMileageGoal(); ok {
		_spec.AddField(fleetcode.FieldMileageGoal, field.TypeFloat64, value)
	}
	if fcuo.mutation.MileageGoalCleared() {
		_spec.ClearField(fleetcode.FieldMileageGoal, field.TypeFloat64)
	}
	if value, ok := fcuo.mutation.Color(); ok {
		_spec.SetField(fleetcode.FieldColor, field.TypeString, value)
	}
	if fcuo.mutation.ColorCleared() {
		_spec.ClearField(fleetcode.FieldColor, field.TypeString)
	}
	if fcuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.OrganizationTable,
			Columns: []string{fleetcode.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.OrganizationTable,
			Columns: []string{fleetcode.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fcuo.mutation.ManagerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.ManagerTable,
			Columns: []string{fleetcode.ManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcuo.mutation.ManagerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fleetcode.ManagerTable,
			Columns: []string{fleetcode.ManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(fcuo.modifiers...)
	_node = &FleetCode{config: fcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fleetcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fcuo.mutation.done = true
	return _node, nil
}
