// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/commodity"
	"github.com/emoss08/trenova/ent/hazardousmaterial"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// Commodity is the model entity for the Commodity schema.
type Commodity struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Status holds the value of the "status" field.
	Status commodity.Status `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsHazmat holds the value of the "is_hazmat" field.
	IsHazmat bool `json:"is_hazmat,omitempty"`
	// UnitOfMeasure holds the value of the "unit_of_measure" field.
	UnitOfMeasure *commodity.UnitOfMeasure `json:"unit_of_measure,omitempty"`
	// MinTemp holds the value of the "min_temp" field.
	MinTemp *float64 `json:"min_temp,omitempty"`
	// MaxTemp holds the value of the "max_temp" field.
	MaxTemp *float64 `json:"max_temp,omitempty"`
	// SetPointTemp holds the value of the "set_point_temp" field.
	SetPointTemp *float64 `json:"set_point_temp,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommodityQuery when eager-loading is set.
	Edges                 CommodityEdges `json:"edges"`
	hazardous_material_id *uuid.UUID
	selectValues          sql.SelectValues
}

// CommodityEdges holds the relations/edges for other nodes in the graph.
type CommodityEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// HazardousMaterial holds the value of the hazardous_material edge.
	HazardousMaterial *HazardousMaterial `json:"hazardous_material,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommodityEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommodityEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// HazardousMaterialOrErr returns the HazardousMaterial value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommodityEdges) HazardousMaterialOrErr() (*HazardousMaterial, error) {
	if e.HazardousMaterial != nil {
		return e.HazardousMaterial, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: hazardousmaterial.Label}
	}
	return nil, &NotLoadedError{edge: "hazardous_material"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Commodity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case commodity.FieldIsHazmat:
			values[i] = new(sql.NullBool)
		case commodity.FieldMinTemp, commodity.FieldMaxTemp, commodity.FieldSetPointTemp:
			values[i] = new(sql.NullFloat64)
		case commodity.FieldStatus, commodity.FieldName, commodity.FieldUnitOfMeasure, commodity.FieldDescription:
			values[i] = new(sql.NullString)
		case commodity.FieldCreatedAt, commodity.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case commodity.FieldID, commodity.FieldBusinessUnitID, commodity.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		case commodity.ForeignKeys[0]: // hazardous_material_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Commodity fields.
func (c *Commodity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case commodity.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case commodity.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				c.BusinessUnitID = *value
			}
		case commodity.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				c.OrganizationID = *value
			}
		case commodity.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case commodity.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case commodity.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = commodity.Status(value.String)
			}
		case commodity.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case commodity.FieldIsHazmat:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_hazmat", values[i])
			} else if value.Valid {
				c.IsHazmat = value.Bool
			}
		case commodity.FieldUnitOfMeasure:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit_of_measure", values[i])
			} else if value.Valid {
				c.UnitOfMeasure = new(commodity.UnitOfMeasure)
				*c.UnitOfMeasure = commodity.UnitOfMeasure(value.String)
			}
		case commodity.FieldMinTemp:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_temp", values[i])
			} else if value.Valid {
				c.MinTemp = new(float64)
				*c.MinTemp = value.Float64
			}
		case commodity.FieldMaxTemp:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_temp", values[i])
			} else if value.Valid {
				c.MaxTemp = new(float64)
				*c.MaxTemp = value.Float64
			}
		case commodity.FieldSetPointTemp:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field set_point_temp", values[i])
			} else if value.Valid {
				c.SetPointTemp = new(float64)
				*c.SetPointTemp = value.Float64
			}
		case commodity.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = new(string)
				*c.Description = value.String
			}
		case commodity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field hazardous_material_id", values[i])
			} else if value.Valid {
				c.hazardous_material_id = new(uuid.UUID)
				*c.hazardous_material_id = *value.S.(*uuid.UUID)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Commodity.
// This includes values selected through modifiers, order, etc.
func (c *Commodity) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the Commodity entity.
func (c *Commodity) QueryBusinessUnit() *BusinessUnitQuery {
	return NewCommodityClient(c.config).QueryBusinessUnit(c)
}

// QueryOrganization queries the "organization" edge of the Commodity entity.
func (c *Commodity) QueryOrganization() *OrganizationQuery {
	return NewCommodityClient(c.config).QueryOrganization(c)
}

// QueryHazardousMaterial queries the "hazardous_material" edge of the Commodity entity.
func (c *Commodity) QueryHazardousMaterial() *HazardousMaterialQuery {
	return NewCommodityClient(c.config).QueryHazardousMaterial(c)
}

// Update returns a builder for updating this Commodity.
// Note that you need to call Commodity.Unwrap() before calling this method if this Commodity
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Commodity) Update() *CommodityUpdateOne {
	return NewCommodityClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Commodity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Commodity) Unwrap() *Commodity {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Commodity is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Commodity) String() string {
	var builder strings.Builder
	builder.WriteString("Commodity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", c.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("is_hazmat=")
	builder.WriteString(fmt.Sprintf("%v", c.IsHazmat))
	builder.WriteString(", ")
	if v := c.UnitOfMeasure; v != nil {
		builder.WriteString("unit_of_measure=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.MinTemp; v != nil {
		builder.WriteString("min_temp=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.MaxTemp; v != nil {
		builder.WriteString("max_temp=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.SetPointTemp; v != nil {
		builder.WriteString("set_point_temp=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Commodities is a parsable slice of Commodity.
type Commodities []*Commodity