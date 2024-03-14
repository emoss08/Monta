// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/accountingcontrol"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/generalledgeraccount"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// AccountingControl is the model entity for the AccountingControl schema.
type AccountingControl struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// RecThreshold holds the value of the "rec_threshold" field.
	RecThreshold int64 `json:"recThreshold"`
	// RecThresholdAction holds the value of the "rec_threshold_action" field.
	RecThresholdAction accountingcontrol.RecThresholdAction `json:"recThresholdAction"`
	// AutoCreateJournalEntries holds the value of the "auto_create_journal_entries" field.
	AutoCreateJournalEntries bool `json:"autoCreateJournalEntries"`
	// RestrictManualJournalEntries holds the value of the "restrict_manual_journal_entries" field.
	RestrictManualJournalEntries bool `json:"restrictManualJournalEntries"`
	// RequireJournalEntryApproval holds the value of the "require_journal_entry_approval" field.
	RequireJournalEntryApproval bool `json:"requireJournalEntryApproval"`
	// EnableRecNotifications holds the value of the "enable_rec_notifications" field.
	EnableRecNotifications bool `json:"enableRecNotifications"`
	// HaltOnPendingRec holds the value of the "halt_on_pending_rec" field.
	HaltOnPendingRec bool `json:"haltOnPendingRec"`
	// CriticalProcesses holds the value of the "critical_processes" field.
	CriticalProcesses string `json:"criticalProcesses"`
	// DefaultRevAccountID holds the value of the "default_rev_account_id" field.
	DefaultRevAccountID uuid.UUID `json:"defaultRevAccountId"`
	// DefaultExpAccountID holds the value of the "default_exp_account_id" field.
	DefaultExpAccountID uuid.UUID `json:"defaultExpAccountId"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountingControlQuery when eager-loading is set.
	Edges            AccountingControlEdges `json:"edges"`
	business_unit_id *uuid.UUID
	organization_id  *uuid.UUID
	selectValues     sql.SelectValues
}

// AccountingControlEdges holds the relations/edges for other nodes in the graph.
type AccountingControlEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// DefaultRevAccount holds the value of the default_rev_account edge.
	DefaultRevAccount *GeneralLedgerAccount `json:"default_rev_account,omitempty"`
	// DefaultExpAccount holds the value of the default_exp_account edge.
	DefaultExpAccount *GeneralLedgerAccount `json:"default_exp_account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountingControlEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountingControlEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// DefaultRevAccountOrErr returns the DefaultRevAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountingControlEdges) DefaultRevAccountOrErr() (*GeneralLedgerAccount, error) {
	if e.DefaultRevAccount != nil {
		return e.DefaultRevAccount, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: generalledgeraccount.Label}
	}
	return nil, &NotLoadedError{edge: "default_rev_account"}
}

// DefaultExpAccountOrErr returns the DefaultExpAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountingControlEdges) DefaultExpAccountOrErr() (*GeneralLedgerAccount, error) {
	if e.DefaultExpAccount != nil {
		return e.DefaultExpAccount, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: generalledgeraccount.Label}
	}
	return nil, &NotLoadedError{edge: "default_exp_account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountingControl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accountingcontrol.FieldAutoCreateJournalEntries, accountingcontrol.FieldRestrictManualJournalEntries, accountingcontrol.FieldRequireJournalEntryApproval, accountingcontrol.FieldEnableRecNotifications, accountingcontrol.FieldHaltOnPendingRec:
			values[i] = new(sql.NullBool)
		case accountingcontrol.FieldRecThreshold:
			values[i] = new(sql.NullInt64)
		case accountingcontrol.FieldRecThresholdAction, accountingcontrol.FieldCriticalProcesses:
			values[i] = new(sql.NullString)
		case accountingcontrol.FieldCreatedAt, accountingcontrol.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case accountingcontrol.FieldID, accountingcontrol.FieldDefaultRevAccountID, accountingcontrol.FieldDefaultExpAccountID:
			values[i] = new(uuid.UUID)
		case accountingcontrol.ForeignKeys[0]: // business_unit_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case accountingcontrol.ForeignKeys[1]: // organization_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountingControl fields.
func (ac *AccountingControl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accountingcontrol.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ac.ID = *value
			}
		case accountingcontrol.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case accountingcontrol.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = value.Time
			}
		case accountingcontrol.FieldRecThreshold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rec_threshold", values[i])
			} else if value.Valid {
				ac.RecThreshold = value.Int64
			}
		case accountingcontrol.FieldRecThresholdAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rec_threshold_action", values[i])
			} else if value.Valid {
				ac.RecThresholdAction = accountingcontrol.RecThresholdAction(value.String)
			}
		case accountingcontrol.FieldAutoCreateJournalEntries:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auto_create_journal_entries", values[i])
			} else if value.Valid {
				ac.AutoCreateJournalEntries = value.Bool
			}
		case accountingcontrol.FieldRestrictManualJournalEntries:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field restrict_manual_journal_entries", values[i])
			} else if value.Valid {
				ac.RestrictManualJournalEntries = value.Bool
			}
		case accountingcontrol.FieldRequireJournalEntryApproval:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field require_journal_entry_approval", values[i])
			} else if value.Valid {
				ac.RequireJournalEntryApproval = value.Bool
			}
		case accountingcontrol.FieldEnableRecNotifications:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable_rec_notifications", values[i])
			} else if value.Valid {
				ac.EnableRecNotifications = value.Bool
			}
		case accountingcontrol.FieldHaltOnPendingRec:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field halt_on_pending_rec", values[i])
			} else if value.Valid {
				ac.HaltOnPendingRec = value.Bool
			}
		case accountingcontrol.FieldCriticalProcesses:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field critical_processes", values[i])
			} else if value.Valid {
				ac.CriticalProcesses = value.String
			}
		case accountingcontrol.FieldDefaultRevAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field default_rev_account_id", values[i])
			} else if value != nil {
				ac.DefaultRevAccountID = *value
			}
		case accountingcontrol.FieldDefaultExpAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field default_exp_account_id", values[i])
			} else if value != nil {
				ac.DefaultExpAccountID = *value
			}
		case accountingcontrol.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value.Valid {
				ac.business_unit_id = new(uuid.UUID)
				*ac.business_unit_id = *value.S.(*uuid.UUID)
			}
		case accountingcontrol.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				ac.organization_id = new(uuid.UUID)
				*ac.organization_id = *value.S.(*uuid.UUID)
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AccountingControl.
// This includes values selected through modifiers, order, etc.
func (ac *AccountingControl) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the AccountingControl entity.
func (ac *AccountingControl) QueryOrganization() *OrganizationQuery {
	return NewAccountingControlClient(ac.config).QueryOrganization(ac)
}

