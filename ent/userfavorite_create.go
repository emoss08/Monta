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
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/user"
	"github.com/emoss08/trenova/ent/userfavorite"
	"github.com/google/uuid"
)

// UserFavoriteCreate is the builder for creating a UserFavorite entity.
type UserFavoriteCreate struct {
	config
	mutation *UserFavoriteMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (ufc *UserFavoriteCreate) SetBusinessUnitID(u uuid.UUID) *UserFavoriteCreate {
	ufc.mutation.SetBusinessUnitID(u)
	return ufc
}

// SetOrganizationID sets the "organization_id" field.
func (ufc *UserFavoriteCreate) SetOrganizationID(u uuid.UUID) *UserFavoriteCreate {
	ufc.mutation.SetOrganizationID(u)
	return ufc
}

// SetCreatedAt sets the "created_at" field.
func (ufc *UserFavoriteCreate) SetCreatedAt(t time.Time) *UserFavoriteCreate {
	ufc.mutation.SetCreatedAt(t)
	return ufc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ufc *UserFavoriteCreate) SetNillableCreatedAt(t *time.Time) *UserFavoriteCreate {
	if t != nil {
		ufc.SetCreatedAt(*t)
	}
	return ufc
}

// SetUpdatedAt sets the "updated_at" field.
func (ufc *UserFavoriteCreate) SetUpdatedAt(t time.Time) *UserFavoriteCreate {
	ufc.mutation.SetUpdatedAt(t)
	return ufc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ufc *UserFavoriteCreate) SetNillableUpdatedAt(t *time.Time) *UserFavoriteCreate {
	if t != nil {
		ufc.SetUpdatedAt(*t)
	}
	return ufc
}

// SetPageLink sets the "page_link" field.
func (ufc *UserFavoriteCreate) SetPageLink(s string) *UserFavoriteCreate {
	ufc.mutation.SetPageLink(s)
	return ufc
}

// SetUserID sets the "user_id" field.
func (ufc *UserFavoriteCreate) SetUserID(u uuid.UUID) *UserFavoriteCreate {
	ufc.mutation.SetUserID(u)
	return ufc
}

// SetID sets the "id" field.
func (ufc *UserFavoriteCreate) SetID(u uuid.UUID) *UserFavoriteCreate {
	ufc.mutation.SetID(u)
	return ufc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ufc *UserFavoriteCreate) SetNillableID(u *uuid.UUID) *UserFavoriteCreate {
	if u != nil {
		ufc.SetID(*u)
	}
	return ufc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (ufc *UserFavoriteCreate) SetBusinessUnit(b *BusinessUnit) *UserFavoriteCreate {
	return ufc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ufc *UserFavoriteCreate) SetOrganization(o *Organization) *UserFavoriteCreate {
	return ufc.SetOrganizationID(o.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ufc *UserFavoriteCreate) SetUser(u *User) *UserFavoriteCreate {
	return ufc.SetUserID(u.ID)
}

// Mutation returns the UserFavoriteMutation object of the builder.
func (ufc *UserFavoriteCreate) Mutation() *UserFavoriteMutation {
	return ufc.mutation
}

// Save creates the UserFavorite in the database.
func (ufc *UserFavoriteCreate) Save(ctx context.Context) (*UserFavorite, error) {
	ufc.defaults()
	return withHooks(ctx, ufc.sqlSave, ufc.mutation, ufc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ufc *UserFavoriteCreate) SaveX(ctx context.Context) *UserFavorite {
	v, err := ufc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufc *UserFavoriteCreate) Exec(ctx context.Context) error {
	_, err := ufc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufc *UserFavoriteCreate) ExecX(ctx context.Context) {
	if err := ufc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufc *UserFavoriteCreate) defaults() {
	if _, ok := ufc.mutation.CreatedAt(); !ok {
		v := userfavorite.DefaultCreatedAt
		ufc.mutation.SetCreatedAt(v)
	}
	if _, ok := ufc.mutation.UpdatedAt(); !ok {
		v := userfavorite.DefaultUpdatedAt
		ufc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ufc.mutation.ID(); !ok {
		v := userfavorite.DefaultID()
		ufc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufc *UserFavoriteCreate) check() error {
	if _, ok := ufc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "UserFavorite.business_unit_id"`)}
	}
	if _, ok := ufc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "UserFavorite.organization_id"`)}
	}
	if _, ok := ufc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserFavorite.created_at"`)}
	}
	if _, ok := ufc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserFavorite.updated_at"`)}
	}
	if _, ok := ufc.mutation.PageLink(); !ok {
		return &ValidationError{Name: "page_link", err: errors.New(`ent: missing required field "UserFavorite.page_link"`)}
	}
	if v, ok := ufc.mutation.PageLink(); ok {
		if err := userfavorite.PageLinkValidator(v); err != nil {
			return &ValidationError{Name: "page_link", err: fmt.Errorf(`ent: validator failed for field "UserFavorite.page_link": %w`, err)}
		}
	}
	if _, ok := ufc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserFavorite.user_id"`)}
	}
	if _, ok := ufc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "UserFavorite.business_unit"`)}
	}
	if _, ok := ufc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "UserFavorite.organization"`)}
	}
	if _, ok := ufc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserFavorite.user"`)}
	}
	return nil
}

