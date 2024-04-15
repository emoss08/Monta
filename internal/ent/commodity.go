// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/commodity"
	"github.com/emoss08/trenova/internal/ent/hazardousmaterial"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/google/uuid"
)

// Commodity is the model entity for the Commodity schema.
type Commodity struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt" validate:"omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version" validate:"omitempty"`
	// Status holds the value of the "status" field.
	Status commodity.Status `json:"status" validate:"required,oneof=A I"`
	// Name holds the value of the "name" field.
	Name string `json:"name" validate:"required"`
	// IsHazmat holds the value of the "is_hazmat" field.
	IsHazmat bool `json:"isHazmat" validate:"omitempty"`
	// UnitOfMeasure holds the value of the "unit_of_measure" field.
	UnitOfMeasure string `json:"unitOfMeasure" validate:"omitempty,oneof=Pallet Tote Drum Cylinder Case Ampule Bag Bottle Pail Pieces IsoTank"`
	// MinTemp holds the value of the "min_temp" field.
	MinTemp int8 `json:"minTemp" validate:"omitempty,max=127"`
	// MaxTemp holds the value of the "max_temp" field.
	MaxTemp int8 `json:"maxTemp" validate:"omitempty,max=127"`
	// Description holds the value of the "description" field.
	Description string `json:"description" validate:"omitempty"`
	// HazardousMaterialID holds the value of the "hazardous_material_id" field.
	HazardousMaterialID *uuid.UUID `json:"hazardousMaterialId" validate:"omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommodityQuery when eager-loading is set.
	Edges        CommodityEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommodityEdges holds the relations/edges for other nodes in the graph.
type CommodityEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// HazardousMaterial holds the value of the hazardous_material edge.
	HazardousMaterial *HazardousMaterial `json:"hazardousMaterial"`
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
		case commodity.FieldHazardousMaterialID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case commodity.FieldIsHazmat:
			values[i] = new(sql.NullBool)
		case commodity.FieldVersion, commodity.FieldMinTemp, commodity.FieldMaxTemp:
			values[i] = new(sql.NullInt64)
		case commodity.FieldStatus, commodity.FieldName, commodity.FieldUnitOfMeasure, commodity.FieldDescription:
			values[i] = new(sql.NullString)
		case commodity.FieldCreatedAt, commodity.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case commodity.FieldID, commodity.FieldBusinessUnitID, commodity.FieldOrganizationID:
			values[i] = new(uuid.UUID)
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
		case commodity.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				c.Version = int(value.Int64)
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
				c.UnitOfMeasure = value.String
			}
		case commodity.FieldMinTemp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_temp", values[i])
			} else if value.Valid {
				c.MinTemp = int8(value.Int64)
			}
		case commodity.FieldMaxTemp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_temp", values[i])
			} else if value.Valid {
				c.MaxTemp = int8(value.Int64)
			}
		case commodity.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case commodity.FieldHazardousMaterialID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field hazardous_material_id", values[i])
			} else if value.Valid {
				c.HazardousMaterialID = new(uuid.UUID)
				*c.HazardousMaterialID = *value.S.(*uuid.UUID)
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
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", c.Version))
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
	builder.WriteString("unit_of_measure=")
	builder.WriteString(c.UnitOfMeasure)
	builder.WriteString(", ")
	builder.WriteString("min_temp=")
	builder.WriteString(fmt.Sprintf("%v", c.MinTemp))
	builder.WriteString(", ")
	builder.WriteString("max_temp=")
	builder.WriteString(fmt.Sprintf("%v", c.MaxTemp))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	if v := c.HazardousMaterialID; v != nil {
		builder.WriteString("hazardous_material_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Commodities is a parsable slice of Commodity.
type Commodities []*Commodity