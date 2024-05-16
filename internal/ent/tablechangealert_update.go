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
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/tablechangealert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// TableChangeAlertUpdate is the builder for updating TableChangeAlert entities.
type TableChangeAlertUpdate struct {
	config
	hooks     []Hook
	mutation  *TableChangeAlertMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TableChangeAlertUpdate builder.
func (tcau *TableChangeAlertUpdate) Where(ps ...predicate.TableChangeAlert) *TableChangeAlertUpdate {
	tcau.mutation.Where(ps...)
	return tcau
}

// SetOrganizationID sets the "organization_id" field.
func (tcau *TableChangeAlertUpdate) SetOrganizationID(u uuid.UUID) *TableChangeAlertUpdate {
	tcau.mutation.SetOrganizationID(u)
	return tcau
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableOrganizationID(u *uuid.UUID) *TableChangeAlertUpdate {
	if u != nil {
		tcau.SetOrganizationID(*u)
	}
	return tcau
}

// SetUpdatedAt sets the "updated_at" field.
func (tcau *TableChangeAlertUpdate) SetUpdatedAt(t time.Time) *TableChangeAlertUpdate {
	tcau.mutation.SetUpdatedAt(t)
	return tcau
}

// SetVersion sets the "version" field.
func (tcau *TableChangeAlertUpdate) SetVersion(i int) *TableChangeAlertUpdate {
	tcau.mutation.ResetVersion()
	tcau.mutation.SetVersion(i)
	return tcau
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableVersion(i *int) *TableChangeAlertUpdate {
	if i != nil {
		tcau.SetVersion(*i)
	}
	return tcau
}

// AddVersion adds i to the "version" field.
func (tcau *TableChangeAlertUpdate) AddVersion(i int) *TableChangeAlertUpdate {
	tcau.mutation.AddVersion(i)
	return tcau
}

// SetStatus sets the "status" field.
func (tcau *TableChangeAlertUpdate) SetStatus(t tablechangealert.Status) *TableChangeAlertUpdate {
	tcau.mutation.SetStatus(t)
	return tcau
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableStatus(t *tablechangealert.Status) *TableChangeAlertUpdate {
	if t != nil {
		tcau.SetStatus(*t)
	}
	return tcau
}

// SetName sets the "name" field.
func (tcau *TableChangeAlertUpdate) SetName(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetName(s)
	return tcau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableName(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetName(*s)
	}
	return tcau
}

// SetDatabaseAction sets the "database_action" field.
func (tcau *TableChangeAlertUpdate) SetDatabaseAction(ta tablechangealert.DatabaseAction) *TableChangeAlertUpdate {
	tcau.mutation.SetDatabaseAction(ta)
	return tcau
}

// SetNillableDatabaseAction sets the "database_action" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableDatabaseAction(ta *tablechangealert.DatabaseAction) *TableChangeAlertUpdate {
	if ta != nil {
		tcau.SetDatabaseAction(*ta)
	}
	return tcau
}

// SetSource sets the "source" field.
func (tcau *TableChangeAlertUpdate) SetSource(t tablechangealert.Source) *TableChangeAlertUpdate {
	tcau.mutation.SetSource(t)
	return tcau
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableSource(t *tablechangealert.Source) *TableChangeAlertUpdate {
	if t != nil {
		tcau.SetSource(*t)
	}
	return tcau
}

// SetTableName sets the "table_name" field.
func (tcau *TableChangeAlertUpdate) SetTableName(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetTableName(s)
	return tcau
}

// SetNillableTableName sets the "table_name" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableTableName(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetTableName(*s)
	}
	return tcau
}

// ClearTableName clears the value of the "table_name" field.
func (tcau *TableChangeAlertUpdate) ClearTableName() *TableChangeAlertUpdate {
	tcau.mutation.ClearTableName()
	return tcau
}

// SetTopicName sets the "topic_name" field.
func (tcau *TableChangeAlertUpdate) SetTopicName(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetTopicName(s)
	return tcau
}

// SetNillableTopicName sets the "topic_name" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableTopicName(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetTopicName(*s)
	}
	return tcau
}