func (ufc *UserFavoriteCreate) sqlSave(ctx context.Context) (*UserFavorite, error) {
	if err := ufc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ufc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ufc.driver, _spec); err != nil {
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
	ufc.mutation.id = &_node.ID
	ufc.mutation.done = true
	return _node, nil
}

func (ufc *UserFavoriteCreate) createSpec() (*UserFavorite, *sqlgraph.CreateSpec) {
	var (
		_node = &UserFavorite{config: ufc.config}
		_spec = sqlgraph.NewCreateSpec(userfavorite.Table, sqlgraph.NewFieldSpec(userfavorite.FieldID, field.TypeUUID))
	)
	if id, ok := ufc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ufc.mutation.CreatedAt(); ok {
		_spec.SetField(userfavorite.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ufc.mutation.UpdatedAt(); ok {
		_spec.SetField(userfavorite.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ufc.mutation.PageLink(); ok {
		_spec.SetField(userfavorite.FieldPageLink, field.TypeString, value)
		_node.PageLink = value
	}
	if nodes := ufc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userfavorite.BusinessUnitTable,
			Columns: []string{userfavorite.BusinessUnitColumn},
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
	if nodes := ufc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userfavorite.OrganizationTable,
			Columns: []string{userfavorite.OrganizationColumn},
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
	if nodes := ufc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfavorite.UserTable,
			Columns: []string{userfavorite.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserFavoriteCreateBulk is the builder for creating many UserFavorite entities in bulk.
type UserFavoriteCreateBulk struct {
	config
	err      error
	builders []*UserFavoriteCreate
}

// Save creates the UserFavorite entities in the database.
func (ufcb *UserFavoriteCreateBulk) Save(ctx context.Context) ([]*UserFavorite, error) {
	if ufcb.err != nil {
		return nil, ufcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ufcb.builders))
	nodes := make([]*UserFavorite, len(ufcb.builders))
	mutators := make([]Mutator, len(ufcb.builders))
	for i := range ufcb.builders {
		func(i int, root context.Context) {
			builder := ufcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserFavoriteMutation)
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
					_, err = mutators[i+1].Mutate(root, ufcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ufcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ufcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ufcb *UserFavoriteCreateBulk) SaveX(ctx context.Context) []*UserFavorite {
	v, err := ufcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufcb *UserFavoriteCreateBulk) Exec(ctx context.Context) error {
	_, err := ufcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufcb *UserFavoriteCreateBulk) ExecX(ctx context.Context) {
	if err := ufcb.Exec(ctx); err != nil {
		panic(err)
	}
}
