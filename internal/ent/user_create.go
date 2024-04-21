// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/shipment"
	"github.com/emoss08/trenova/internal/ent/shipmentcharges"
	"github.com/emoss08/trenova/internal/ent/shipmentcomment"
	"github.com/emoss08/trenova/internal/ent/user"
	"github.com/emoss08/trenova/internal/ent/userfavorite"
	"github.com/emoss08/trenova/internal/ent/usernotification"
	"github.com/emoss08/trenova/internal/ent/userreport"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (uc *UserCreate) SetBusinessUnitID(u uuid.UUID) *UserCreate {
	uc.mutation.SetBusinessUnitID(u)
	return uc
}

// SetOrganizationID sets the "organization_id" field.
func (uc *UserCreate) SetOrganizationID(u uuid.UUID) *UserCreate {
	uc.mutation.SetOrganizationID(u)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetVersion sets the "version" field.
func (uc *UserCreate) SetVersion(i int) *UserCreate {
	uc.mutation.SetVersion(i)
	return uc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (uc *UserCreate) SetNillableVersion(i *int) *UserCreate {
	if i != nil {
		uc.SetVersion(*i)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(u user.Status) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(u *user.Status) *UserCreate {
	if u != nil {
		uc.SetStatus(*u)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetTimezone sets the "timezone" field.
func (uc *UserCreate) SetTimezone(u user.Timezone) *UserCreate {
	uc.mutation.SetTimezone(u)
	return uc
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (uc *UserCreate) SetNillableTimezone(u *user.Timezone) *UserCreate {
	if u != nil {
		uc.SetTimezone(*u)
	}
	return uc
}

// SetProfilePicURL sets the "profile_pic_url" field.
func (uc *UserCreate) SetProfilePicURL(s string) *UserCreate {
	uc.mutation.SetProfilePicURL(s)
	return uc
}

// SetNillableProfilePicURL sets the "profile_pic_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableProfilePicURL(s *string) *UserCreate {
	if s != nil {
		uc.SetProfilePicURL(*s)
	}
	return uc
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (uc *UserCreate) SetThumbnailURL(s string) *UserCreate {
	uc.mutation.SetThumbnailURL(s)
	return uc
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableThumbnailURL(s *string) *UserCreate {
	if s != nil {
		uc.SetThumbnailURL(*s)
	}
	return uc
}

// SetPhoneNumber sets the "phone_number" field.
func (uc *UserCreate) SetPhoneNumber(s string) *UserCreate {
	uc.mutation.SetPhoneNumber(s)
	return uc
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhoneNumber(s *string) *UserCreate {
	if s != nil {
		uc.SetPhoneNumber(*s)
	}
	return uc
}

// SetIsAdmin sets the "is_admin" field.
func (uc *UserCreate) SetIsAdmin(b bool) *UserCreate {
	uc.mutation.SetIsAdmin(b)
	return uc
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (uc *UserCreate) SetNillableIsAdmin(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsAdmin(*b)
	}
	return uc
}

// SetIsSuperAdmin sets the "is_super_admin" field.
func (uc *UserCreate) SetIsSuperAdmin(b bool) *UserCreate {
	uc.mutation.SetIsSuperAdmin(b)
	return uc
}

// SetNillableIsSuperAdmin sets the "is_super_admin" field if the given value is not nil.
func (uc *UserCreate) SetNillableIsSuperAdmin(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsSuperAdmin(*b)
	}
	return uc
}

// SetLastLogin sets the "last_login" field.
func (uc *UserCreate) SetLastLogin(t time.Time) *UserCreate {
	uc.mutation.SetLastLogin(t)
	return uc
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastLogin(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastLogin(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (uc *UserCreate) SetBusinessUnit(b *BusinessUnit) *UserCreate {
	return uc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (uc *UserCreate) SetOrganization(o *Organization) *UserCreate {
	return uc.SetOrganizationID(o.ID)
}

// AddUserFavoriteIDs adds the "user_favorites" edge to the UserFavorite entity by IDs.
func (uc *UserCreate) AddUserFavoriteIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddUserFavoriteIDs(ids...)
	return uc
}

// AddUserFavorites adds the "user_favorites" edges to the UserFavorite entity.
func (uc *UserCreate) AddUserFavorites(u ...*UserFavorite) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserFavoriteIDs(ids...)
}

// AddUserNotificationIDs adds the "user_notifications" edge to the UserNotification entity by IDs.
func (uc *UserCreate) AddUserNotificationIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddUserNotificationIDs(ids...)
	return uc
}

// AddUserNotifications adds the "user_notifications" edges to the UserNotification entity.
func (uc *UserCreate) AddUserNotifications(u ...*UserNotification) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserNotificationIDs(ids...)
}

// AddShipmentIDs adds the "shipments" edge to the Shipment entity by IDs.
func (uc *UserCreate) AddShipmentIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddShipmentIDs(ids...)
	return uc
}

// AddShipments adds the "shipments" edges to the Shipment entity.
func (uc *UserCreate) AddShipments(s ...*Shipment) *UserCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddShipmentIDs(ids...)
}

// AddShipmentCommentIDs adds the "shipment_comments" edge to the ShipmentComment entity by IDs.
func (uc *UserCreate) AddShipmentCommentIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddShipmentCommentIDs(ids...)
	return uc
}

// AddShipmentComments adds the "shipment_comments" edges to the ShipmentComment entity.
func (uc *UserCreate) AddShipmentComments(s ...*ShipmentComment) *UserCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddShipmentCommentIDs(ids...)
}

// AddShipmentChargeIDs adds the "shipment_charges" edge to the ShipmentCharges entity by IDs.
func (uc *UserCreate) AddShipmentChargeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddShipmentChargeIDs(ids...)
	return uc
}

// AddShipmentCharges adds the "shipment_charges" edges to the ShipmentCharges entity.
func (uc *UserCreate) AddShipmentCharges(s ...*ShipmentCharges) *UserCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddShipmentChargeIDs(ids...)
}

// AddReportIDs adds the "reports" edge to the UserReport entity by IDs.
func (uc *UserCreate) AddReportIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddReportIDs(ids...)
	return uc
}

// AddReports adds the "reports" edges to the UserReport entity.
func (uc *UserCreate) AddReports(u ...*UserReport) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddReportIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.Version(); !ok {
		v := user.DefaultVersion
		uc.mutation.SetVersion(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.Timezone(); !ok {
		v := user.DefaultTimezone
		uc.mutation.SetTimezone(v)
	}
	if _, ok := uc.mutation.IsAdmin(); !ok {
		v := user.DefaultIsAdmin
		uc.mutation.SetIsAdmin(v)
	}
	if _, ok := uc.mutation.IsSuperAdmin(); !ok {
		v := user.DefaultIsSuperAdmin
		uc.mutation.SetIsSuperAdmin(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "User.business_unit_id"`)}
	}
	if _, ok := uc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "User.organization_id"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "User.version"`)}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Timezone(); !ok {
		return &ValidationError{Name: "timezone", err: errors.New(`ent: missing required field "User.timezone"`)}
	}
	if v, ok := uc.mutation.Timezone(); ok {
		if err := user.TimezoneValidator(v); err != nil {
			return &ValidationError{Name: "timezone", err: fmt.Errorf(`ent: validator failed for field "User.timezone": %w`, err)}
		}
	}
	if _, ok := uc.mutation.IsAdmin(); !ok {
		return &ValidationError{Name: "is_admin", err: errors.New(`ent: missing required field "User.is_admin"`)}
	}
	if _, ok := uc.mutation.IsSuperAdmin(); !ok {
		return &ValidationError{Name: "is_super_admin", err: errors.New(`ent: missing required field "User.is_super_admin"`)}
	}
	if _, ok := uc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "User.business_unit"`)}
	}
	if _, ok := uc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "User.organization"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
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
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Version(); ok {
		_spec.SetField(user.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Timezone(); ok {
		_spec.SetField(user.FieldTimezone, field.TypeEnum, value)
		_node.Timezone = value
	}
	if value, ok := uc.mutation.ProfilePicURL(); ok {
		_spec.SetField(user.FieldProfilePicURL, field.TypeString, value)
		_node.ProfilePicURL = value
	}
	if value, ok := uc.mutation.ThumbnailURL(); ok {
		_spec.SetField(user.FieldThumbnailURL, field.TypeString, value)
		_node.ThumbnailURL = value
	}
	if value, ok := uc.mutation.PhoneNumber(); ok {
		_spec.SetField(user.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := uc.mutation.IsAdmin(); ok {
		_spec.SetField(user.FieldIsAdmin, field.TypeBool, value)
		_node.IsAdmin = value
	}
	if value, ok := uc.mutation.IsSuperAdmin(); ok {
		_spec.SetField(user.FieldIsSuperAdmin, field.TypeBool, value)
		_node.IsSuperAdmin = value
	}
	if value, ok := uc.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
		_node.LastLogin = &value
	}
	if nodes := uc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.BusinessUnitTable,
			Columns: []string{user.BusinessUnitColumn},
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
	if nodes := uc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.OrganizationTable,
			Columns: []string{user.OrganizationColumn},
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
	if nodes := uc.mutation.UserFavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserFavoritesTable,
			Columns: []string{user.UserFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userfavorite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserNotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserNotificationsTable,
			Columns: []string{user.UserNotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usernotification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ShipmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShipmentsTable,
			Columns: []string{user.ShipmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ShipmentCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShipmentCommentsTable,
			Columns: []string{user.ShipmentCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmentcomment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ShipmentChargesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShipmentChargesTable,
			Columns: []string{user.ShipmentChargesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmentcharges.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ReportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReportsTable,
			Columns: []string{user.ReportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userreport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