// ClearTopicName clears the value of the "topic_name" field.
func (tcau *TableChangeAlertUpdate) ClearTopicName() *TableChangeAlertUpdate {
	tcau.mutation.ClearTopicName()
	return tcau
}

// SetDescription sets the "description" field.
func (tcau *TableChangeAlertUpdate) SetDescription(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetDescription(s)
	return tcau
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableDescription(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetDescription(*s)
	}
	return tcau
}

// ClearDescription clears the value of the "description" field.
func (tcau *TableChangeAlertUpdate) ClearDescription() *TableChangeAlertUpdate {
	tcau.mutation.ClearDescription()
	return tcau
}

// SetCustomSubject sets the "custom_subject" field.
func (tcau *TableChangeAlertUpdate) SetCustomSubject(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetCustomSubject(s)
	return tcau
}

// SetNillableCustomSubject sets the "custom_subject" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableCustomSubject(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetCustomSubject(*s)
	}
	return tcau
}

// ClearCustomSubject clears the value of the "custom_subject" field.
func (tcau *TableChangeAlertUpdate) ClearCustomSubject() *TableChangeAlertUpdate {
	tcau.mutation.ClearCustomSubject()
	return tcau
}

// SetFunctionName sets the "function_name" field.
func (tcau *TableChangeAlertUpdate) SetFunctionName(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetFunctionName(s)
	return tcau
}

// SetNillableFunctionName sets the "function_name" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableFunctionName(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetFunctionName(*s)
	}
	return tcau
}

// ClearFunctionName clears the value of the "function_name" field.
func (tcau *TableChangeAlertUpdate) ClearFunctionName() *TableChangeAlertUpdate {
	tcau.mutation.ClearFunctionName()
	return tcau
}

// SetTriggerName sets the "trigger_name" field.
func (tcau *TableChangeAlertUpdate) SetTriggerName(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetTriggerName(s)
	return tcau
}

// SetNillableTriggerName sets the "trigger_name" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableTriggerName(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetTriggerName(*s)
	}
	return tcau
}

// ClearTriggerName clears the value of the "trigger_name" field.
func (tcau *TableChangeAlertUpdate) ClearTriggerName() *TableChangeAlertUpdate {
	tcau.mutation.ClearTriggerName()
	return tcau
}

// SetListenerName sets the "listener_name" field.
func (tcau *TableChangeAlertUpdate) SetListenerName(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetListenerName(s)
	return tcau
}

// SetNillableListenerName sets the "listener_name" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableListenerName(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetListenerName(*s)
	}
	return tcau
}

// ClearListenerName clears the value of the "listener_name" field.
func (tcau *TableChangeAlertUpdate) ClearListenerName() *TableChangeAlertUpdate {
	tcau.mutation.ClearListenerName()
	return tcau
}

// SetEmailRecipients sets the "email_recipients" field.
func (tcau *TableChangeAlertUpdate) SetEmailRecipients(s string) *TableChangeAlertUpdate {
	tcau.mutation.SetEmailRecipients(s)
	return tcau
}

// SetNillableEmailRecipients sets the "email_recipients" field if the given value is not nil.
func (tcau *TableChangeAlertUpdate) SetNillableEmailRecipients(s *string) *TableChangeAlertUpdate {
	if s != nil {
		tcau.SetEmailRecipients(*s)
	}
	return tcau
}

// ClearEmailRecipients clears the value of the "email_recipients" field.
func (tcau *TableChangeAlertUpdate) ClearEmailRecipients() *TableChangeAlertUpdate {
	tcau.mutation.ClearEmailRecipients()
	return tcau
}

// SetEffectiveDate sets the "effective_date" field.
func (tcau *TableChangeAlertUpdate) SetEffectiveDate(pg *pgtype.Date) *TableChangeAlertUpdate {
	tcau.mutation.SetEffectiveDate(pg)
	return tcau
}

// ClearEffectiveDate clears the value of the "effective_date" field.
func (tcau *TableChangeAlertUpdate) ClearEffectiveDate() *TableChangeAlertUpdate {
	tcau.mutation.ClearEffectiveDate()
	return tcau
}

