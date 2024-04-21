// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/delaycode"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/google/uuid"
)

// DelayCode is the model entity for the DelayCode schema.
type DelayCode struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// The time that this entity was created.
	CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
	// The last time that this entity was updated.
	UpdatedAt time.Time `json:"updatedAt" validate:"omitempty"`
	// The current version of this entity.
	Version int `json:"version" validate:"omitempty"`
	// Status holds the value of the "status" field.
	Status delaycode.Status `json:"status" validate:"required,oneof=A I"`
	// Code holds the value of the "code" field.
	Code string `json:"code" validate:"required,max=4"`
	// Description holds the value of the "description" field.
	Description string `json:"description" validate:"omitempty"`
	// FCarrierOrDriver holds the value of the "f_carrier_or_driver" field.
	FCarrierOrDriver bool `json:"fCarrierOrDriver" validate:"omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color" validate:"omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DelayCodeQuery when eager-loading is set.
	Edges        DelayCodeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DelayCodeEdges holds the relations/edges for other nodes in the graph.
type DelayCodeEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DelayCodeEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DelayCodeEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DelayCode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case delaycode.FieldFCarrierOrDriver:
			values[i] = new(sql.NullBool)
		case delaycode.FieldVersion:
			values[i] = new(sql.NullInt64)
		case delaycode.FieldStatus, delaycode.FieldCode, delaycode.FieldDescription, delaycode.FieldColor:
			values[i] = new(sql.NullString)
		case delaycode.FieldCreatedAt, delaycode.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case delaycode.FieldID, delaycode.FieldBusinessUnitID, delaycode.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DelayCode fields.
func (dc *DelayCode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case delaycode.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dc.ID = *value
			}
		case delaycode.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				dc.BusinessUnitID = *value
			}
		case delaycode.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				dc.OrganizationID = *value
			}
		case delaycode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dc.CreatedAt = value.Time
			}
		case delaycode.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dc.UpdatedAt = value.Time
			}
		case delaycode.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				dc.Version = int(value.Int64)
			}
		case delaycode.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dc.Status = delaycode.Status(value.String)
			}
		case delaycode.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				dc.Code = value.String
			}
		case delaycode.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				dc.Description = value.String
			}
		case delaycode.FieldFCarrierOrDriver:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field f_carrier_or_driver", values[i])
			} else if value.Valid {
				dc.FCarrierOrDriver = value.Bool
			}
		case delaycode.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				dc.Color = value.String
			}
		default:
			dc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DelayCode.
// This includes values selected through modifiers, order, etc.
func (dc *DelayCode) Value(name string) (ent.Value, error) {
	return dc.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the DelayCode entity.
func (dc *DelayCode) QueryBusinessUnit() *BusinessUnitQuery {
	return NewDelayCodeClient(dc.config).QueryBusinessUnit(dc)
}

// QueryOrganization queries the "organization" edge of the DelayCode entity.
func (dc *DelayCode) QueryOrganization() *OrganizationQuery {
	return NewDelayCodeClient(dc.config).QueryOrganization(dc)
}

// Update returns a builder for updating this DelayCode.
// Note that you need to call DelayCode.Unwrap() before calling this method if this DelayCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (dc *DelayCode) Update() *DelayCodeUpdateOne {
	return NewDelayCodeClient(dc.config).UpdateOne(dc)
}

// Unwrap unwraps the DelayCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dc *DelayCode) Unwrap() *DelayCode {
	_tx, ok := dc.config.driver.(*txDriver)
	if !ok {
		panic("ent: DelayCode is not a transactional entity")
	}
	dc.config.driver = _tx.drv
	return dc
}

// String implements the fmt.Stringer.
func (dc *DelayCode) String() string {
	var builder strings.Builder
	builder.WriteString("DelayCode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dc.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", dc.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", dc.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(dc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", dc.Version))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", dc.Status))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(dc.Code)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(dc.Description)
	builder.WriteString(", ")
	builder.WriteString("f_carrier_or_driver=")
	builder.WriteString(fmt.Sprintf("%v", dc.FCarrierOrDriver))
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(dc.Color)
	builder.WriteByte(')')
	return builder.String()
}

// DelayCodes is a parsable slice of DelayCode.
type DelayCodes []*DelayCode