// QueryBusinessUnit queries the "business_unit" edge of the AccountingControl entity.
func (ac *AccountingControl) QueryBusinessUnit() *BusinessUnitQuery {
	return NewAccountingControlClient(ac.config).QueryBusinessUnit(ac)
}

// QueryDefaultRevAccount queries the "default_rev_account" edge of the AccountingControl entity.
func (ac *AccountingControl) QueryDefaultRevAccount() *GeneralLedgerAccountQuery {
	return NewAccountingControlClient(ac.config).QueryDefaultRevAccount(ac)
}

// QueryDefaultExpAccount queries the "default_exp_account" edge of the AccountingControl entity.
func (ac *AccountingControl) QueryDefaultExpAccount() *GeneralLedgerAccountQuery {
	return NewAccountingControlClient(ac.config).QueryDefaultExpAccount(ac)
}

// Update returns a builder for updating this AccountingControl.
// Note that you need to call AccountingControl.Unwrap() before calling this method if this AccountingControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AccountingControl) Update() *AccountingControlUpdateOne {
	return NewAccountingControlClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AccountingControl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AccountingControl) Unwrap() *AccountingControl {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccountingControl is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AccountingControl) String() string {
	var builder strings.Builder
	builder.WriteString("AccountingControl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ac.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("rec_threshold=")
	builder.WriteString(fmt.Sprintf("%v", ac.RecThreshold))
	builder.WriteString(", ")
	builder.WriteString("rec_threshold_action=")
	builder.WriteString(fmt.Sprintf("%v", ac.RecThresholdAction))
	builder.WriteString(", ")
	builder.WriteString("auto_create_journal_entries=")
	builder.WriteString(fmt.Sprintf("%v", ac.AutoCreateJournalEntries))
	builder.WriteString(", ")
	builder.WriteString("restrict_manual_journal_entries=")
	builder.WriteString(fmt.Sprintf("%v", ac.RestrictManualJournalEntries))
	builder.WriteString(", ")
	builder.WriteString("require_journal_entry_approval=")
	builder.WriteString(fmt.Sprintf("%v", ac.RequireJournalEntryApproval))
	builder.WriteString(", ")
	builder.WriteString("enable_rec_notifications=")
	builder.WriteString(fmt.Sprintf("%v", ac.EnableRecNotifications))
	builder.WriteString(", ")
	builder.WriteString("halt_on_pending_rec=")
	builder.WriteString(fmt.Sprintf("%v", ac.HaltOnPendingRec))
	builder.WriteString(", ")
	builder.WriteString("critical_processes=")
	builder.WriteString(ac.CriticalProcesses)
	builder.WriteString(", ")
	builder.WriteString("default_rev_account_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.DefaultRevAccountID))
	builder.WriteString(", ")
	builder.WriteString("default_exp_account_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.DefaultExpAccountID))
	builder.WriteByte(')')
	return builder.String()
}

// AccountingControls is a parsable slice of AccountingControl.
type AccountingControls []*AccountingControl