// SetExpirationDate sets the "expiration_date" field.
func (tcau *TableChangeAlertUpdate) SetExpirationDate(pg *pgtype.Date) *TableChangeAlertUpdate {
	tcau.mutation.SetExpirationDate(pg)
	return tcau
}

// ClearExpirationDate clears the value of the "expiration_date" field.
func (tcau *TableChangeAlertUpdate) ClearExpirationDate() *TableChangeAlertUpdate {
	tcau.mutation.ClearExpirationDate()
	return tcau
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (tcau *TableChangeAlertUpdate) SetOrganization(o *Organization) *TableChangeAlertUpdate {
	return tcau.SetOrganizationID(o.ID)
}

// Mutation returns the TableChangeAlertMutation object of the builder.
func (tcau *TableChangeAlertUpdate) Mutation() *TableChangeAlertMutation {
	return tcau.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (tcau *TableChangeAlertUpdate) ClearOrganization() *TableChangeAlertUpdate {
	tcau.mutation.ClearOrganization()
	return tcau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcau *TableChangeAlertUpdate) Save(ctx context.Context) (int, error) {
	if err := tcau.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tcau.sqlSave, tcau.mutation, tcau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcau *TableChangeAlertUpdate) SaveX(ctx context.Context) int {
	affected, err := tcau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcau *TableChangeAlertUpdate) Exec(ctx context.Context) error {
	_, err := tcau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcau *TableChangeAlertUpdate) ExecX(ctx context.Context) {
	if err := tcau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcau *TableChangeAlertUpdate) defaults() error {
	if _, ok := tcau.mutation.UpdatedAt(); !ok {
		if tablechangealert.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tablechangealert.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tablechangealert.UpdateDefaultUpdatedAt()
		tcau.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tcau *TableChangeAlertUpdate) check() error {
	if v, ok := tcau.mutation.Status(); ok {
		if err := tablechangealert.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.status": %w`, err)}
		}
	}
	if v, ok := tcau.mutation.Name(); ok {
		if err := tablechangealert.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.name": %w`, err)}
		}
	}
	if v, ok := tcau.mutation.DatabaseAction(); ok {
		if err := tablechangealert.DatabaseActionValidator(v); err != nil {
			return &ValidationError{Name: "database_action", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.database_action": %w`, err)}
		}
	}
	if v, ok := tcau.mutation.Source(); ok {
		if err := tablechangealert.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.source": %w`, err)}
		}
	}
	if v, ok := tcau.mutation.FunctionName(); ok {
		if err := tablechangealert.FunctionNameValidator(v); err != nil {
			return &ValidationError{Name: "function_name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.function_name": %w`, err)}
		}
	}
	if v, ok := tcau.mutation.TriggerName(); ok {
		if err := tablechangealert.TriggerNameValidator(v); err != nil {
			return &ValidationError{Name: "trigger_name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.trigger_name": %w`, err)}
		}
	}
	if v, ok := tcau.mutation.ListenerName(); ok {
		if err := tablechangealert.ListenerNameValidator(v); err != nil {
			return &ValidationError{Name: "listener_name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.listener_name": %w`, err)}
		}
	}
	if _, ok := tcau.mutation.BusinessUnitID(); tcau.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TableChangeAlert.business_unit"`)
	}
	if _, ok := tcau.mutation.OrganizationID(); tcau.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TableChangeAlert.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tcau *TableChangeAlertUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TableChangeAlertUpdate {
	tcau.modifiers = append(tcau.modifiers, modifiers...)
	return tcau
}

func (tcau *TableChangeAlertUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tcau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tablechangealert.Table, tablechangealert.Columns, sqlgraph.NewFieldSpec(tablechangealert.FieldID, field.TypeUUID))
	if ps := tcau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcau.mutation.UpdatedAt(); ok {
		_spec.SetField(tablechangealert.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tcau.mutation.Version(); ok {
		_spec.SetField(tablechangealert.FieldVersion, field.TypeInt, value)
	}
	if value, ok := tcau.mutation.AddedVersion(); ok {
		_spec.AddField(tablechangealert.FieldVersion, field.TypeInt, value)
	}
	if value, ok := tcau.mutation.Status(); ok {
		_spec.SetField(tablechangealert.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tcau.mutation.Name(); ok {
		_spec.SetField(tablechangealert.FieldName, field.TypeString, value)
	}
	if value, ok := tcau.mutation.DatabaseAction(); ok {
		_spec.SetField(tablechangealert.FieldDatabaseAction, field.TypeEnum, value)
	}
	if value, ok := tcau.mutation.Source(); ok {
		_spec.SetField(tablechangealert.FieldSource, field.TypeEnum, value)
	}
	if value, ok := tcau.mutation.TableName(); ok {
		_spec.SetField(tablechangealert.FieldTableName, field.TypeString, value)
	}
	if tcau.mutation.TableNameCleared() {
		_spec.ClearField(tablechangealert.FieldTableName, field.TypeString)
	}
	if value, ok := tcau.mutation.TopicName(); ok {
		_spec.SetField(tablechangealert.FieldTopicName, field.TypeString, value)
	}
	if tcau.mutation.TopicNameCleared() {
		_spec.ClearField(tablechangealert.FieldTopicName, field.TypeString)
	}
	if value, ok := tcau.mutation.Description(); ok {
		_spec.SetField(tablechangealert.FieldDescription, field.TypeString, value)
	}
	if tcau.mutation.DescriptionCleared() {
		_spec.ClearField(tablechangealert.FieldDescription, field.TypeString)
	}
	if value, ok := tcau.mutation.CustomSubject(); ok {
		_spec.SetField(tablechangealert.FieldCustomSubject, field.TypeString, value)
	}
	if tcau.mutation.CustomSubjectCleared() {
		_spec.ClearField(tablechangealert.FieldCustomSubject, field.TypeString)
	}
	if value, ok := tcau.mutation.FunctionName(); ok {
		_spec.SetField(tablechangealert.FieldFunctionName, field.TypeString, value)
	}
	if tcau.mutation.FunctionNameCleared() {
		_spec.ClearField(tablechangealert.FieldFunctionName, field.TypeString)
	}
	if value, ok := tcau.mutation.TriggerName(); ok {
		_spec.SetField(tablechangealert.FieldTriggerName, field.TypeString, value)
	}
	if tcau.mutation.TriggerNameCleared() {
		_spec.ClearField(tablechangealert.FieldTriggerName, field.TypeString)
	}
	if value, ok := tcau.mutation.ListenerName(); ok {
		_spec.SetField(tablechangealert.FieldListenerName, field.TypeString, value)
	}
	if tcau.mutation.ListenerNameCleared() {
		_spec.ClearField(tablechangealert.FieldListenerName, field.TypeString)
	}
	if value, ok := tcau.mutation.EmailRecipients(); ok {
		_spec.SetField(tablechangealert.FieldEmailRecipients, field.TypeString, value)
	}
	if tcau.mutation.EmailRecipientsCleared() {
		_spec.ClearField(tablechangealert.FieldEmailRecipients, field.TypeString)
	}
	if value, ok := tcau.mutation.EffectiveDate(); ok {
		_spec.SetField(tablechangealert.FieldEffectiveDate, field.TypeOther, value)
	}
	if tcau.mutation.EffectiveDateCleared() {
		_spec.ClearField(tablechangealert.FieldEffectiveDate, field.TypeOther)
	}
	if value, ok := tcau.mutation.ExpirationDate(); ok {
		_spec.SetField(tablechangealert.FieldExpirationDate, field.TypeOther, value)
	}
	if tcau.mutation.ExpirationDateCleared() {
		_spec.ClearField(tablechangealert.FieldExpirationDate, field.TypeOther)
	}
	if tcau.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tablechangealert.OrganizationTable,
			Columns: []string{tablechangealert.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcau.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tablechangealert.OrganizationTable,
			Columns: []string{tablechangealert.OrganizationColumn},
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
	_spec.AddModifiers(tcau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tcau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tablechangealert.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcau.mutation.done = true
	return n, nil
}

// TableChangeAlertUpdateOne is the builder for updating a single TableChangeAlert entity.
type TableChangeAlertUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TableChangeAlertMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetOrganizationID sets the "organization_id" field.
func (tcauo *TableChangeAlertUpdateOne) SetOrganizationID(u uuid.UUID) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetOrganizationID(u)
	return tcauo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableOrganizationID(u *uuid.UUID) *TableChangeAlertUpdateOne {
	if u != nil {
		tcauo.SetOrganizationID(*u)
	}
	return tcauo
}

// SetUpdatedAt sets the "updated_at" field.
func (tcauo *TableChangeAlertUpdateOne) SetUpdatedAt(t time.Time) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetUpdatedAt(t)
	return tcauo
}

// SetVersion sets the "version" field.
func (tcauo *TableChangeAlertUpdateOne) SetVersion(i int) *TableChangeAlertUpdateOne {
	tcauo.mutation.ResetVersion()
	tcauo.mutation.SetVersion(i)
	return tcauo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableVersion(i *int) *TableChangeAlertUpdateOne {
	if i != nil {
		tcauo.SetVersion(*i)
	}
	return tcauo
}

// AddVersion adds i to the "version" field.
func (tcauo *TableChangeAlertUpdateOne) AddVersion(i int) *TableChangeAlertUpdateOne {
	tcauo.mutation.AddVersion(i)
	return tcauo
}

// SetStatus sets the "status" field.
func (tcauo *TableChangeAlertUpdateOne) SetStatus(t tablechangealert.Status) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetStatus(t)
	return tcauo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableStatus(t *tablechangealert.Status) *TableChangeAlertUpdateOne {
	if t != nil {
		tcauo.SetStatus(*t)
	}
	return tcauo
}

// SetName sets the "name" field.
func (tcauo *TableChangeAlertUpdateOne) SetName(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetName(s)
	return tcauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableName(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetName(*s)
	}
	return tcauo
}

// SetDatabaseAction sets the "database_action" field.
func (tcauo *TableChangeAlertUpdateOne) SetDatabaseAction(ta tablechangealert.DatabaseAction) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetDatabaseAction(ta)
	return tcauo
}

// SetNillableDatabaseAction sets the "database_action" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableDatabaseAction(ta *tablechangealert.DatabaseAction) *TableChangeAlertUpdateOne {
	if ta != nil {
		tcauo.SetDatabaseAction(*ta)
	}
	return tcauo
}

// SetSource sets the "source" field.
func (tcauo *TableChangeAlertUpdateOne) SetSource(t tablechangealert.Source) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetSource(t)
	return tcauo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableSource(t *tablechangealert.Source) *TableChangeAlertUpdateOne {
	if t != nil {
		tcauo.SetSource(*t)
	}
	return tcauo
}

// SetTableName sets the "table_name" field.
func (tcauo *TableChangeAlertUpdateOne) SetTableName(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetTableName(s)
	return tcauo
}

// SetNillableTableName sets the "table_name" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableTableName(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetTableName(*s)
	}
	return tcauo
}

// ClearTableName clears the value of the "table_name" field.
func (tcauo *TableChangeAlertUpdateOne) ClearTableName() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearTableName()
	return tcauo
}

// SetTopicName sets the "topic_name" field.
func (tcauo *TableChangeAlertUpdateOne) SetTopicName(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetTopicName(s)
	return tcauo
}

// SetNillableTopicName sets the "topic_name" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableTopicName(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetTopicName(*s)
	}
	return tcauo
}

// ClearTopicName clears the value of the "topic_name" field.
func (tcauo *TableChangeAlertUpdateOne) ClearTopicName() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearTopicName()
	return tcauo
}

// SetDescription sets the "description" field.
func (tcauo *TableChangeAlertUpdateOne) SetDescription(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetDescription(s)
	return tcauo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableDescription(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetDescription(*s)
	}
	return tcauo
}

// ClearDescription clears the value of the "description" field.
func (tcauo *TableChangeAlertUpdateOne) ClearDescription() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearDescription()
	return tcauo
}

// SetCustomSubject sets the "custom_subject" field.
func (tcauo *TableChangeAlertUpdateOne) SetCustomSubject(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetCustomSubject(s)
	return tcauo
}

// SetNillableCustomSubject sets the "custom_subject" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableCustomSubject(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetCustomSubject(*s)
	}
	return tcauo
}

// ClearCustomSubject clears the value of the "custom_subject" field.
func (tcauo *TableChangeAlertUpdateOne) ClearCustomSubject() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearCustomSubject()
	return tcauo
}

// SetFunctionName sets the "function_name" field.
func (tcauo *TableChangeAlertUpdateOne) SetFunctionName(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetFunctionName(s)
	return tcauo
}

// SetNillableFunctionName sets the "function_name" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableFunctionName(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetFunctionName(*s)
	}
	return tcauo
}

// ClearFunctionName clears the value of the "function_name" field.
func (tcauo *TableChangeAlertUpdateOne) ClearFunctionName() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearFunctionName()
	return tcauo
}

// SetTriggerName sets the "trigger_name" field.
func (tcauo *TableChangeAlertUpdateOne) SetTriggerName(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetTriggerName(s)
	return tcauo
}

// SetNillableTriggerName sets the "trigger_name" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableTriggerName(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetTriggerName(*s)
	}
	return tcauo
}

// ClearTriggerName clears the value of the "trigger_name" field.
func (tcauo *TableChangeAlertUpdateOne) ClearTriggerName() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearTriggerName()
	return tcauo
}

// SetListenerName sets the "listener_name" field.
func (tcauo *TableChangeAlertUpdateOne) SetListenerName(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetListenerName(s)
	return tcauo
}

// SetNillableListenerName sets the "listener_name" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableListenerName(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetListenerName(*s)
	}
	return tcauo
}

// ClearListenerName clears the value of the "listener_name" field.
func (tcauo *TableChangeAlertUpdateOne) ClearListenerName() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearListenerName()
	return tcauo
}

// SetEmailRecipients sets the "email_recipients" field.
func (tcauo *TableChangeAlertUpdateOne) SetEmailRecipients(s string) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetEmailRecipients(s)
	return tcauo
}

// SetNillableEmailRecipients sets the "email_recipients" field if the given value is not nil.
func (tcauo *TableChangeAlertUpdateOne) SetNillableEmailRecipients(s *string) *TableChangeAlertUpdateOne {
	if s != nil {
		tcauo.SetEmailRecipients(*s)
	}
	return tcauo
}

// ClearEmailRecipients clears the value of the "email_recipients" field.
func (tcauo *TableChangeAlertUpdateOne) ClearEmailRecipients() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearEmailRecipients()
	return tcauo
}

// SetEffectiveDate sets the "effective_date" field.
func (tcauo *TableChangeAlertUpdateOne) SetEffectiveDate(pg *pgtype.Date) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetEffectiveDate(pg)
	return tcauo
}

// ClearEffectiveDate clears the value of the "effective_date" field.
func (tcauo *TableChangeAlertUpdateOne) ClearEffectiveDate() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearEffectiveDate()
	return tcauo
}

// SetExpirationDate sets the "expiration_date" field.
func (tcauo *TableChangeAlertUpdateOne) SetExpirationDate(pg *pgtype.Date) *TableChangeAlertUpdateOne {
	tcauo.mutation.SetExpirationDate(pg)
	return tcauo
}

// ClearExpirationDate clears the value of the "expiration_date" field.
func (tcauo *TableChangeAlertUpdateOne) ClearExpirationDate() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearExpirationDate()
	return tcauo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (tcauo *TableChangeAlertUpdateOne) SetOrganization(o *Organization) *TableChangeAlertUpdateOne {
	return tcauo.SetOrganizationID(o.ID)
}

// Mutation returns the TableChangeAlertMutation object of the builder.
func (tcauo *TableChangeAlertUpdateOne) Mutation() *TableChangeAlertMutation {
	return tcauo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (tcauo *TableChangeAlertUpdateOne) ClearOrganization() *TableChangeAlertUpdateOne {
	tcauo.mutation.ClearOrganization()
	return tcauo
}

// Where appends a list predicates to the TableChangeAlertUpdate builder.
func (tcauo *TableChangeAlertUpdateOne) Where(ps ...predicate.TableChangeAlert) *TableChangeAlertUpdateOne {
	tcauo.mutation.Where(ps...)
	return tcauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcauo *TableChangeAlertUpdateOne) Select(field string, fields ...string) *TableChangeAlertUpdateOne {
	tcauo.fields = append([]string{field}, fields...)
	return tcauo
}

// Save executes the query and returns the updated TableChangeAlert entity.
func (tcauo *TableChangeAlertUpdateOne) Save(ctx context.Context) (*TableChangeAlert, error) {
	if err := tcauo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tcauo.sqlSave, tcauo.mutation, tcauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcauo *TableChangeAlertUpdateOne) SaveX(ctx context.Context) *TableChangeAlert {
	node, err := tcauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcauo *TableChangeAlertUpdateOne) Exec(ctx context.Context) error {
	_, err := tcauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcauo *TableChangeAlertUpdateOne) ExecX(ctx context.Context) {
	if err := tcauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcauo *TableChangeAlertUpdateOne) defaults() error {
	if _, ok := tcauo.mutation.UpdatedAt(); !ok {
		if tablechangealert.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tablechangealert.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tablechangealert.UpdateDefaultUpdatedAt()
		tcauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tcauo *TableChangeAlertUpdateOne) check() error {
	if v, ok := tcauo.mutation.Status(); ok {
		if err := tablechangealert.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.status": %w`, err)}
		}
	}
	if v, ok := tcauo.mutation.Name(); ok {
		if err := tablechangealert.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.name": %w`, err)}
		}
	}
	if v, ok := tcauo.mutation.DatabaseAction(); ok {
		if err := tablechangealert.DatabaseActionValidator(v); err != nil {
			return &ValidationError{Name: "database_action", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.database_action": %w`, err)}
		}
	}
	if v, ok := tcauo.mutation.Source(); ok {
		if err := tablechangealert.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.source": %w`, err)}
		}
	}
	if v, ok := tcauo.mutation.FunctionName(); ok {
		if err := tablechangealert.FunctionNameValidator(v); err != nil {
			return &ValidationError{Name: "function_name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.function_name": %w`, err)}
		}
	}
	if v, ok := tcauo.mutation.TriggerName(); ok {
		if err := tablechangealert.TriggerNameValidator(v); err != nil {
			return &ValidationError{Name: "trigger_name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.trigger_name": %w`, err)}
		}
	}
	if v, ok := tcauo.mutation.ListenerName(); ok {
		if err := tablechangealert.ListenerNameValidator(v); err != nil {
			return &ValidationError{Name: "listener_name", err: fmt.Errorf(`ent: validator failed for field "TableChangeAlert.listener_name": %w`, err)}
		}
	}
	if _, ok := tcauo.mutation.BusinessUnitID(); tcauo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TableChangeAlert.business_unit"`)
	}
	if _, ok := tcauo.mutation.OrganizationID(); tcauo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TableChangeAlert.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tcauo *TableChangeAlertUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TableChangeAlertUpdateOne {
	tcauo.modifiers = append(tcauo.modifiers, modifiers...)
	return tcauo
}

func (tcauo *TableChangeAlertUpdateOne) sqlSave(ctx context.Context) (_node *TableChangeAlert, err error) {
	if err := tcauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tablechangealert.Table, tablechangealert.Columns, sqlgraph.NewFieldSpec(tablechangealert.FieldID, field.TypeUUID))
	id, ok := tcauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TableChangeAlert.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tablechangealert.FieldID)
		for _, f := range fields {
			if !tablechangealert.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tablechangealert.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcauo.mutation.UpdatedAt(); ok {
		_spec.SetField(tablechangealert.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tcauo.mutation.Version(); ok {
		_spec.SetField(tablechangealert.FieldVersion, field.TypeInt, value)
	}
	if value, ok := tcauo.mutation.AddedVersion(); ok {
		_spec.AddField(tablechangealert.FieldVersion, field.TypeInt, value)
	}
	if value, ok := tcauo.mutation.Status(); ok {
		_spec.SetField(tablechangealert.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tcauo.mutation.Name(); ok {
		_spec.SetField(tablechangealert.FieldName, field.TypeString, value)
	}
	if value, ok := tcauo.mutation.DatabaseAction(); ok {
		_spec.SetField(tablechangealert.FieldDatabaseAction, field.TypeEnum, value)
	}
	if value, ok := tcauo.mutation.Source(); ok {
		_spec.SetField(tablechangealert.FieldSource, field.TypeEnum, value)
	}
	if value, ok := tcauo.mutation.TableName(); ok {
		_spec.SetField(tablechangealert.FieldTableName, field.TypeString, value)
	}
	if tcauo.mutation.TableNameCleared() {
		_spec.ClearField(tablechangealert.FieldTableName, field.TypeString)
	}
	if value, ok := tcauo.mutation.TopicName(); ok {
		_spec.SetField(tablechangealert.FieldTopicName, field.TypeString, value)
	}
	if tcauo.mutation.TopicNameCleared() {
		_spec.ClearField(tablechangealert.FieldTopicName, field.TypeString)
	}
	if value, ok := tcauo.mutation.Description(); ok {
		_spec.SetField(tablechangealert.FieldDescription, field.TypeString, value)
	}
	if tcauo.mutation.DescriptionCleared() {
		_spec.ClearField(tablechangealert.FieldDescription, field.TypeString)
	}
	if value, ok := tcauo.mutation.CustomSubject(); ok {
		_spec.SetField(tablechangealert.FieldCustomSubject, field.TypeString, value)
	}
	if tcauo.mutation.CustomSubjectCleared() {
		_spec.ClearField(tablechangealert.FieldCustomSubject, field.TypeString)
	}
	if value, ok := tcauo.mutation.FunctionName(); ok {
		_spec.SetField(tablechangealert.FieldFunctionName, field.TypeString, value)
	}
	if tcauo.mutation.FunctionNameCleared() {
		_spec.ClearField(tablechangealert.FieldFunctionName, field.TypeString)
	}
	if value, ok := tcauo.mutation.TriggerName(); ok {
		_spec.SetField(tablechangealert.FieldTriggerName, field.TypeString, value)
	}
	if tcauo.mutation.TriggerNameCleared() {
		_spec.ClearField(tablechangealert.FieldTriggerName, field.TypeString)
	}
	if value, ok := tcauo.mutation.ListenerName(); ok {
		_spec.SetField(tablechangealert.FieldListenerName, field.TypeString, value)
	}
	if tcauo.mutation.ListenerNameCleared() {
		_spec.ClearField(tablechangealert.FieldListenerName, field.TypeString)
	}
	if value, ok := tcauo.mutation.EmailRecipients(); ok {
		_spec.SetField(tablechangealert.FieldEmailRecipients, field.TypeString, value)
	}
	if tcauo.mutation.EmailRecipientsCleared() {
		_spec.ClearField(tablechangealert.FieldEmailRecipients, field.TypeString)
	}
	if value, ok := tcauo.mutation.EffectiveDate(); ok {
		_spec.SetField(tablechangealert.FieldEffectiveDate, field.TypeOther, value)
	}
	if tcauo.mutation.EffectiveDateCleared() {
		_spec.ClearField(tablechangealert.FieldEffectiveDate, field.TypeOther)
	}
	if value, ok := tcauo.mutation.ExpirationDate(); ok {
		_spec.SetField(tablechangealert.FieldExpirationDate, field.TypeOther, value)
	}
	if tcauo.mutation.ExpirationDateCleared() {
		_spec.ClearField(tablechangealert.FieldExpirationDate, field.TypeOther)
	}
	if tcauo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tablechangealert.OrganizationTable,
			Columns: []string{tablechangealert.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcauo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tablechangealert.OrganizationTable,
			Columns: []string{tablechangealert.OrganizationColumn},
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
	_spec.AddModifiers(tcauo.modifiers...)
	_node = &TableChangeAlert{config: tcauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tablechangealert.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcauo.mutation.done = true
	return _node, nil
}
